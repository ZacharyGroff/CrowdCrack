package reader

import (
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"testing"
)

type wordlistReaderTestObject struct {
	queue  *mocks.MockQueue
	reader WordlistReader
}

func setupQueue() mocks.MockQueue {
	return mocks.MockQueue{
		PutCalls: 0,
	}
}

func setupWordlistReader() wordlistReaderTestObject {
	config := setupConfig()
	queue := setupQueue()
	reader := WordlistReader{
		config:    &config,
		passwords: &queue,
	}

	return wordlistReaderTestObject{
		queue:  &queue,
		reader: reader,
	}
}

func assertQueuePutNotCalled(t *testing.T, m *mocks.MockQueue) {
	expected := uint64(0)
	actual := m.PutCalls
	if actual != expected {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertQueuePutCalledNTimes(t *testing.T, m *mocks.MockQueue, expected uint64) {
	actual := m.PutCalls
	if actual != expected {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestNewWordlistReader(t *testing.T) {
	configProvider := setupConfigProvider()
	mockQueue := mocks.MockQueue{}
	NewWordlistReader(&configProvider, &mockQueue)
	assertConfigProviderCalled(t, configProvider)
}

func TestWordlistReader_LoadPasswords_Success(t *testing.T) {
	passwords := []string{"password1"}

	setupFile(wordlistPath, passwords)
	defer cleanupFile(wordlistPath)

	testObject := setupWordlistReader()

	err := testObject.reader.LoadPasswords()

	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestWordlistReader_LoadPasswords_Error(t *testing.T) {
	testObject := setupWordlistReader()

	err := testObject.reader.LoadPasswords()

	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestLoadPasswordsPutNoCalls(t *testing.T) {
	passwords := []string{}

	setupFile(wordlistPath, passwords)
	defer cleanupFile(wordlistPath)

	testObject := setupWordlistReader()
	testObject.reader.LoadPasswords()

	assertQueuePutNotCalled(t, testObject.queue)
}

func TestWordlistReader_LoadPasswords_MultiplePutCalls(t *testing.T) {
	expected := uint64(2)
	passwords := []string{"password1", "password2"}

	setupFile(wordlistPath, passwords)
	defer cleanupFile(wordlistPath)

	testObject := setupWordlistReader()
	testObject.reader.LoadPasswords()

	assertQueuePutCalledNTimes(t, testObject.queue, expected)
}
