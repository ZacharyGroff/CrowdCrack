package requester

import (
	"hash"
	"crypto/sha256"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/queue"
)

type PasswordRequester struct {
	config *config.ClientConfig
	passwords queue.Queue
}

func NewPasswordRequester(c *config.ClientConfig, q *queue.PasswordQueue) *PasswordRequester {
	return &PasswordRequester{c, q}
}

func (p PasswordRequester) Request() (hash.Hash, error) {
	return sha256.New(), nil
}
