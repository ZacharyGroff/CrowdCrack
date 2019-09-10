package reader

import (
	"os"
	"testing"
	"github.com/ZacharyGroff/CrowdCrack/config"
)

func TestGetHashesNoError(t *testing.T) {
	testPath := "hashlist_test.txt"
	hashes := []string{
		"f52fbd32b2b3b86ff88ef6c490628285f482af15ddcb29541f94bcf526a3f6c7",
		"ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f",
		"4339b2e3c470e9822e1c4f1caa6b6f3ef044d3701e35df7ff9735470e9aa014c",
	}
	setupFile(testPath, hashes)

	config := config.ServerConfig{HashlistPath: testPath}
	reader := HashlistReader{&config}

	_, err := reader.GetHashes()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}

	os.Remove(testPath)
}

func TestGetHashesError(t *testing.T) {
	testPath := "hashlist_test.txt"

	config := config.ServerConfig{HashlistPath: testPath}
	reader := HashlistReader{&config}

	_, err := reader.GetHashes()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestGetHashesCorrectResults(t *testing.T) {
	testPath := "hashlist_test.txt"
	expected := []string{
		"f52fbd32b2b3b86ff88ef6c490628285f482af15ddcb29541f94bcf526a3f6c7",
		"ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f",
		"4339b2e3c470e9822e1c4f1caa6b6f3ef044d3701e35df7ff9735470e9aa014c",
	}
	setupFile(testPath, expected)

	config := config.ServerConfig{HashlistPath: testPath}
	reader := HashlistReader{&config}

	actual, _ := reader.GetHashes()
	for _, hash := range expected { 
		if !actual[hash] {
			t.Errorf("Expected: %s to be in map: %+v\n", hash, actual)
		}
	}

	os.Remove(testPath)
}