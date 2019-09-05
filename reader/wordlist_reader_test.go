package reader

import (
	"bufio"
	"os"
	"testing"
	"github.com/ZacharyGroff/CrowdCrack/config"
)

type mockQueue struct {
	PutCalls uint64
}

func (m mockQueue) Size() int {
	return -1
}

func (m mockQueue) Get() (string, error) {
	return "", nil
}

func (m *mockQueue) Put(password string) error {
	m.PutCalls++
	return nil
}

func setupWordlist(testPath string, passwords []string) {
	file, _ := os.Create(testPath)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, password := range passwords {
		writer.WriteString(password)
		writer.WriteString("\n")
	}
	writer.Flush()
}

func TestLoadPasswordsNoError(t *testing.T) {
	testPath := "wordlist_test.txt"
	passwords := []string{"password1"}
	setupWordlist(testPath, passwords)

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
	setupWordlist(testPath, passwords)

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
	setupWordlist(testPath, passwords)

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
