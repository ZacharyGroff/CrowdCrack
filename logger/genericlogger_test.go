package logger

import (
	"bufio"
	"os"
	"strings"
	"testing"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

func TestGenericLogger_LogMessage_Error(t *testing.T) {
	config := models.Config{LogPath: ""}
	GenericLogger := GenericLogger{&config}
	err := GenericLogger.LogMessage("test")
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestGenericLogger_logToFile_Success(t *testing.T) {
	logPath := "test_log.txt"
	config := models.Config{LogPath: logPath}
	GenericLogger := GenericLogger{&config}
	err := GenericLogger.logToFile("test")
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}

	os.Remove(logPath)
}

func TestGenericLogger_logToFile_CorrectWrite(t *testing.T) {
	expected := "test"

	logPath := "test_log.txt"
	config := models.Config{LogPath: logPath}
	GenericLogger := GenericLogger{&config}
	GenericLogger.logToFile(expected)

	f, _ := os.Open(logPath)
	reader := bufio.NewReader(f)
	line, _, _ := reader.ReadLine()

	actual := string(line)
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}

	os.Remove(logPath)
}

func TestGenericLogger_logToFile_Error(t *testing.T) {
	config := models.Config{LogPath: ""}
	GenericLogger := GenericLogger{&config}
	err := GenericLogger.logToFile("test")
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}
