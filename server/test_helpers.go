package server

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

type mockVerifier struct {
	verifyCalls uint64
}

func (m *mockVerifier) Verify() {
	m.verifyCalls++
}

type mockApi struct {
	handleRequestsCalls uint64
}

func (m *mockApi) HandleRequests() {
	m.handleRequestsCalls++
}
