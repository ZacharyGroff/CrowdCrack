package submitter

import (
	"bytes"
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
)

type HashSubmitter struct {
	config *config.ClientConfig
	submissionQueue queue.SubmissionQueue
	waiter waiter.Waiter
}

func NewHashSubmitter(c *config.ClientConfig, q *queue.HashingSubmissionQueue) *HashSubmitter {
	w := getWaiter()
	return &HashSubmitter{c, q, w}
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
	jsonHashSubmission, err := json.Marshal(hashSubmission)
	if err != nil {
		return err
	}
	response, err := http.Post(h.config.ServerAddress+"/hashes", "application/json", bytes.NewBuffer(jsonHashSubmission))
	if err != nil {
		return err
	}
	log.Println(response)

	return nil
}
