package mocks

import (
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
)

type MockEncoderFactory struct {
	GetNewEncoderCalls uint64
	encoderToReturn    interfaces.Encoder
}

func NewMockEncoderFactory(e interfaces.Encoder) MockEncoderFactory {
	return MockEncoderFactory{
		GetNewEncoderCalls: 0,
		encoderToReturn:    e,
	}
}

func (m *MockEncoderFactory) GetNewEncoder() interfaces.Encoder {
	m.GetNewEncoderCalls++
	return m.encoderToReturn
}

