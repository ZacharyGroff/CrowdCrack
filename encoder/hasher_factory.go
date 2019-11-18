package encoder

import (
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
)

type HasherFactory struct {
	config          *models.Config
	logger          interfaces.Logger
	requestQueue    interfaces.RequestQueue
	stopQueue       interfaces.ClientStopQueue
	submissionQueue interfaces.SubmissionQueue
	waiter          interfaces.Waiter
}

func NewHasherFactory(p interfaces.ConfigProvider, l *logger.ConcurrentLogger, r *queue.HashingRequestQueue, s *queue.HashingSubmissionQueue, c *queue.ClientStopReasonQueue, w waiter.Sleeper) *HasherFactory {
	return &HasherFactory{
		config:          p.GetConfig(),
		logger:          l,
		requestQueue:    r,
		stopQueue:       c,
		submissionQueue: s,
		waiter:          w,
	}
}

func (h *HasherFactory) GetNewEncoder() interfaces.Encoder {
	return NewHasher(h.config, h.logger, h.requestQueue, h.submissionQueue, h.stopQueue, h.waiter)
}
