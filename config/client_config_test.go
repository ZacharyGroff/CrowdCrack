package config

import (
	"strings"
	"testing"
)

func TestParseServerAddress(t *testing.T) {
	config := ClientConfig{}
	config.parseConfig("client_config_test.json")
	
	expected := "192.168.0.1:42"
	actual := config.ServerAddress
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}
