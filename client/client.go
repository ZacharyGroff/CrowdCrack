package client

import (
	"log"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/encoder"
)

type Client struct {
	config *config.ClientConfig
	encoder encoder.Encoder
}

func NewClient(c *config.ClientConfig, e *encoder.Hasher) Client {
	return Client{c, e}
}

func (client Client) Start() {
	log.Println("Starting Client...")
}

func (client Client) Stop() {
	log.Println("Stopping Client...")
}
