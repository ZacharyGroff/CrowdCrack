package reader

import (
	"bufio"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"os"
	"testing"
)

var threads = uint16(3)

func setupConfig() models.Config {
	return models.Config{
		Threads: threads,
	}
}

func setupConfigProvider() mocks.MockConfigProvider {
	config := setupConfig()
	return mocks.NewMockConfigProvider(&config)
}

func assertConfigProviderCalled(t *testing.T, m mocks.MockConfigProvider) {
	expected := uint64(1)

	actual := m.GetConfigCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
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
