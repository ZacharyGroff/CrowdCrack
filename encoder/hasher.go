package encoder

import (
	"fmt"
	"log"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
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
	passwordHashes, err := getPasswordHashes(hashingRequest.HashName, hashingRequest.Passwords)
	if err != nil {
		return models.HashSubmission{}, err
	}

	return models.HashSubmission{hashingRequest.HashName, passwordHashes}, nil
}

func getPasswordHashes(hashName string, passwords []string) ([]string, error) {
	var passwordHashes []string
	for _, password := range passwords {
		passwordHash, err := getPasswordHash(hashName, password)
		if err != nil {
			return nil, err
		}
		passwordHashes = append(passwordHashes, passwordHash)
	}

	return passwordHashes, nil
}

func getPasswordHash(hashName string, password string) (string, error) {
	var humanReadableHash string
	switch hashName {
	case "md5":
		hash := md5.Sum([]byte(password))
		humanReadableHash = hex.EncodeToString(hash[:])
	case "sha256":
		hash := sha256.Sum256([]byte(password))
		humanReadableHash = hex.EncodeToString(hash[:])
	default:
		return "", fmt.Errorf("%s is not a supported hash. If the hash is currently available in golang crypto package, please create a GitHub issue to have support for it added.", hashName)
	}

	return password + ":" + humanReadableHash, nil
}
