package mocks

import (
	"github.com/ZacharyGroff/CrowdCrack/encoder"
)

type MockEncoderFactory struct {
	GetNewEncoderCalls uint64
	encoderToReturn encoder.Encoder
}

func (m *MockEncoderFactory) GetNewEncoder() encoder.Encoder {
	m.GetNewEncoderCalls++
	return m.encoderToReturn
}

