package encoder

import (
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
)

type HasherFactory struct {
	config *models.Config
	logger logger.Logger
	requestQueue queue.RequestQueue
	submissionQueue queue.SubmissionQueue
	waiter waiter.Waiter
}

func NewHasherFactory(p userinput.CmdLineConfigProvider, l *logger.ConcurrentLogger, r *queue.HashingRequestQueue, s *queue.HashingSubmissionQueue, w waiter.Sleeper) *HasherFactory {
	return &HasherFactory {
		config:          p.GetConfig(),
		logger:          l,
		requestQueue:    r,
		submissionQueue: s,
		waiter:          w,
	}
}

func (h *HasherFactory) GetNewEncoder() Encoder {
	return &Hasher {
		config:          h.config,
		logger:          h.logger,
		requestQueue:    h.requestQueue,
		submissionQueue: h.submissionQueue,
		waiter:          h.waiter,
	}
}