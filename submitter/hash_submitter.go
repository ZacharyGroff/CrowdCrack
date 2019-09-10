package submitter

import (
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/queue"
)

type HashSubmitter struct {
	config *config.ClientConfig
	hashes queue.FlushingQueue
}

func NewHashSubmitter(c *config.ClientConfig, q *queue.HashQueue) *HashSubmitter {
	return &HashSubmitter{c, q}
}

func (h HashSubmitter) Submit() error {
	return nil
}
