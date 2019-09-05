package reader

import (
	"bufio"
	"fmt"
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

func (m mockQueue) Put(password string) error {
	m.PutCalls++
	return nil
}

func setupWordlist(testPath string, passwords []string) {
	file, _ := os.Create(testPath)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, password := range passwords {
		fmt.Fprintln(writer, password) 
	}
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
		os.Remove(testPath)
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}

	os.Remove(testPath)
}
