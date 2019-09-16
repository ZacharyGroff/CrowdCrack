package encoder

import (
	"log"
	"github.com/ZacharyGroff/CrowdCrack/config"
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

func (e Hasher) Encode() error {
	log.Println("Starting encoding...")
	return nil
}
