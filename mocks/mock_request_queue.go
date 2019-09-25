package mocks

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type MockRequestQueue struct {
	GetCalls uint64
	PutCalls uint64
	SizeCalls uint64
	errorToReturn error
	hashingRequestToReturn models.HashingRequest
}

func NewMockRequestQueue(e error, h models.HashingRequest) MockRequestQueue {
	return MockRequestQueue{0, 0, 0, e, h}
}

func (m *MockRequestQueue) Size() int {
	m.SizeCalls++
	return int(m.PutCalls) - int(m.GetCalls)
}

func (m *MockRequestQueue) Get() (models.HashingRequest, error) {
	m.GetCalls++
	return m.hashingRequestToReturn, m.errorToReturn
}

func (m *MockRequestQueue) Put(h models.HashingRequest) error {
	m.PutCalls++
	return m.errorToReturn
}
