package encoder

import (
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"sync"
)

type HasherFactory struct {
	config          *models.Config
	logger          interfaces.Logger
	mux             *sync.Mutex
	requestQueue    interfaces.RequestQueue
	stopQueue       interfaces.ClientStopQueue
	submissionQueue interfaces.SubmissionQueue
	waiter          interfaces.Waiter
}

func NewHasherFactory(p interfaces.ConfigProvider, l interfaces.Logger, r interfaces.RequestQueue, s interfaces.SubmissionQueue, c interfaces.ClientStopQueue, w interfaces.Waiter) interfaces.EncoderFactory {
	m := new(sync.Mutex)
	return &HasherFactory{
		config:          p.GetConfig(),
		logger:          l,
		mux:             m,
		requestQueue:    r,
		stopQueue:       c,
		submissionQueue: s,
		waiter:          w,
	}
}

func (h *HasherFactory) GetNewEncoder() interfaces.Encoder {
	return NewHasher(h.config, h.logger, h.requestQueue, h.submissionQueue, h.stopQueue, h.waiter, h.mux)
}
