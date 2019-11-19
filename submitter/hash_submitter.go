package submitter

import (
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/apiclient"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
)

type HashSubmitter struct {
	config          *models.Config
	client          interfaces.ApiClient
	logger          interfaces.Logger
	submissionQueue interfaces.SubmissionQueue
	stopQueue       interfaces.ClientStopQueue
	waiter          interfaces.Waiter
}

func NewHashSubmitter(p interfaces.ConfigProvider, c *apiclient.HashApiClient, l *logger.ConcurrentLogger, s *queue.HashingSubmissionQueue, cl *queue.ClientStopReasonQueue, w waiter.Sleeper) *HashSubmitter {
	return &HashSubmitter{
		config:          p.GetConfig(),
		client:          c,
		logger:          l,
		stopQueue:       cl,
		submissionQueue: s,
		waiter:          w,
	}
}

func (h HashSubmitter) Start() error {
	h.logger.LogMessage("Starting HashSubmitter...")
	for {
		stopReason, err := h.stopQueue.Get()
		if err == nil {
			h.stop()
			err = fmt.Errorf("Submitter observed updateStopQueue reason:\n\t%+v", stopReason)
			return err
		}

		err = h.processOrSleep()
		if err != nil {
			h.updateStopQueue(err)
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

func (h HashSubmitter) updateStopQueue(err error) {
	stopReason := models.ClientStopReason{
		Requester: "",
		Encoder:   "",
		Submitter: err.Error(),
	}

	var i uint16
	for i = 0; i < h.config.Threads - 1; i++ {
		h.stopQueue.Put(stopReason)
	}
}

func (h HashSubmitter) stop() {
	return
}
