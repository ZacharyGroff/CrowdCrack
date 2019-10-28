package requester

import (
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/apiclient"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
	"hash"
)

type PasswordRequester struct {
	config *models.Config
	client apiclient.ApiClient
	requestQueue queue.RequestQueue
	supportedHashes map[string]hash.Hash
	waiter waiter.Waiter
}

func NewPasswordRequester(p userinput.CmdLineConfigProvider, cl *apiclient.HashApiClient, r *queue.HashingRequestQueue, logger *logger.GenericLogger) *PasswordRequester {
	c := p.GetConfig()
	s := getSupportedHashes()
	w := getWaiter(logger)
	return &PasswordRequester{c, cl, r, s, w}
}

func getSupportedHashes() map[string]hash.Hash {
	return models.GetSupportedHashFunctions()
}

func getWaiter(logger logger.Logger) waiter.Sleeper {
	sleepSeconds := 60
	isLogging := true
	logMessage := fmt.Sprintf("Password requester sleeping for %d seconds", sleepSeconds)

	return waiter.NewSleeper(sleepSeconds, isLogging, logMessage, logger)
}

func (p PasswordRequester) Start() error {
	for {
		err := p.processOrSleep()
		if err != nil {
			return err
		}
	}
}

func (p PasswordRequester) processOrSleep() error {
	if p.requestQueue.Size() < 2 {
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

	if len(passwords) < 1 {
		p.waiter.Wait()
	} else {
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
	numPasswords := 1000

	statusCode, passwords := p.client.GetPasswords(numPasswords)
	if statusCode != 200 {
		err := fmt.Errorf("Unexpected response from api on password request with status code: %d\n", statusCode)
		return passwords, err
	}

	return passwords, nil
}
