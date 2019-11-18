package requester

import (
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/apiclient"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
	"hash"
	"os"
)

type PasswordRequester struct {
	config          *models.Config
	client          interfaces.ApiClient
	logger          interfaces.Logger
	requestQueue    interfaces.RequestQueue
	stopQueue       interfaces.ClientStopQueue
	supportedHashes map[string]hash.Hash
	waiter          interfaces.Waiter
}

func NewPasswordRequester(p interfaces.ConfigProvider, cl *apiclient.HashApiClient, l *logger.ConcurrentLogger, r *queue.HashingRequestQueue, c *queue.ClientStopReasonQueue, w waiter.Sleeper) *PasswordRequester {
	return &PasswordRequester{
		config:          p.GetConfig(),
		client:          cl,
		logger:          l,
		requestQueue:    r,
		stopQueue:       c,
		supportedHashes: models.GetSupportedHashFunctions(),
		waiter:          w,
	}
}

func (p PasswordRequester) Start() error {
	p.logger.LogMessage("Starting password requester")
	for {
		err := p.processOrStop()
		if err != nil {
			p.stopWithError(err)
			return err
		}
	}
}

func (p PasswordRequester) processOrStop() error {
	stopReason, err := p.stopQueue.Get()
	if err != nil {
		p.stop(stopReason)
		return nil
	}
	return p.processOrWait()
}

func (p PasswordRequester) processOrWait() error {
	if p.requestQueue.Size() < 10 {
		err := p.process()
		if err != nil {
			return err
		}
	} else {
		p.waiter.Wait()
	}

	return nil
}

func (p PasswordRequester) process() error {
	hashingRequest, err := p.getHashingRequest()
	if err != nil {
		return err
	}
	numPasswords := len(hashingRequest.Passwords)

	if numPasswords < 1 {
		p.logger.LogMessage("Requester received a response with zero passwords contained.")
		p.waiter.Wait()
	} else {
		err := p.addRequestToQueue(hashingRequest)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p PasswordRequester) addRequestToQueue(hashingRequest models.HashingRequest) error {
	if p.config.Verbose {
		p.logRequestCreated(hashingRequest)
	}

	err := p.requestQueue.Put(hashingRequest)

	return err
}

func (p PasswordRequester) logRequestCreated(hashingRequest models.HashingRequest) {
	numPasswords := len(hashingRequest.Passwords)
	logMessage := fmt.Sprintf("Requester has created hashing request with hash name: %s and %d passwords", hashingRequest.HashName, numPasswords)
	p.logger.LogMessage(logMessage)
}

func (p PasswordRequester) getHashingRequest() (models.HashingRequest, error) {
	var hashingRequest models.HashingRequest

	hash, hashName, err := p.getHash()
	if err != nil {
		return hashingRequest, err
	}

	passwords, err := p.getPasswords()
	if err != nil {
		return hashingRequest, err
	}

	hashingRequest = models.HashingRequest{hash, hashName, passwords}

	return hashingRequest, nil
}

func (p PasswordRequester) getHash() (hash.Hash, string, error) {
	hashName, err := p.requestHashName()
	if err != nil {
		return nil, "", err
	}

	hashFunction, err := p.getHashFunction(hashName)
	if err != nil {
		return nil, "", err
	}

	return hashFunction, hashName, nil
}

func (p PasswordRequester) getHashFunction(hashName string) (hash.Hash, error) {
	currentHash, isSupported := p.supportedHashes[hashName]
	if !isSupported {
		err := fmt.Errorf("Current hash: %s is unsupported\n", hashName)
		return nil, err
	}

	return currentHash, nil
}

func (p PasswordRequester) requestHashName() (string, error) {
	statusCode, hashName := p.client.GetHashName()
	if statusCode != 200 {
		err := fmt.Errorf("Unexpected response from api on hash name request with status code: %d\n", statusCode)
		return "", err
	}

	return hashName, nil
}

func (p PasswordRequester) getPasswords() ([]string, error) {
	statusCode, passwords := p.client.GetPasswords(int(p.config.PasswordRequestSize))
	if statusCode != 200 {
		err := fmt.Errorf("Unexpected response from api on password request with status code: %d\n", statusCode)
		return passwords, err
	}

	return passwords, nil
}

func (p PasswordRequester) stopWithError(err error) {
	stopReason := models.ClientStopReason{
		Requester: err.Error(),
		Encoder:   "",
		Submitter: "",
	}

	var i uint16
	for i = 0; i < p.config.Threads - 1; i++ {
		p.stopQueue.Put(stopReason)
	}

	p.stop(stopReason)
}

func (p PasswordRequester) stop(reason models.ClientStopReason) {
	if reason.Encoder != "" {
		logMessage := fmt.Sprintf("Stopping requester because encoder hit exception: %s", reason.Encoder)
		p.logger.LogMessage(logMessage)
	} else if reason.Requester != "" {
		logMessage := fmt.Sprintf("Stopping requester because requester hit exception: %s", reason.Requester)
		p.logger.LogMessage(logMessage)
	} else if reason.Submitter != "" {
		logMessage := fmt.Sprintf("Stopping requester because submitter hit exception: %s", reason.Submitter)
		p.logger.LogMessage(logMessage)
	}

	os.Exit(0)
}
