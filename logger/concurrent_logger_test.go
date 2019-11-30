package logger

import (
	"bufio"
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"os"
	"strings"
	"testing"
	"time"
)

type testObject struct{
	configProvider *mocks.MockConfigProvider
	logger interfaces.Logger
}

func assertConfigProviderCalled(t *testing.T, testObject testObject) {
	expected := uint64(1)

	actual := testObject.configProvider.GetConfigCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestNewConcurrentLogger(t *testing.T) {
	config := models.Config{LogPath: "testPath"}
	configProvider := mocks.NewMockConfigProvider(&config)
	logger := NewConcurrentLogger(&configProvider)

	testObject := testObject{
		configProvider: &configProvider,
		logger:         logger,
	}

	assertConfigProviderCalled(t, testObject)
}

func TestConcurrentLogger_LogMessage_Error(t *testing.T) {
	config := models.Config{LogPath: ""}
	ConcurrentLogger := ConcurrentLogger{config: &config}
	err := ConcurrentLogger.LogMessage("test")
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestConcurrentLogger_logToFile_Success(t *testing.T) {
	logPath := "test_log.txt"
	config := models.Config{LogPath: logPath}
	ConcurrentLogger := ConcurrentLogger{config: &config}
	err := ConcurrentLogger.logToFile("test")
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}

	os.Remove(logPath)
}

func TestConcurrentLogger_logToFile_CorrectWrite(t *testing.T) {
	expected := "test"

	logPath := "test_log.txt"
	config := models.Config{LogPath: logPath}
	ConcurrentLogger := ConcurrentLogger{config: &config}
	ConcurrentLogger.logToFile(expected)

	f, _ := os.Open(logPath)
	reader := bufio.NewReader(f)
	line, _, _ := reader.ReadLine()

	actual := string(line)
	if !strings.Contains(actual, expected) {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}

	os.Remove(logPath)
}

func TestConcurrentLogger_logToFile_Error(t *testing.T) {
	config := models.Config{LogPath: ""}
	ConcurrentLogger := ConcurrentLogger{config: &config}
	err := ConcurrentLogger.logToFile("test")
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
