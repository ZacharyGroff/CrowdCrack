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
	encoderFactory encoder.EncoderFactory
	logger logger.Logger
	requester requester.Requester
	submitter submitter.Submitter
}

func NewClient(p userinput.CmdLineConfigProvider, e *encoder.HasherFactory, l *logger.ConcurrentLogger, r *requester.PasswordRequester, s *submitter.HashSubmitter) Client {
	c := p.GetConfig()
	return Client{
		config:           c,
		encoderFactory:   e,
		logger:           l,
		requester:        r,
		submitter:        s,
	}
}

func (c Client) Start() {
	c.logger.LogMessage("Starting Client...")

	var wg sync.WaitGroup
	wg.Add(3)

	go c.startRequester(&wg)
	go c.startEncoder(&wg)
	go c.startSubmitter(&wg)

	wg.Wait()
}

func (c Client) startRequester(wg *sync.WaitGroup) {
	err := c.requester.Start()
	if err != nil {
		c.logger.LogMessage(err.Error())
	}
	c.logger.LogMessage("Requester Done!")
	wg.Done()
}

func (c Client) startEncoder(wg *sync.WaitGroup) {
	encoder := c.encoderFactory.GetNewEncoder()
	err := encoder.Start()
	if err != nil {
		c.logger.LogMessage(err.Error())
	}
	c.logger.LogMessage("Encoder Done!")
	wg.Done()
}

func (c Client) startSubmitter(wg *sync.WaitGroup) {
	err := c.submitter.Start()
	if err != nil {
		c.logger.LogMessage(err.Error())
	}
	c.logger.LogMessage("Submitter Done!")
	wg.Done()
}

func (c Client) Stop() {
	c.logger.LogMessage("Stopping Client...")
}
