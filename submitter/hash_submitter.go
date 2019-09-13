package submitter

import (
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/queue"
)

type HashSubmitter struct {
	config *config.ClientConfig
	hashes queue.FlushingQueue
	submissionQueue queue.SubmissionQueue
}

func NewHashSubmitter(c *config.ClientConfig, q *queue.HashQueue, s *queue.HashingSubmissionQueue) *HashSubmitter {
	return &HashSubmitter{c, q, s}
}

func (h HashSubmitter) Submit() error {
	return nil
}
