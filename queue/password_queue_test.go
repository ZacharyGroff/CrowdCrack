package queue

import (
	"testing"
)

func TestNewPasswordQueue(t *testing.T) {
	configProvider := setupConfigProvider()
	NewPasswordQueue(&configProvider)
	assertConfigProviderCalled(t, configProvider)
}

func TestPasswordQueue_Put_Success(t *testing.T) {
	q := PasswordQueue{passwords: make(chan string, 1)}
	password := "hunter2"
	err := q.Put(password)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPasswordQueue_Put_Error(t *testing.T) {
	q := PasswordQueue{passwords: make(chan string, 0)}
	password := "hunter2"
	err := q.Put(password)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestPasswordQueue_Get_Success(t *testing.T) {
	expected := "hunter2"
	q := PasswordQueue{passwords: make(chan string, 1)}
	q.Put(expected)

	actual, _ := q.Get()
	if expected != actual {
		t.Errorf("Expected: %q\nActual: %q\n", expected, actual)
	}
}

func TestPasswordQueue_Get_Error(t *testing.T) {
	q := PasswordQueue{passwords: make(chan string, 0)}

	_, err := q.Get()
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestPasswordQueue_Size_Zero(t *testing.T) {
	expected := 0

	q := PasswordQueue{passwords: make(chan string, 5)}
	actual := q.Size()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestPasswordQueue_Size_NotZero(t *testing.T) {
	expected := 2

	q := PasswordQueue{passwords: make(chan string, 5)}
	password := "hunter2"

	q.Put(password)
	q.Put(password)
	actual := q.Size()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}
