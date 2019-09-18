package submitter

import (
	"bytes"
	"log"
	"time"
	"encoding/json"
	"net/http"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/queue"
)

type HashSubmitter struct {
	config *config.ClientConfig
	submissionQueue queue.SubmissionQueue
}

func NewHashSubmitter(c *config.ClientConfig, q *queue.HashingSubmissionQueue) *HashSubmitter {
	return &HashSubmitter{c, q}
}

func (h HashSubmitter) Submit() error {
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
			sleepDurationSeconds := time.Duration(5)
			log.Printf("No submissions in queue. HashSubmitter sleeping for %d seconds\n", sleepDurationSeconds)
			time.Sleep(sleepDurationSeconds * time.Second)
		}
	}
}
