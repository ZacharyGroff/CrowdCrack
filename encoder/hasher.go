package encoder

import (
	"fmt"
	"hash"
	"io"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
)

type Hasher struct {
	config *models.Config
	logger logger.Logger
	requestQueue queue.RequestQueue
	submissionQueue queue.SubmissionQueue
	waiter waiter.Waiter
}

func NewHasher(p userinput.CmdLineConfigProvider, l *logger.GenericLogger, r *queue.HashingRequestQueue, s *queue.HashingSubmissionQueue) *Hasher {
	c := p.GetConfig()
	w := getWaiter(l)
	return &Hasher{
		config:          c,
		logger:          l,
		requestQueue:    r,
		submissionQueue: s,
		waiter:          w,
	}
}

func getWaiter(logger logger.Logger) waiter.Waiter {
	sleepDuration := 5
	isLogging := true
	logMessage := fmt.Sprintf("No requests in queue. Hasher sleeping for %d seconds", sleepDuration)

	return waiter.NewSleeper(sleepDuration, isLogging, logMessage, logger)
}

func (e Hasher) Start() error {
	e.logger.LogMessage("Starting hasher...")
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
	hashSubmission := e.getHashSubmission(hashingRequest)

	err := e.submissionQueue.Put(hashSubmission)
	for err != nil {
		return err
	}

	return nil
}

func (e Hasher) getHashSubmission(hashingRequest models.HashingRequest) models.HashSubmission {
	passwordHashes := getPasswordHashes(hashingRequest.Hash, hashingRequest.Passwords)

	return models.HashSubmission{hashingRequest.HashName, passwordHashes}
}

func getPasswordHashes(hash hash.Hash, passwords []string) []string {
	var passwordHashes []string
	for _, password := range passwords {
		passwordHash := getPasswordHash(hash, password)
		passwordHashes = append(passwordHashes, passwordHash)
	}

	return passwordHashes
}

func getPasswordHash(hash hash.Hash, password string) string {
	io.WriteString(hash, password)
	humanReadableHash := fmt.Sprintf("%x", hash.Sum(nil))
	hash.Reset()
	return password + ":" + humanReadableHash
}
