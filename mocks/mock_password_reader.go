package mocks

import (
	"errors"
)

type MockPasswordReader struct {
	LoadPasswordsCalls uint64
	isErrorCall        bool
}

func NewMockPasswordReader(isErrorCall bool) MockPasswordReader {
	return MockPasswordReader{0, isErrorCall}
}

func (m *MockPasswordReader) LoadPasswords() error {
	m.LoadPasswordsCalls++
	if m.isErrorCall {
		return errors.New("test error")
	}

	return nil
}
