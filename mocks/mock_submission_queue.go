package mocks

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type MockSubmissionQueue struct {
	GetCalls               uint64
	PutCalls               uint64
	SizeCalls              uint64
	errorToReturn          error
	hashSubmissionToReturn models.HashSubmission
	sizeToReturn           int
}

func NewMockSubmissionQueue(e error, h models.HashSubmission, s int) MockSubmissionQueue {
	return MockSubmissionQueue{0, 0, 0, e, h, s}
}

func (m *MockSubmissionQueue) Size() int {
	m.SizeCalls++
	return m.sizeToReturn
}

func (m *MockSubmissionQueue) Get() (models.HashSubmission, error) {
	m.GetCalls++
	return m.hashSubmissionToReturn, m.errorToReturn
}

func (m *MockSubmissionQueue) Put(h models.HashSubmission) error {
	m.PutCalls++
	return m.errorToReturn
}
