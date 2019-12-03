package logger

import (
	"bufio"
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

type testObject struct{
	config         *models.Config
	configProvider *mocks.MockConfigProvider
	logger         ConcurrentLogger
}

func setupConcurrentLoggerForSuccess() testObject {
	config := models.Config{LogPath: "test_log.txt"}
	configProvider := mocks.NewMockConfigProvider(&config)
	logger := ConcurrentLogger{
		config: &config,
		mux:    sync.Mutex{},
	}

	return testObject{
		config:         &config,
		configProvider: &configProvider,
		logger:         logger,
	}
}

func setupConcurrentLoggerForInvalidLogPath() testObject {
	config := models.Config{LogPath: ""}
	configProvider := mocks.NewMockConfigProvider(&config)
	logger := ConcurrentLogger{
		config: &config,
		mux:    sync.Mutex{},
	}

	return testObject{
		config:         &config,
		configProvider: &configProvider,
		logger:         logger,
	}
}

func assertConfigProviderCalled(t *testing.T, configProvider mocks.MockConfigProvider) {
	expected := uint64(1)

	actual := configProvider.GetConfigCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestNewConcurrentLogger(t *testing.T) {
	config := models.Config{LogPath: "test_log.txt"}
	configProvider := mocks.NewMockConfigProvider(&config)
	NewConcurrentLogger(&configProvider)

	assertConfigProviderCalled(t, configProvider)
}

func TestConcurrentLogger_LogMessage_Error(t *testing.T) {
	testObject := setupConcurrentLoggerForInvalidLogPath()

	err := testObject.logger.LogMessage("test")
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestConcurrentLogger_logToFile_Success(t *testing.T) {
	testObject := setupConcurrentLoggerForSuccess()

	err := testObject.logger.logToFile("test")
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}

	os.Remove(testObject.config.LogPath)
}

func TestConcurrentLogger_logToFile_CorrectWrite(t *testing.T) {
	expected := "test"

	testObject := setupConcurrentLoggerForSuccess()
	testObject.logger.logToFile(expected)

	f, _ := os.Open(testObject.config.LogPath)
	reader := bufio.NewReader(f)
	line, _, _ := reader.ReadLine()

	actual := string(line)
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}

	os.Remove(testObject.config.LogPath)
}

func TestConcurrentLogger_logToFile_Error(t *testing.T) {
	testObject := setupConcurrentLoggerForInvalidLogPath()
	err := testObject.logger.logToFile("test")
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestConcurrentLogger_getTimeFormattedMessage_CorrectResult(t *testing.T) {
	testMessage := "testMessage"
	currentTime := time.Now()
	expected := fmt.Sprintf("%s: %s", currentTime.Format(time.RFC822), testMessage)

	actual := getTimeFormattedMessage(currentTime, testMessage)
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}
