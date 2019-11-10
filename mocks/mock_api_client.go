package mocks

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type MockApiClient struct {
	GetHashNameCalls               uint64
	GetPasswordsCalls              uint64
	SubmitHashesCalls              uint64
	getHashNameStatusCodeToReturn  int
	getPasswordsStatusCodeToReturn int
	submitHashesStatusCodeToReturn int
	hashNameToReturn               string
	passwordsToReturn              []string
}

func NewMockApiClient(getHashNameStatusCodeToReturn int, getPasswordsStatusCodeToReturn int, submitHashesStatusCodeToReturn int, hashNameToReturn string, passwordsToReturn []string) MockApiClient {
	return MockApiClient{0, 0, 0, getHashNameStatusCodeToReturn, getPasswordsStatusCodeToReturn, submitHashesStatusCodeToReturn, hashNameToReturn, passwordsToReturn}
}

func (m *MockApiClient) GetHashName() (int, string) {
	m.GetHashNameCalls++
	return m.getHashNameStatusCodeToReturn, m.hashNameToReturn
}

func (m *MockApiClient) GetPasswords(int) (int, []string) {
	m.GetPasswordsCalls++
	return m.getPasswordsStatusCodeToReturn, m.passwordsToReturn
}

func (m *MockApiClient) SubmitHashes(models.HashSubmission) int {
	m.SubmitHashesCalls++
	return m.submitHashesStatusCodeToReturn
}
