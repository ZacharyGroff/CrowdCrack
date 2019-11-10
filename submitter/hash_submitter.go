package submitter

import (
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/apiclient"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
)

type HashSubmitter struct {
	config          *models.Config
	client          interfaces.ApiClient
	logger          interfaces.Logger
	submissionQueue interfaces.SubmissionQueue
	waiter          interfaces.Waiter
}

func NewHashSubmitter(p userinput.CmdLineConfigProvider, cl *apiclient.HashApiClient, l *logger.ConcurrentLogger, s *queue.HashingSubmissionQueue, w waiter.Sleeper) *HashSubmitter {
	return &HashSubmitter{
		config:          p.GetConfig(),
		client:          cl,
		logger:          l,
		submissionQueue: s,
		waiter:          w,
	}
}

func (h HashSubmitter) Start() error {
	h.logger.LogMessage("Starting HashSubmitter...")
	for {
		err := h.processOrSleep()
		if err != nil {
			return err
		}
	}
}

func (h HashSubmitter) processOrSleep() error {
	if h.submissionQueue.Size() > 0 {
		err := h.processSubmission()
		if err != nil {
			return err
		}
	} else {
		h.waiter.Wait()
	}

	return nil
}

func (h HashSubmitter) processSubmission() error {
	hashSubmission, err := h.submissionQueue.Get()
	if err != nil {
		return err
	}

	if h.config.Verbose {
		numResults := len(hashSubmission.Results)
		logMessage := fmt.Sprintf("Submitter has received hash submission with hash type: %s and %d results", hashSubmission.HashType, numResults)
		h.logger.LogMessage(logMessage)
	}

	statusCode := h.client.SubmitHashes(hashSubmission)
	if statusCode != 200 {
		return fmt.Errorf("Unexpected response from api on hash submission with status code: %d\n", statusCode)
	}

	if h.config.Verbose {
		h.logger.LogMessage("Submitter has successfully submitted")
	}

	return nil
}
