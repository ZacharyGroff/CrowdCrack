package client

import (
	"log"
	"github.com/ZacharyGroff/CrowdCrack/config"
)

type Client struct {
	config *config.ClientConfig
}

func NewClient(c *config.ClientConfig) Client {
	return Client{c}
}

func (client Client) Start() {
	log.Println("Starting Client...")
}

func (client Client) Stop() {
	log.Println("Stopping Client...")
}
