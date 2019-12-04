package encoder

import (
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"sync"
	"testing"
)

type factoryTestObject struct {
	config        *models.Config
	hasherFactory *HasherFactory
}

func setupConfig() models.Config {
	return models.Config{
		LogPath: "test",
	}
}

func setupConfigProvider() mocks.MockConfigProvider {
	config := setupConfig()
	return mocks.NewMockConfigProvider(&config)
}

func setupHashFactory() factoryTestObject {
	config := setupConfig()
	HasherFactory := HasherFactory{
		config:          &config,
		logger:          &mocks.MockLogger{},
		mux:             &sync.Mutex{},
		requestQueue:    &mocks.MockRequestQueue{},
		stopQueue:       &mocks.MockClientStopQueue{},
		submissionQueue: &mocks.MockSubmissionQueue{},
		waiter:          &mocks.MockWaiter{},
	}
	return factoryTestObject{
		config:         &config,
		hasherFactory:  &HasherFactory,
	}
}

func assertConfigProviderCalled(t *testing.T, m *mocks.MockConfigProvider) {
	expected := uint64(1)

	actual := m.GetConfigCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestNewHasherFactory(t *testing.T) {
	configProvider := setupConfigProvider()
	NewHasherFactory(&configProvider, nil, nil, nil, nil, nil)
	assertConfigProviderCalled(t, &configProvider)
}

