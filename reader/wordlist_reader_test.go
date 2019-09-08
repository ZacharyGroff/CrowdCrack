package reader

import (
	"os"
	"testing"
	"github.com/ZacharyGroff/CrowdCrack/config"
)

func TestLoadPasswordsNoError(t *testing.T) {
	testPath := "wordlist_test.txt"
	passwords := []string{"password1"}
	setupFile(testPath, passwords)

	config := config.ServerConfig{WordlistPath: testPath}
	queue := mockQueue{PutCalls: 0}
	reader := WordlistReader{config: &config, passwords: &queue}

	err := reader.LoadPasswords()
	
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
	
	os.Remove(testPath)
}

func TestLoadPasswordsError(t *testing.T) {
	testPath := "wordlist_test.txt"

	config := config.ServerConfig{WordlistPath: testPath}
	queue := mockQueue{PutCalls: 0}
	reader := WordlistReader{config: &config, passwords: &queue}

	err := reader.LoadPasswords()
	
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestLoadPasswordsPutNoCalls(t *testing.T) {
	expected := uint64(0)

	testPath := "wordlist_test.txt"
	passwords := []string{}
	setupFile(testPath, passwords)

	config := config.ServerConfig{WordlistPath: testPath}
	queue := mockQueue{PutCalls: 0}
	reader := WordlistReader{config: &config, passwords: &queue}

	reader.LoadPasswords()
	
	actual := queue.PutCalls
	if actual != expected {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}

	os.Remove(testPath)
}

func TestLoadPasswordsPutMultipleCalls(t *testing.T) {
	expected := uint64(2)

	testPath := "wordlist_test.txt"
	passwords := []string{"password1", "password2"}
	setupFile(testPath, passwords)

	config := config.ServerConfig{WordlistPath: testPath}
	queue := mockQueue{PutCalls: 0}
	reader := WordlistReader{config: &config, passwords: &queue}

	reader.LoadPasswords()
	
	actual := queue.PutCalls
	if actual != expected {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}

	os.Remove(testPath)
}
