package client

import (
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"sync"
)

type Client struct {
	config         *models.Config
	encoderFactory interfaces.EncoderFactory
	logger         interfaces.Logger
	requester      interfaces.Requester
	submitter      interfaces.Submitter
}

func NewClient(p interfaces.ConfigProvider, e interfaces.EncoderFactory, l interfaces.Logger, r interfaces.Requester, s interfaces.Submitter) Client {
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

	var wg sync.WaitGroup

	go c.startRequester(&wg)
	c.startEncoders(&wg)
	go c.startSubmitter(&wg)

	wg.Wait()
	c.Stop()
}

func (c Client) startRequester(wg *sync.WaitGroup) {
	wg.Add(1)
	err := c.requester.Start()
	if err != nil {
		c.logger.LogMessage(err.Error())
	}
	c.logger.LogMessage("Requester Done!")
	wg.Done()
}

func (c Client) startEncoders(wg *sync.WaitGroup) {
	var encoderNum uint16
	for encoderNum = 0; encoderNum < c.config.Threads - 2; encoderNum++ {
		wg.Add(1)
		go c.startEncoder(encoderNum, wg)
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
	wg.Add(1)
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
