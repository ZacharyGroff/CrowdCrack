package submitter

import (
	"fmt"
	"log"
	"github.com/ZacharyGroff/CrowdCrack/apiclient"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
)

type HashSubmitter struct {
	config *config.ClientConfig
	client apiclient.ApiClient
	submissionQueue queue.SubmissionQueue
	waiter waiter.Waiter
}

func NewHashSubmitter(c *config.ClientConfig, q *queue.HashingSubmissionQueue, cl *apiclient.HashApiClient) *HashSubmitter {
	w := getWaiter()
	return &HashSubmitter{c, cl, q, w}
}

func getWaiter() waiter.Sleeper {
	sleepSeconds := 5
	logMessage := fmt.Sprintf("No submissions in queue. HashSubmitter sleeping for %d seconds\n", sleepSeconds)
	isLogging := true
	
	return waiter.NewSleeper(sleepSeconds, isLogging, logMessage)
}

func (h HashSubmitter) Start() error {
	log.Println("Starting HashSubmitter...")
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
