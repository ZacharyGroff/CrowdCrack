package encoder

import (
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
	"hash"
	"io"
)

type Hasher struct {
	config          *models.Config
	logger          interfaces.Logger
	requestQueue    interfaces.RequestQueue
	submissionQueue interfaces.SubmissionQueue
	waiter          interfaces.Waiter
}

func NewHasher(p userinput.CmdLineConfigProvider, l *logger.ConcurrentLogger, r *queue.HashingRequestQueue, s *queue.HashingSubmissionQueue, w waiter.Sleeper) *Hasher {
	return &Hasher{
		config:          p.GetConfig(),
		logger:          l,
		requestQueue:    r,
		submissionQueue: s,
		waiter:          w,
	}
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
	if e.config.Verbose {
		numResults := len(hashSubmission.Results)
		logMessage := fmt.Sprintf("Hasher has created hash submission with hash type: %s and %d results", hashSubmission.HashType, numResults)
		e.logger.LogMessage(logMessage)
	}

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
