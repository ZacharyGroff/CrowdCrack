package queue

import (
	"testing"
	"github.com/ZacharyGroff/CrowdCrack/config"
)

func TestPutPasswordSuccess(t *testing.T) {
	config := config.ServerConfig{PasswordQueueBuffer: 1}	
	q := NewServerPasswordQueue(&config)
	password := "hunter2"
	err := q.Put(password)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPutPasswordError(t *testing.T) {
	config := config.ServerConfig{PasswordQueueBuffer: 0}	
	q := NewServerPasswordQueue(&config)
	password := "hunter2"
	err := q.Put(password)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestGetPasswordSuccess(t *testing.T) {
	expected := "hunter2"
	config := config.ServerConfig{PasswordQueueBuffer: 1}	
	q := NewServerPasswordQueue(&config)
	q.Put(expected)

	actual, _ := q.Get()
	if expected != actual {
		t.Errorf("Expected: %q\nActual: %q\n", expected, actual)
	}
}

func TestGetPasswordError(t *testing.T) {
	config := config.ServerConfig{PasswordQueueBuffer: 0}
	q := NewServerPasswordQueue(&config)

	_, err := q.Get()
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestSizeZeroPasswords(t *testing.T) {
	expected := 0

	config := config.ServerConfig{PasswordQueueBuffer: 5}
	q := NewServerPasswordQueue(&config)
	actual := q.Size()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestSizeNotZeroPasswords(t *testing.T) {
	expected := 2

	config := config.ServerConfig{PasswordQueueBuffer: 5}
	q := NewServerPasswordQueue(&config)
	password := "hunter2"

	q.Put(password)
	q.Put(password)
	actual := q.Size()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}
