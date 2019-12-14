package reader

import (
	"bufio"
	"errors"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"os"
	"testing"
)

var nilError = error(nil)
var testError = errors.New("test error")

var hashlistPath = "hashlist_test.txt"
var requestBackupPath = "test_request_backup.json"
var submissionBackupPath = "test_submission_backup.json"
var wordlistPath = "wordlist_test.txt"
var threads = uint16(3)

func setupConfig() models.Config {
	return models.Config{
		HashlistPath: hashlistPath,
		RequestBackupPath: requestBackupPath,
		SubmissionBackupPath: submissionBackupPath,
		Threads: threads,
		WordlistPath: wordlistPath,
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

func setupFiles(testPaths []string, linesForFiles [][]string) {
	for i, _ := range testPaths {
		setupFile(testPaths[i], linesForFiles[i])
	}
}

func setupFile(testPath string, linesForFile []string) {
	file, _ := os.Create(testPath)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range linesForFile {
		writer.WriteString(line)
		writer.WriteString("\n")
	}
	writer.Flush()
}

func cleanupFiles(paths []string) {
	for _, path := range paths {
		cleanupFile(path)
	}
}

func cleanupFile(path string) {
	os.Remove(path)
}
