package submitter

import (
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/apiclient"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
)

type HashSubmitter struct {
	config *models.Config
	client apiclient.ApiClient
	logger logger.Logger
	submissionQueue queue.SubmissionQueue
	waiter waiter.Waiter
}

func NewHashSubmitter(p userinput.CmdLineConfigProvider, cl *apiclient.HashApiClient, l *logger.GenericLogger, s *queue.HashingSubmissionQueue) *HashSubmitter {
	c := p.GetConfig()
	w := getWaiter()
	return &HashSubmitter{
		config:          c,
		client:          cl,
		logger:          l,
		submissionQueue: s,
		waiter:          w,
	}
}

func getWaiter() waiter.Sleeper {
	sleepSeconds := 5
	logMessage := fmt.Sprintf("No submissions in queue. HashSubmitter sleeping for %d seconds\n", sleepSeconds)
	isLogging := true
	
	return waiter.NewSleeper(sleepSeconds, isLogging, logMessage)
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

	statusCode := h.client.SubmitHashes(hashSubmission)
	if statusCode != 200 {
		return fmt.Errorf("Unexpected response from api on hash submission with status code: %d\n", statusCode)
	}

	return nil
}
