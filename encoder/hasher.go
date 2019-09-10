package encoder

import (
	"hash"
	"log"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/queue"
)

type Hasher struct {
	config *config.ClientConfig
 	hashes queue.FlushingQueue
	passwords queue.Queue
}

func NewHasher(c *config.ClientConfig, h *queue.HashQueue, p *queue.PasswordQueue) *Hasher {
	return &Hasher{c, h, p}
}

func (e Hasher) Encode(h hash.Hash) error {
	log.Println("Starting encoding...")
	return nil
}
