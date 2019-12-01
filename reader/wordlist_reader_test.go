package reader

import (
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"os"
	"testing"
)

func TestNewWordlistReader(t *testing.T) {
	configProvider := setupConfigProvider()
	mockQueue := mocks.MockQueue{}
	NewWordlistReader(&configProvider, &mockQueue)
	assertConfigProviderCalled(t, configProvider)
}

func TestWordlistReader_LoadPasswords_Success(t *testing.T) {
	testPath := "wordlist_test.txt"
	passwords := []string{"password1"}
	setupFile(testPath, passwords)

	config := models.Config{WordlistPath: testPath}
	queue := mocks.MockQueue{PutCalls: 0}
	reader := WordlistReader{config: &config, passwords: &queue}

	err := reader.LoadPasswords()

	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}

	os.Remove(testPath)
}

func TestWordlistReader_LoadPasswords_Error(t *testing.T) {
	testPath := "wordlist_test.txt"

	config := models.Config{WordlistPath: testPath}
	queue := mocks.MockQueue{PutCalls: 0}
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

	config := models.Config{WordlistPath: testPath}
	queue := mocks.MockQueue{PutCalls: 0}
	reader := WordlistReader{config: &config, passwords: &queue}

	reader.LoadPasswords()

	actual := queue.PutCalls
	if actual != expected {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}

	os.Remove(testPath)
}

func TestWordlistReader_LoadPasswords_MultiplePutCalls(t *testing.T) {
	expected := uint64(2)

	testPath := "wordlist_test.txt"
	passwords := []string{"password1", "password2"}
	setupFile(testPath, passwords)

	config := models.Config{WordlistPath: testPath}
	queue := mocks.MockQueue{PutCalls: 0}
	reader := WordlistReader{config: &config, passwords: &queue}

	reader.LoadPasswords()

	actual := queue.PutCalls
	if actual != expected {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}

	os.Remove(testPath)
}
