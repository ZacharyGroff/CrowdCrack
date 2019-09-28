package submitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
	"log"
	"net/http"
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
		if h.submissionQueue.Size() > 0 {
			hashSubmission, err := h.submissionQueue.Get()
			if err != nil {
				return err
			}

			jsonHashSubmission, err := json.Marshal(hashSubmission)
			if err != nil {
				return err
			}		

			response, err := http.Post(h.config.ServerAddress + "/hashes", "application/json", bytes.NewBuffer(jsonHashSubmission))
			if err != nil {
				return err
			}

			log.Println(response)
		} else {
			h.waiter.Wait()
		}
	}
}
