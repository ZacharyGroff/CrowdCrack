package mocks

import (
	"errors"
)

type mockPasswordReader struct {
    loadPasswordsCalls uint64
    isErrorCall bool
}

func (m *mockPasswordReader) LoadPasswords() error {
    m.loadPasswordsCalls++
    if m.isErrorCall {
        return errors.New("test error")
    }

    return nil
}
