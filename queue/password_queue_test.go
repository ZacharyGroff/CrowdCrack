package queue

import (
	"testing"
	"github.com/ZacharyGroff/CrowdCrack/config"
)

func TestPutSuccess(t *testing.T) {
	config := config.ServerConfig{PasswordQueueBuffer: 1}	
	q := NewPasswordQueue(&config)
	password := "hunter2"
	err := q.Put(password)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPutError(t *testing.T) {
	config := config.ServerConfig{PasswordQueueBuffer: 0}	
	q := NewPasswordQueue(&config)
	password := "hunter2"
	err := q.Put(password)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestGetSuccess(t *testing.T) {
	expected := "hunter2"
	config := config.ServerConfig{PasswordQueueBuffer: 1}	
	q := NewPasswordQueue(&config)
	q.Put(expected)

	actual, _ := q.Get()
	if expected != actual {
		t.Errorf("Expected: %q\nActual: %q\n", expected, actual)
	}
}

func TestGetError(t *testing.T) {
	config := config.ServerConfig{PasswordQueueBuffer: 0}
	q := NewPasswordQueue(&config)

	_, err := q.Get()
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestSizeZero(t *testing.T) {
	expected := 0

	config := config.ServerConfig{PasswordQueueBuffer: 5}
	q := NewPasswordQueue(&config)
	actual := q.Size()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestSizeNotZero(t *testing.T) {
	expected := 2

	config := config.ServerConfig{PasswordQueueBuffer: 5}
	q := NewPasswordQueue(&config)
	password := "hunter2"

	q.Put(password)
	q.Put(password)
	actual := q.Size()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}
