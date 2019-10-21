package requester

import (
	"fmt"
	"hash"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"crypto/sha256"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
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
	return map[string]hash.Hash {
		"md4": md4.New(),
		"md5": md5.New(),
		"sha1": sha1.New(),
		"sha256": sha256.New(),
		"sha512": sha512.New(),
		"ripemd160": ripemd160.New(),
		"sha3_224": sha3.New224(),
		"sha3_256": sha3.New256(),
		"sha3_384": sha3.New384(),
		"sha3_512": sha3.New512(),
		"sha512_224": sha512.New512_224(),
		"sha512_256": sha512.New512_256(),
	}
}

func getWaiter(logger logger.Logger) waiter.Sleeper {
	sleepSeconds := 60
	isLogging := true
	logMessage := fmt.Sprintf("Request queue full. Password requester sleeping for %d seconds", sleepSeconds)

	return waiter.NewSleeper(sleepSeconds, isLogging, logMessage, logger)
}

func (p PasswordRequester) Start() error {
	for {
		err := p.processOrWait()
		if err != nil {
			return err
		}
	}
}

func (p PasswordRequester) processOrWait() error {
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

	hashingRequest := models.HashingRequest{hash, hashName, passwords}
	err = p.requestQueue.Put(hashingRequest)

	return err
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
