package reader

import (
	"bufio"
	"os"
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

func setupFile(testPath string, lines []string) {
	file, _ := os.Create(testPath)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		writer.WriteString(line)
		writer.WriteString("\n")
	}
	writer.Flush()
}
