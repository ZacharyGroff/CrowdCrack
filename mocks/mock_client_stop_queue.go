package mocks

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type MockClientStopQueue struct {
	GetCalls                 uint64
	PutCalls                 uint64
	clientStopReasonToReturn models.ClientStopReason
	getErrorToReturn         error
	putErrorToReturn         error
}

func NewMockClientStopQueue(c models.ClientStopReason, g error, p error) MockClientStopQueue {
	return MockClientStopQueue{
		GetCalls:                 0,
		PutCalls:                 0,
		clientStopReasonToReturn: c,
		getErrorToReturn:         g,
		putErrorToReturn:         p,
	}
}

func (m *MockClientStopQueue) Get() (models.ClientStopReason, error) {
	m.GetCalls++
	return m.clientStopReasonToReturn, m.getErrorToReturn
}

func (m *MockClientStopQueue) Put(models.ClientStopReason) error {
	m.PutCalls++
	return m.putErrorToReturn
}
