package encoder

import (
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
)

type HasherFactory struct {
	config          *models.Config
	logger          interfaces.Logger
	requestQueue    interfaces.RequestQueue
	submissionQueue interfaces.SubmissionQueue
	waiter          interfaces.Waiter
}

func NewHasherFactory(p userinput.CmdLineConfigProvider, l *logger.ConcurrentLogger, r *queue.HashingRequestQueue, s *queue.HashingSubmissionQueue, w waiter.Sleeper) *HasherFactory {
	return &HasherFactory{
		config:          p.GetConfig(),
		logger:          l,
		requestQueue:    r,
		submissionQueue: s,
		waiter:          w,
	}
}

func (h *HasherFactory) GetNewEncoder() interfaces.Encoder {
	return NewHasher(h.config, h.logger, h.requestQueue, h.submissionQueue, h.waiter)
}
