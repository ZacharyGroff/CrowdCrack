package encoder

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
	"log"
)

type Hasher struct {
	config *models.ClientConfig
	requestQueue queue.RequestQueue
	submissionQueue queue.SubmissionQueue
	waiter waiter.Waiter
}

func NewHasher(p userinput.CmdLineConfigProvider, r *queue.HashingRequestQueue, s *queue.HashingSubmissionQueue) *Hasher {
	c := p.GetClientConfig()
	w := getWaiter()
	return &Hasher{c, r, s, w}
}

func getWaiter() waiter.Waiter {
	sleepDuration := 5
	isLogging := true
	logMessage := fmt.Sprintf("No requests in queue. Hasher sleeping for %d seconds", sleepDuration)

	return waiter.NewSleeper(sleepDuration, isLogging, logMessage)
}

func (e Hasher) Start() error {
	log.Println("Starting hasher...")
	for {
		err := e.processOrSleep()
		if err != nil {
			return err
		}
	}
}

func (e Hasher) processOrSleep() error {
	hashingRequest, err := e.requestQueue.Get()
	if err != nil {
		e.waiter.Wait()
	} else {
		err = e.handleHashingRequest(hashingRequest)
		if err != nil {
			return err
		}
	}
	
	return nil
}

func (e Hasher) handleHashingRequest(hashingRequest models.HashingRequest) error {
	hashSubmission, err := e.getHashSubmission(hashingRequest)
	if err != nil {
		return err
	}
	
	err = e.submissionQueue.Put(hashSubmission)
	for err != nil {
		return err
	}

	return nil
}

func (e Hasher) getHashSubmission(hashingRequest models.HashingRequest) (models.HashSubmission, error) {
	hashFunction, err := e.getHashFunction(hashingRequest.HashName)
	if err != nil {
		return models.HashSubmission{}, err
	}

	passwordHashes := getPasswordHashes(hashFunction, hashingRequest.Passwords)

	return models.HashSubmission{hashingRequest.HashName, passwordHashes}, nil
}

func (e Hasher) getHashFunction(hashName string) (func([]byte) [32]byte, error) {
	switch hashName {
	case "sha256":
		return sha256.Sum256, nil
	default:
		return nil, fmt.Errorf("%s is not a supported hash. If the hash is currently available in golang crypto package, please create a GitHub issue to have support for it added.", hashName)
	}
}

func getPasswordHashes(hashFunction func([]byte) [32]byte, passwords []string) []string {
	var passwordHashes []string
	for _, password := range passwords {
		passwordHash := getPasswordHash(hashFunction, password)
		passwordHashes = append(passwordHashes, passwordHash)
	}

	return passwordHashes
}

func getPasswordHash(hashFunction func([]byte) [32]byte, password string) string {
	hash := hashFunction([]byte(password))
	humanReadableHash := hex.EncodeToString(hash[:])

	return password + ":" + humanReadableHash
}
