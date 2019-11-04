package requester

import (
	"fmt"
	"hash"
	"github.com/ZacharyGroff/CrowdCrack/apiclient"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
)

type PasswordRequester struct {
	config *models.Config
	client apiclient.ApiClient
	logger logger.Logger
	requestQueue queue.RequestQueue
	supportedHashes map[string]hash.Hash
	waiter waiter.Waiter
}

func NewPasswordRequester(p userinput.CmdLineConfigProvider, cl *apiclient.HashApiClient, l *logger.ConcurrentLogger, r *queue.HashingRequestQueue, w waiter.Sleeper) *PasswordRequester {
	return &PasswordRequester{
		config:          p.GetConfig(),
		client:          cl,
		logger:          l,
		requestQueue:    r,
		supportedHashes: models.GetSupportedHashFunctions(),
		waiter:          w,
	}
}

func (p PasswordRequester) Start() error {
	p.logger.LogMessage("Starting password requester")
	for {
		err := p.processOrSleep()
		if err != nil {
			return err
		}
	}
}

func (p PasswordRequester) processOrSleep() error {
	if p.requestQueue.Size() < 10 {
		err := p.addRequestToQueue()
		if err != nil {
			return err
		}
	} else {
		p.waiter.Wait()
	}

	return nil
}

func (p PasswordRequester) addRequestToQueue() error {
	hash, hashName, err := p.getHash()
	if err != nil {
		return err
	}

	passwords, err := p.getPasswords()
	if err != nil {
		return err
	}
	numPasswords := len(passwords)

	if numPasswords < 1 {
		p.logger.LogMessage("Requester received a response with zero passwords contained.")
		p.waiter.Wait()
	} else {
		if p.config.Verbose {
			logMessage := fmt.Sprintf("Requester has created hashing request with hash name: %s and %d passwords", hashName, numPasswords)
			p.logger.LogMessage(logMessage)
		}
		hashingRequest := models.HashingRequest{hash, hashName, passwords}
		err = p.requestQueue.Put(hashingRequest)
		if err != nil {
			return err
		}
	}

	return nil
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
