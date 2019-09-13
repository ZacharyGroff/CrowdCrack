package queue

import (
	"testing"
	"crypto/sha256"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

func TestPutHashingRequestSuccess(t *testing.T) {
	request := models.HashingRequest{sha256.New(), 5}
	q := NewHashingRequestQueue()

	err := q.Put(request)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPutHashingRequestError(t *testing.T) {
	request := models.HashingRequest{sha256.New(), 5}
	q := NewHashingRequestQueue()

	q.Put(request)
	q.Put(request)
	err := q.Put(request)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestGetHashingRequestSuccess(t *testing.T) {
	expected := models.HashingRequest{sha256.New(), 5}
	q := NewHashingRequestQueue()

	q.Put(expected)
	actual, _ := q.Get()

	if expected != actual {
		t.Errorf("Expected: %+v\nActual: %+v\n", expected, actual)
	}
}

func TestGetHashingRequestError(t *testing.T) {
	q := NewHashingRequestQueue()
	
	_, err := q.Get()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestSizeZeroHashingRequests(t *testing.T) {
	expected := 0	
	q := NewHashingRequestQueue()
	actual := q.Size()

	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestSizeNotZeroHashingRequests(t *testing.T) {
	expected := 2

	request := models.HashingRequest{sha256.New(), 5}
	q := NewHashingRequestQueue()
	q.Put(request)
	q.Put(request)

	actual := q.Size()
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)	
	}
}