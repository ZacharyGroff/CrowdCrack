package encoder

import (
	"fmt"
	"log"
	"time"
	"crypto/sha256"
	"encoding/hex"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
)

type Hasher struct {
	config *config.ClientConfig
	requestQueue queue.RequestQueue
	submissionQueue queue.SubmissionQueue
}

func NewHasher(c *config.ClientConfig, r *queue.HashingRequestQueue, s *queue.HashingSubmissionQueue) *Hasher {
	return &Hasher{c, r, s}
}

func (e Hasher) Start() error {
	log.Println("Starting hasher...")
	for {
		err := e.processOrSleep()
		if err != nil {
			return err
		}
	}

	return nil
}

func (e Hasher) processOrSleep() error {
	hashingRequest, err := e.requestQueue.Get()
	if err != nil {
		sleep()
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

	var passwordHashes []string
	for _, password := range hashingRequest.Passwords {
		passwordHash := getPasswordHash(hashFunction, password)
		passwordHashes = append(passwordHashes, passwordHash)
	}

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

func sleep() {
	sleepDurationSeconds := time.Duration(5)
	log.Printf("No requests in queue. Hasher sleeping for %d seconds\n", sleepDurationSeconds)
	time.Sleep(sleepDurationSeconds * time.Second)
}

func getPasswordHash(hashFunction func([]byte) [32]byte, password string) string {
	hash := hashFunction([]byte(password))
	humanReadableHash := hex.EncodeToString(hash[:])

	return password + ":" + humanReadableHash
}

