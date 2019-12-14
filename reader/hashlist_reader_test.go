package reader

import (
	"testing"
)

var hashes = []string{
	"f52fbd32b2b3b86ff88ef6c490628285f482af15ddcb29541f94bcf526a3f6c7",
	"ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f",
	"4339b2e3c470e9822e1c4f1caa6b6f3ef044d3701e35df7ff9735470e9aa014c",
}

func setupHashlistReader() HashlistReader {
	config := setupConfig()
	return HashlistReader{
		config: &config,
	}
}

func TestNewHashlistReader(t *testing.T) {
	configProvider := setupConfigProvider()
	NewHashlistReader(&configProvider)
	assertConfigProviderCalled(t, configProvider)
}

func TestHashlistReader_GetHashes_Success(t *testing.T) {
	setupFile(hashlistPath, hashes)
	defer cleanupFile(hashlistPath)

	reader := setupHashlistReader()

	_, err := reader.GetHashes()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHashlistReader_GetHashes_Error(t *testing.T) {
	reader := setupHashlistReader()

	_, err := reader.GetHashes()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHashlistReader_GetHashes_CorrectResults(t *testing.T) {
	expected := hashes

	setupFile(hashlistPath, expected)
	defer cleanupFile(hashlistPath)

	reader := setupHashlistReader()

	actual, _ := reader.GetHashes()
	for _, hash := range expected {
		if !actual[hash] {
			t.Errorf("Expected: %s to be in map: %+v\n", hash, actual)
		}
	}
}
