package mocks

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type MockConfigProvider struct {
	GetConfigCalls uint64
	configToReturn *models.Config
}

func NewMockConfigProvider(c *models.Config) MockConfigProvider {
	return MockConfigProvider{
		GetConfigCalls: 0,
		configToReturn: c,
	}
}

func (m *MockConfigProvider) GetConfig() *models.Config {
	m.GetConfigCalls++
	return m.configToReturn
}

