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
