package flusher

import (
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type ClientQueueFlusher struct {
	config          *models.Config
	logger          interfaces.Logger
	requestQueue    interfaces.RequestQueue
	submissionQueue interfaces.SubmissionQueue
}

func NewClientQueueFlusher(p interfaces.ConfigProvider, l interfaces.Logger, r interfaces.RequestQueue, s interfaces.SubmissionQueue) interfaces.Flusher {
	return &ClientQueueFlusher{
		config:          p.GetConfig(),
		logger:          l,
		requestQueue:    r,
		submissionQueue: s,
	}
}

func (c *ClientQueueFlusher) NeedsFlushed() bool {
	return c.requestQueue.Size() > 0  || c.submissionQueue.Size() > 0
}

func (c *ClientQueueFlusher) Flush() error {
	panic("implement me")
}

