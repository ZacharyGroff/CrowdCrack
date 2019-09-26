package encoder

import (
	"errors"
	"reflect"
	"strings"
	"testing"
	"crypto/sha256"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

func TestHasherStartError(t *testing.T) {
	requestQueueError := error(nil)
	hashingRequest := models.HashingRequest {
		HashName: "fakeHashFunction",
		Passwords: []string {
			"password123",
		},
	}
	mockRequestQueue := mocks.NewMockRequestQueue(requestQueueError, hashingRequest)
	mockWaiter := mocks.MockWaiter{0}
	hasher := Hasher{waiter: &mockWaiter, requestQueue: &mockRequestQueue}

	err := hasher.Start()

	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHasherProcessOrSleepProcessSuccess(t *testing.T) {
	queueError := error(nil)
	hashingRequest := models.HashingRequest {
		HashName: "sha256",
		Passwords: []string {
			"password123",
		},
	}
	hashSubmission := models.HashSubmission{}
	mockRequestQueue := mocks.NewMockRequestQueue(queueError, hashingRequest)
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(queueError,hashSubmission) 
	mockWaiter := mocks.MockWaiter{0}
	hasher := Hasher{waiter: &mockWaiter, requestQueue: &mockRequestQueue, submissionQueue: &mockSubmissionQueue}
	
	err := hasher.processOrSleep()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHasherProcessOrSleepWaiterSleepsSuccess(t *testing.T) {
	requestQueueError := errors.New("test error")
	hashingRequest := models.HashingRequest{}
	mockRequestQueue := mocks.NewMockRequestQueue(requestQueueError, hashingRequest)
	mockWaiter := mocks.MockWaiter{0}
	hasher := Hasher{waiter: &mockWaiter, requestQueue: &mockRequestQueue}
	
	err := hasher.processOrSleep()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHasherProcessOrSleepWaiterSleepsWaiterCalled(t *testing.T) {
	requestQueueError := errors.New("test error")
	hashingRequest := models.HashingRequest{}
	mockRequestQueue := mocks.NewMockRequestQueue(requestQueueError, hashingRequest)
	mockWaiter := mocks.MockWaiter{0}
	hasher := Hasher{waiter: &mockWaiter, requestQueue: &mockRequestQueue}
	
	hasher.processOrSleep()
	expected := uint64(1)
	actual := mockWaiter.WaitCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestHasherProcessOrSleepError(t *testing.T) {
	requestQueueError := error(nil)
	hashingRequest := models.HashingRequest {
		HashName: "fakeHashFunction",
		Passwords: []string {
			"password123",
		},
	}
	mockRequestQueue := mocks.NewMockRequestQueue(requestQueueError, hashingRequest)
	mockWaiter := mocks.MockWaiter{0}
	hasher := Hasher{waiter: &mockWaiter, requestQueue: &mockRequestQueue}
	
	err := hasher.processOrSleep()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHasherHandleHashingRequestSuccess(t *testing.T) {
	submissionQueueError := error(nil)
	hashSubmission := models.HashSubmission{}
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(submissionQueueError, hashSubmission)
	hasher := Hasher{submissionQueue: &mockSubmissionQueue}		

	hashingRequest := models.HashingRequest {
		HashName: "sha256",
		Passwords: []string {
			"password123",
		},
	}

	err := hasher.handleHashingRequest(hashingRequest)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHasherHandleHashingRequestHashSubmissionError(t *testing.T) {
	submissionQueueError := error(nil)
	hashSubmission := models.HashSubmission{}
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(submissionQueueError, hashSubmission)
	hasher := Hasher{submissionQueue: &mockSubmissionQueue}		

	hashingRequest := models.HashingRequest {
		HashName: "fakeHashFunction",
		Passwords: []string {
			"password123",
		},
	}

	err := hasher.handleHashingRequest(hashingRequest)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestHasherHandleHashingRequestHashSubmissionErrorSubmissionQueuePutNotCalled(t *testing.T) {
	submissionQueueError := error(nil)
	hashSubmission := models.HashSubmission{}
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(submissionQueueError, hashSubmission)
	hasher := Hasher{submissionQueue: &mockSubmissionQueue}		

	hashingRequest := models.HashingRequest {
		HashName: "fakeHashFunction",
		Passwords: []string {
			"password123",
		},
	}

	hasher.handleHashingRequest(hashingRequest)

	expected := uint64(0)
	actual := mockSubmissionQueue.PutCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestHasherHandleHashingRequestSubmissionQueueError(t *testing.T) {
	submissionQueueError := errors.New("test error")
	hashSubmission := models.HashSubmission{}
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(submissionQueueError, hashSubmission)
	hasher := Hasher{submissionQueue: &mockSubmissionQueue}		

	hashingRequest := models.HashingRequest {
		HashName: "sha256",
		Passwords: []string {
			"password123",
		},
	}

	err := hasher.handleHashingRequest(hashingRequest)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestHasherHandleHashingRequestSubmissionQueueErrorPutCalled(t *testing.T) {
	submissionQueueError := errors.New("test error")
	hashSubmission := models.HashSubmission{}
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(submissionQueueError, hashSubmission)
	hasher := Hasher{submissionQueue: &mockSubmissionQueue}		

	hashingRequest := models.HashingRequest {
		HashName: "sha256",
		Passwords: []string {
			"password123",
		},
	}

	hasher.handleHashingRequest(hashingRequest)
	
	expected := uint64(1)
	actual := mockSubmissionQueue.PutCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestHasherGetHashSubmissionCorrectResults(t *testing.T) {
	expected := models.HashSubmission {
		HashType: "sha256",
		Results: []string {
			"password123:ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f",
			"hunter2:f52fbd32b2b3b86ff88ef6c490628285f482af15ddcb29541f94bcf526a3f6c7",
		},
	}
	passwords := []string{
		"password123",
		"hunter2",
	}
	hashingRequest := models.HashingRequest {
		HashName: "sha256",
		Passwords: passwords,
	}
	hasher := Hasher{}

	actual, _ := hasher.getHashSubmission(hashingRequest)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", expected, actual)
	}
}

func TestHasherGetHashSubmissionSuccess(t *testing.T) {
	passwords := []string{"password123"}
	hashingRequest := models.HashingRequest {
		HashName: "sha256",
		Passwords: passwords,
	}
	hasher := Hasher{}

	_, err := hasher.getHashSubmission(hashingRequest)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHasherGetHashSubmissionError(t *testing.T) {
	passwords := []string{"password123"}
	hashingRequest := models.HashingRequest {
		HashName: "fakeHashFunction",
		Passwords: passwords,
	}
	hasher := Hasher{}

	_, err := hasher.getHashSubmission(hashingRequest)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestHasherGetHashFunctionHashSupported(t *testing.T) {
	hasher := Hasher{}
	_, err := hasher.getHashFunction("sha256")
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHasherGetHashFunctionHashNotSupported(t *testing.T) {
	hasher := Hasher{}
	_, err := hasher.getHashFunction("fakeHashFunction")
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestHasherGetPasswordHashesCorrectResults(t *testing.T) {
	hashResults := []string {
		"f52fbd32b2b3b86ff88ef6c490628285f482af15ddcb29541f94bcf526a3f6c7",
		"ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f",
	}
	passwords := []string {
		"hunter2",
		"password123",
	}
	hashFunction := sha256.Sum256

	var expected []string
	for i, _ := range passwords {
		expectedResult := passwords[i] + ":" + hashResults[i]
		expected = append(expected, expectedResult)
	}

	actual := getPasswordHashes(hashFunction, passwords)
	for i, _ := range expected {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
		}
	}
}

func TestHasherGetPasswordHashCorrectResult(t *testing.T) {
	hashResult := "f52fbd32b2b3b86ff88ef6c490628285f482af15ddcb29541f94bcf526a3f6c7"
	password := "hunter2"
	hashFunction := sha256.Sum256

	expected := password + ":" + hashResult
	actual := getPasswordHash(hashFunction, password)
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}
