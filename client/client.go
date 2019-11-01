package client

import (
	"sync"
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

func NewClient(p userinput.CmdLineConfigProvider, e *encoder.Hasher, l *logger.ConcurrentLogger, r *requester.PasswordRequester, s *submitter.HashSubmitter) Client {
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

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		err := c.requester.Start()
		if err != nil {
			c.logger.LogMessage(err.Error())
		}
		c.logger.LogMessage("Requester Done!")
		wg.Done()
	}()
	go func() {
		err := c.encoder.Start()
		if err != nil {
			c.logger.LogMessage(err.Error())
		}
		c.logger.LogMessage("Encoder Done!")
		wg.Done()
	}()
	go func() {
		err := c.submitter.Start()
		if err != nil {
			c.logger.LogMessage(err.Error())
		}
		c.logger.LogMessage("Submitter Done!")
		wg.Done()
	}()

	wg.Wait()
}

func (c Client) Stop() {
	c.logger.LogMessage("Stopping Client...")
}
