package queue

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
	"reflect"
	"testing"
)

func TestHashingSubmissionQueue_Size_Zero(t *testing.T) {
	expected := 0
	q := NewHashingSubmissionQueue()

	actual := q.Size()
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestHashingSubmissionQueue_Size_NotZero(t *testing.T) {
	expected := 2
	submission := models.HashSubmission{}
	q := NewHashingSubmissionQueue()
	q.Put(submission)
	q.Put(submission)

	actual := q.Size()
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestHashingSubmissionQueue_Put_Success(t *testing.T) {
	submission := models.HashSubmission{}
	q := NewHashingSubmissionQueue()

	err := q.Put(submission)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHashingSubmissionQueue_Put_Error(t *testing.T) {
	submission := models.HashSubmission{}
	q := NewHashingSubmissionQueue()

	for i := 0; i < 100; i++ {
		q.Put(submission)
	}

	err := q.Put(submission)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestHashingSubmissionQueue_Get_Success(t *testing.T) {
	expected := models.HashSubmission{"sha256", []string{"result1"}}
	q := NewHashingSubmissionQueue()

	q.Put(expected)
	actual, _ := q.Get()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", expected, actual)
	}
}

func TestHashingSubmissionQueue_Get_Error(t *testing.T) {
	q := NewHashingSubmissionQueue()

	_, err := q.Get()
	if err == nil {
		t.Errorf("Expected error but nil returned")
	}
}
