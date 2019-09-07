package queue

import (
	"os"
	"strings"
	"testing"
	"github.com/ZacharyGroff/CrowdCrack/config"
)

func TestPutHashSuccess(t *testing.T) {
	config := config.ServerConfig{HashQueueBuffer: 1}	
	q := NewHashQueue(&config)
	hash := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	err := q.Put(hash)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPutHashError(t *testing.T) {
	config := config.ServerConfig{HashQueueBuffer: 0}	
	q := NewHashQueue(&config)
	hash := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	err := q.Put(hash)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestGetHashSuccess(t *testing.T) {
	expected := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	config := config.ServerConfig{HashQueueBuffer: 1}	
	q := NewHashQueue(&config)
	q.Put(expected)

	actual, _ := q.Get()
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %x\nActual: %x\n", expected, actual)
	}
}

func TestGetHashError(t *testing.T) {
	config := config.ServerConfig{HashQueueBuffer: 0}
	q := NewHashQueue(&config)

	_, err := q.Get()
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestFlushSize(t *testing.T) {
	testPath := "hash_test.txt"
	os.Create(testPath)

	config := config.ServerConfig{ComputedHashOverflowPath: testPath, HashQueueBuffer: 1, FlushToFile: false}
	q := NewHashQueue(&config)

	hash := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	q.Put(hash)
	q.Flush()
	
	expected := 0
	actual := len(q.hashes)
	if expected != actual {
		os.Remove(testPath)
		t.Errorf("Expected: %x\tActual: %x\n", expected, actual)	
	}
}

func TestFlushToFileSuccess(t *testing.T) {
	testPath := "hash_test.txt"
	f, err := os.Create(testPath)
	f.Close()

	config := config.ServerConfig{ComputedHashOverflowPath: testPath, HashQueueBuffer: 1, FlushToFile: true}
	q := NewHashQueue(&config)

	hash := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	q.Put(hash)
	err = q.Flush()
	
	if err != nil {
		os.Remove(testPath)
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}

	os.Remove(testPath)
}

func TestFlushToFileError(t *testing.T) {
	testPath := "hash_test.txt"

	config := config.ServerConfig{ComputedHashOverflowPath: testPath, HashQueueBuffer: 1, FlushToFile: true}
	q := NewHashQueue(&config)

	hash := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	q.Put(hash)
	err := q.Flush()
	
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestFlushWithoutFileSuccess(t *testing.T) {
	config := config.ServerConfig{HashQueueBuffer: 1, FlushToFile: false}
	q := NewHashQueue(&config)

	hash := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	q.Put(hash)
	err := q.Flush()
	
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestSizeZeroHashes(t *testing.T) {
	expected := 0

	config := config.ServerConfig{HashQueueBuffer: 5}
	q := NewHashQueue(&config)
	actual := q.Size()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestSizeNotZeroHashes(t *testing.T) {
	expected := 2

	config := config.ServerConfig{HashQueueBuffer: 5}
	q := NewHashQueue(&config)

	hash := "2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED"
	q.Put(hash)
	q.Put(hash)
	actual := q.Size()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}
