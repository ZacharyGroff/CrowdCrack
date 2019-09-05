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

func TestLoadPasswordsNoError(t *testing.T) {
	testPath := "wordlist_test.txt"
	file, _ := os.Create(testPath)
	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, "testpassword1") 
	fmt.Fprintln(writer, "testpassword2") 
	file.Close()
	
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
