package queue

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
	"os"
	"strings"
	"testing"
)

func TestNewHashQueue(t *testing.T) {
	configProvider := setupConfigProvider()
	NewHashQueue(&configProvider)
	assertConfigProviderCalled(t, configProvider)
}

func TestHashQueue_PutHash_Success(t *testing.T) {
	q := HashQueue{hashes: make(chan string, 1)}
	hash := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	err := q.Put(hash)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHashQueue_PutHash_Error(t *testing.T) {
	q := HashQueue{hashes: make(chan string, 0)}
	hash := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	err := q.Put(hash)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestHashQueue_GetHash_Success(t *testing.T) {
	expected := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	q := HashQueue{hashes: make(chan string, 1)}
	q.Put(expected)

	actual, _ := q.Get()
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %x\nActual: %x\n", expected, actual)
	}
}

func TestHashQueue_GetHash_Error(t *testing.T) {
	q := HashQueue{hashes: make(chan string, 0)}

	_, err := q.Get()
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestHashQueue_Flush_Size(t *testing.T) {
	testPath := "hash_test.txt"
	config := models.Config{ComputedHashOverflowPath: testPath, HashQueueBuffer: 1, FlushToFile: false}
	q := HashQueue{hashes: make(chan string, 1), config: config}

	hash := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	q.Put(hash)
	q.Flush()

	expected := 0
	actual := len(q.hashes)
	if expected != actual {
		t.Errorf("Expected: %x\tActual: %x\n", expected, actual)
	}

	os.Remove(testPath)
}

func TestHashQueue_Flush_WithoutFileSuccess(t *testing.T) {
	config := models.Config{HashQueueBuffer: 1, FlushToFile: false}
	q := HashQueue{hashes: make(chan string, 1), config: config}

	hash := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	q.Put(hash)
	err := q.Flush()

	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHashQueue_FlushToFile_Success(t *testing.T) {
	testPath := "hash_test.txt"

	config := models.Config{ComputedHashOverflowPath: testPath, HashQueueBuffer: 1, FlushToFile: true}
	q := HashQueue{hashes: make(chan string, 1), config: config}

	hash := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	q.Put(hash)
	err := q.flushToFile()

	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}

	os.Remove(testPath)
}

func TestHashQueue_FlushToFile_OpenFileError(t *testing.T) {
	testPath := ""

	config := models.Config{ComputedHashOverflowPath: testPath, HashQueueBuffer: 1, FlushToFile: true}
	q := HashQueue{hashes: make(chan string, 1), config: config}

	hash := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	q.Put(hash)
	err := q.flushToFile()

	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestHashQueue_Size_Zero(t *testing.T) {
	expected := 0

	q := HashQueue{hashes: make(chan string, 5)}
	actual := q.Size()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestHashQueue_Size_NotZero(t *testing.T) {
	expected := 2
	q := HashQueue{hashes: make(chan string, 5)}

	hash := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	q.Put(hash)
	q.Put(hash)
	actual := q.Size()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}
