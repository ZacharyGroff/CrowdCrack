package client

import (
	"log"
	"github.com/ZacharyGroff/CrowdCrack/encoder"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/requester"
	"github.com/ZacharyGroff/CrowdCrack/submitter"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
)

type Client struct {
	config *models.ClientConfig
	encoder encoder.Encoder
	requester requester.Requester
	submitter submitter.Submitter
}

func NewClient(p userinput.CmdLineConfigProvider, e *encoder.Hasher, r *requester.PasswordRequester, s *submitter.HashSubmitter) Client {
	c := p.GetClientConfig()
	return Client{c, e, r, s}
}

func (c Client) Start() {
	log.Println("Starting Client...")
	go func() {
		err := c.requester.Start()
		if err != nil {
			log.Println(err)
		}
		c.Stop()
	}()
	go func() {
		err := c.encoder.Start()
		if err != nil {
			log.Println(err)
		}
		c.Stop()
	}()
	go func() {
		err := c.submitter.Start()
		if err != nil {
			log.Println(err)
		}
		c.Stop()
	}()

	for {}
}

func (c Client) Stop() {
	log.Println("Stopping Client...")
}
