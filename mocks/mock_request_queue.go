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
	intToReturn int
}

func NewMockRequestQueue(e error, h models.HashingRequest, i int) MockRequestQueue {
	return MockRequestQueue{0, 0, 0, e, h, i}
}

func (m *MockRequestQueue) Size() int {
	m.SizeCalls++
	return m.intToReturn
}

func (m *MockRequestQueue) Get() (models.HashingRequest, error) {
	m.GetCalls++
	return m.hashingRequestToReturn, m.errorToReturn
}

func (m *MockRequestQueue) Put(h models.HashingRequest) error {
	m.PutCalls++
	return m.errorToReturn
}
