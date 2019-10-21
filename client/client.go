package client

import (
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/encoder"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/requester"
	"github.com/ZacharyGroff/CrowdCrack/submitter"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
)

type Client struct {
	config *models.Config
	encoder encoder.Encoder
	logger logger.Logger
	requester requester.Requester
	submitter submitter.Submitter
}

func NewClient(p userinput.CmdLineConfigProvider, e *encoder.Hasher, l *logger.GenericLogger, r *requester.PasswordRequester, s *submitter.HashSubmitter) Client {
	c := p.GetConfig()
	return Client{
		config:    c,
		encoder:   e,
		logger:    l,
		requester: r,
		submitter: s,
	}
}

func (c Client) Start() {
	c.logger.LogMessage("Starting Client...")
	go func() {
		err := c.requester.Start()
		if err != nil {
			c.logger.LogMessage(err.Error())
		}
		c.Stop()
	}()
	go func() {
		err := c.encoder.Start()
		if err != nil {
			c.logger.LogMessage(err.Error())
		}
		c.Stop()
	}()
	go func() {
		err := c.submitter.Start()
		if err != nil {
			c.logger.LogMessage(err.Error())
		}
		c.Stop()
	}()

	for {}
}

func (c Client) Stop() {
	c.logger.LogMessage("Stopping Client...")
}
