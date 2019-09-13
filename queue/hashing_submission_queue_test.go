package queue

import (
	"testing"
)

func TestPutHashingSubmissionSuccess(t *testing.T) {
	submission := uint64(42)
	q := NewHashingSubmissionQueue()

	err := q.Put(submission)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPutHashingSubmissionError(t *testing.T) {
	submission := uint64(42)
	q := NewHashingSubmissionQueue()

	q.Put(submission)	
	q.Put(submission)	
	err := q.Put(submission)	
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestGetHashingSubmissionSuccess(t *testing.T) {
	expected := uint64(42)
	q := NewHashingSubmissionQueue()

	q.Put(expected)
	actual, _ := q.Get()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestGetHashingSubmissionError(t *testing.T) {
	q := NewHashingSubmissionQueue()

	_, err := q.Get()
	if err == nil {
		t.Errorf("Expected error but nil returned")
	}
}
