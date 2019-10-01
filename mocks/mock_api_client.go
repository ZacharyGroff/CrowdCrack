package mocks

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type MockApiClient struct {
	GetHashNameCalls uint64
	GetPasswordsCalls uint64
	SubmitHashesCalls uint64
	statusCodeToReturn int
	hashNameToReturn string
	passwordsToReturn []string
}

func (m *MockApiClient) GetHashName() (int, string) {
	m.GetHashNameCalls++
	return m.statusCodeToReturn, m.hashNameToReturn
}

func (m *MockApiClient) GetPasswords(int) (int, []string) {
	m.GetPasswordsCalls++
	return m.statusCodeToReturn, m.passwordsToReturn
}

func (m *MockApiClient) SubmitHashes(models.HashSubmission) int {
	m.SubmitHashesCalls++
	return m.statusCodeToReturn
}

