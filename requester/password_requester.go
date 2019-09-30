package requester

import (
	"fmt"
	"hash"
	"log"
	"time"
	"crypto/sha256"
	"github.com/ZacharyGroff/CrowdCrack/apiclient"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
)

type PasswordRequester struct {
	config *config.ClientConfig
	client apiclient.ApiClient
	supportedHashes map[string]hash.Hash
	requestQueue queue.RequestQueue
}

func NewPasswordRequester(c *config.ClientConfig, cl *apiclient.HashApiClient, r *queue.HashingRequestQueue) *PasswordRequester {
	s := getSupportedHashes()
	return &PasswordRequester{c,cl, s, r}
}

func getSupportedHashes() map[string]hash.Hash {
	return map[string]hash.Hash {
		"sha256": sha256.New(),
	}
}

func (p PasswordRequester) Start() error {
	for {
		if p.requestQueue.Size() < 2 { 
			err := p.addRequestToQueue()
			if err != nil {
				return err
			}
		} else {
			p.sleep()
		}
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

	hashingRequest := models.HashingRequest{hash, hashName, passwords}
	p.requestQueue.Put(hashingRequest)
	
	return nil
}

func (p PasswordRequester) sleep() {
	sleepDurationSeconds := time.Duration(60)
	log.Printf("Request queue full. Password requester sleeping for %d seconds\n", sleepDurationSeconds)
	time.Sleep(sleepDurationSeconds * time.Second)
}

func (p PasswordRequester) getHash() (hash.Hash, string, error) {
	hashName, err := p.requestHashName()
	if err != nil {
		return nil, "", err
	}

	currentHash, isSupported := p.supportedHashes[hashName]
	if !isSupported {
		err = fmt.Errorf("Current hash: %s is unsupported", hashName)
		return nil, "", err
	}

	return currentHash, hashName, nil
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
