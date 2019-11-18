package client

import (
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/encoder"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/requester"
	"github.com/ZacharyGroff/CrowdCrack/submitter"
	"sync"
)

type Client struct {
	config         *models.Config
	encoderFactory interfaces.EncoderFactory
	logger         interfaces.Logger
	requester      interfaces.Requester
	submitter      interfaces.Submitter
}

func NewClient(p interfaces.ConfigProvider, e *encoder.HasherFactory, l *logger.ConcurrentLogger, r *requester.PasswordRequester, s *submitter.HashSubmitter) Client {
	c := p.GetConfig()
	return Client{
		config:         c,
		encoderFactory: e,
		logger:         l,
		requester:      r,
		submitter:      s,
	}
}

func (c Client) Start() {
	c.logger.LogMessage("Starting Client...")

	availableThreads := int(c.config.Threads)
	var wg sync.WaitGroup
	wg.Add(availableThreads)

	go c.startRequester(&wg)
	go c.startEncoders(&wg)
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

func (c Client) startEncoders(wg *sync.WaitGroup) {
	var encoderNum uint16
	for encoderNum = 0; encoderNum < c.config.Threads; encoderNum++ {
		c.startEncoder(encoderNum, wg)
	}
}

func (c Client) startEncoder(encoderNum uint16, wg *sync.WaitGroup) {
	encoder := c.encoderFactory.GetNewEncoder()
	err := encoder.Start()
	if err != nil {
		c.logger.LogMessage(err.Error())
	}
	logMessage := fmt.Sprintf("Encoder #%d Done!", encoderNum)
	c.logger.LogMessage(logMessage)
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
