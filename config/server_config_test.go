package config

import (
	"strings"
	"testing"
)

func TestParseWordlistPath(t *testing.T) {
	config := ServerConfig{}
	config.parseConfig("server_config_test.json")
	
	expected := "path/to/wordlist.txt"
	actual := config.WordlistPath
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestParseHashlistPath(t *testing.T) {
	config := ServerConfig{}
	config.parseConfig("server_config_test.json")
	
	expected := "path/to/hashlist.txt"
	actual := config.HashlistPath
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestParseHashFunction(t *testing.T) {
	config := ServerConfig{}
	config.parseConfig("server_config_test.json")
	
	expected := "testHash"
	actual := config.HashFunction
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestParseApiPort(t *testing.T) {
	config := ServerConfig{}
	config.parseConfig("server_config_test.json")
	
	expected := uint16(42)
	actual := config.ApiPort
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestParsePasswordQueueBuffer(t *testing.T) {
	config := ServerConfig{}
	config.parseConfig("server_config_test.json")
	
	expected := uint64(43)
	actual := config.PasswordQueueBuffer
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestParseHashQueueBuffer(t *testing.T) {
	config := ServerConfig{}
	config.parseConfig("server_config_test.json")
	
	expected := uint64(44)
	actual := config.HashQueueBuffer
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestParseFlushToFile(t *testing.T) {
	config := ServerConfig{}
	config.parseConfig("server_config_test.json")
	
	expected := true
	actual := config.FlushToFile
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}

func TestParseHashPath(t *testing.T) {
	config := ServerConfig{}
	config.parseConfig("server_config_test.json")
	
	expected := "test/path/to/hashFile"
	actual := config.HashPath
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}
