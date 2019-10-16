package logger

import (
	"os"
	"testing"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

func TestServerLogger_LogMessage_Error(t *testing.T) {
	config := models.ServerConfig{LogPath: ""}
	serverLogger := NewServerLogger(&config)
	err := serverLogger.LogMessage("test")
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestServerLogger_logToFile_Success(t *testing.T) {
	logPath := "test_log.txt"
	config := models.ServerConfig{LogPath: logPath}
	serverLogger := NewServerLogger(&config)
	err := serverLogger.logToFile("test")
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}

	os.Remove(logPath)
}

func TestServerLogger_logToFile_Error(t *testing.T) {
	config := models.ServerConfig{LogPath: ""}
	serverLogger := NewServerLogger(&config)
	err := serverLogger.logToFile("test")
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}
