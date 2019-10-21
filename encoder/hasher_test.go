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

var nilError = error(nil)

type testObject struct {
	logger *mocks.MockLogger
	requestQueue *mocks.MockRequestQueue
	submissionQueue *mocks.MockSubmissionQueue
	waiter *mocks.MockWaiter
	hasher *Hasher
}

var hashingRequest = models.HashingRequest {
	Hash: sha256.New(),
	HashName: "sha256",
	Passwords: []string {
		"password123",
	},
}

func setupHasherForSuccess() testObject {
	hashSubmission := models.HashSubmission{}
	mockLogger := mocks.NewMockLogger(nilError)
	mockRequestQueue := mocks.NewMockRequestQueue(nilError, hashingRequest, 0)
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(nilError, hashSubmission, 0)
	mockWaiter := mocks.MockWaiter{0}
	hasher := Hasher {
		config:          nil,
		logger:          &mockLogger,
		requestQueue:    &mockRequestQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject {
		logger:          &mockLogger,
		requestQueue:    &mockRequestQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
		hasher:          &hasher,
	}
}

func setupHasherForSubmissionQueueError() testObject {
	submissionQueueError := errors.New("test error")
	hashSubmission := models.HashSubmission{}

	mockLogger := mocks.NewMockLogger(nilError)
	mockRequestQueue := mocks.NewMockRequestQueue(nilError, hashingRequest, 0)
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(submissionQueueError, hashSubmission, 0)
	mockWaiter := mocks.MockWaiter{0}
	hasher := Hasher {
		config:          nil,
		logger:          &mockLogger,
		requestQueue:    &mockRequestQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject {
		logger:          &mockLogger,
		requestQueue:    &mockRequestQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
		hasher:          &hasher,
	}
}

func setupHasherForRequestQueueError() testObject {
	requestQueueError := errors.New("test error")
	hashSubmission := models.HashSubmission{}

	mockLogger := mocks.NewMockLogger(nilError)
	mockRequestQueue := mocks.NewMockRequestQueue(requestQueueError, hashingRequest, 0)
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(nilError, hashSubmission, 0)
	mockWaiter := mocks.MockWaiter{0}
	hasher := Hasher {
		config:          nil,
		logger:          &mockLogger,
		requestQueue:    &mockRequestQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject {
		logger:          &mockLogger,
		requestQueue:    &mockRequestQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
		hasher:          &hasher,
	}
}

func TestHasherStartError(t *testing.T) {
	testObject := setupHasherForSubmissionQueueError()

	err := testObject.hasher.Start()

	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHasherProcessOrSleepProcessSuccess(t *testing.T) {
	testObject := setupHasherForSuccess()
	err := testObject.hasher.processOrSleep()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHasherProcessOrSleepWaiterSleeps(t *testing.T) {
	testObject := setupHasherForRequestQueueError()

	err := testObject.hasher.processOrSleep()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHasherProcessOrSleepWaiterSleepsWaiterCalled(t *testing.T) {
	testObject := setupHasherForRequestQueueError()

	testObject.hasher.processOrSleep()

	expected := uint64(1)
	actual := testObject.waiter.WaitCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestHasherProcessOrSleepError(t *testing.T) {
	testObject := setupHasherForSubmissionQueueError()

	err := testObject.hasher.processOrSleep()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHasherHandleHashingRequestSuccess(t *testing.T) {
	testObject := setupHasherForSuccess()

	err := testObject.hasher.handleHashingRequest(hashingRequest)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHasherHandleHashingRequestHashSubmissionError(t *testing.T) {
	testObject := setupHasherForSubmissionQueueError()

	err := testObject.hasher.handleHashingRequest(hashingRequest)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestHasherHandleHashingRequestSubmissionQueueError(t *testing.T) {
	testObject := setupHasherForSubmissionQueueError()

	err := testObject.hasher.handleHashingRequest(hashingRequest)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestHasherHandleHashingRequestSubmissionQueueErrorPutCalled(t *testing.T) {
	testObject := setupHasherForSubmissionQueueError()

	testObject.hasher.handleHashingRequest(hashingRequest)
	
	expected := uint64(1)
	actual := testObject.submissionQueue.PutCalls
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
		Hash: sha256.New(),
		HashName: "sha256",
		Passwords: passwords,
	}
	testObject := setupHasherForSuccess()

	actual := testObject.hasher.getHashSubmission(hashingRequest)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", expected, actual)
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
	hashFunction := sha256.New()

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
	hashFunction := sha256.New()

	expected := password + ":" + hashResult
	actual := getPasswordHash(hashFunction, password)
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}
