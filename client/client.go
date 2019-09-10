package client

import (
	"log"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/encoder"
	"github.com/ZacharyGroff/CrowdCrack/requester"
	"github.com/ZacharyGroff/CrowdCrack/submitter"
)

type Client struct {
	config *config.ClientConfig
	encoder encoder.Encoder
	requester requester.Requester
	submitter submitter.Submitter
}

func NewClient(c *config.ClientConfig, e *encoder.Hasher, r *requester.PasswordRequester, s *submitter.HashSubmitter) Client {
	return Client{c, e, r, s}
}

func (c Client) Start() {
	log.Println("Starting Client...")
	for {
		hash, err := c.requester.Request()
		if err != nil {
			log.Println(err)
		}
		c.encoder.Encode(hash)
		c.submitter.Submit()
	}
}

func (c Client) Stop() {
	log.Println("Stopping Client...")
}
