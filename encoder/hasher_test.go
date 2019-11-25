package encoder

import (
	"crypto/sha256"
	"errors"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"reflect"
	"strings"
	"sync"
	"testing"
)

var threads = uint16(43)
var nilError = error(nil)
var testError = errors.New("testError")
var verboseConfig = models.Config{Verbose: true, Threads: threads}
var nonVerboseConfig = models.Config{Verbose: false, Threads: threads}

type testObject struct {
	logger          *mocks.MockLogger
	requestQueue    *mocks.MockRequestQueue
	stopQueue       *mocks.MockClientStopQueue
	submissionQueue *mocks.MockSubmissionQueue
	waiter          *mocks.MockWaiter
	hasher          *Hasher
}

var hashingRequest = models.HashingRequest{
	Hash:     sha256.New(),
	HashName: "sha256",
	Passwords: []string{
		"password123",
	},
}

func setupStopQueueForStopReasonReturn() mocks.MockClientStopQueue {
	stopReason := models.ClientStopReason{
		Requester: "",
		Encoder:   "",
		Submitter: "",
	}
	return mocks.NewMockClientStopQueue(stopReason, nilError, nilError)
}

func setupStopQueueForEmptyReturn() mocks.MockClientStopQueue {
	stopReason := models.ClientStopReason{
		Requester: "",
		Encoder:   "",
		Submitter: "",
	}
	return mocks.NewMockClientStopQueue(stopReason, nilError, nilError)
}

func setupHasherForSuccess() testObject {
	hashSubmission := models.HashSubmission{}
	mockLogger := mocks.NewMockLogger(nilError)
	mockRequestQueue := mocks.NewMockRequestQueue(nilError, hashingRequest, 0)
	mockStopQueue := setupStopQueueForEmptyReturn()
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(nilError, hashSubmission, 0)
	mockWaiter := mocks.MockWaiter{0}
	mux := new(sync.Mutex)
	hasher := Hasher{
		config:          &verboseConfig,
		logger:          &mockLogger,
		mux:             mux,
		requestQueue:    &mockRequestQueue,
		stopQueue:       &mockStopQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject{
		logger:          &mockLogger,
		requestQueue:    &mockRequestQueue,
		stopQueue:       &mockStopQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
		hasher:          &hasher,
	}
}

func setupHasherForStopReason() testObject {
	hashSubmission := models.HashSubmission{}
	mockLogger := mocks.NewMockLogger(nilError)
	mockRequestQueue := mocks.NewMockRequestQueue(nilError, hashingRequest, 0)
	mockStopQueue := setupStopQueueForStopReasonReturn()
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(nilError, hashSubmission, 0)
	mockWaiter := mocks.MockWaiter{0}
	mux := new(sync.Mutex)
	hasher := Hasher{
		config:          &verboseConfig,
		logger:          &mockLogger,
		mux:             mux,
		requestQueue:    &mockRequestQueue,
		stopQueue:       &mockStopQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject{
		logger:          &mockLogger,
		requestQueue:    &mockRequestQueue,
		stopQueue:       &mockStopQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
		hasher:          &hasher,
	}
}

func setupHasherForSuccessNonVerbose() testObject {
	hashSubmission := models.HashSubmission{}
	mockLogger := mocks.NewMockLogger(nilError)
	mockRequestQueue := mocks.NewMockRequestQueue(nilError, hashingRequest, 0)
	mockStopQueue := setupStopQueueForEmptyReturn()
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(nilError, hashSubmission, 0)
	mockWaiter := mocks.MockWaiter{0}
	mux := new(sync.Mutex)
	hasher := Hasher{
		config:          &nonVerboseConfig,
		logger:          &mockLogger,
		mux:             mux,
		requestQueue:    &mockRequestQueue,
		stopQueue:       &mockStopQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject{
		logger:          &mockLogger,
		requestQueue:    &mockRequestQueue,
		stopQueue:       &mockStopQueue,
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
	mockStopQueue := setupStopQueueForEmptyReturn()
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(submissionQueueError, hashSubmission, 0)
	mockWaiter := mocks.MockWaiter{0}
	mux := new(sync.Mutex)
	hasher := Hasher{
		config:          &verboseConfig,
		logger:          &mockLogger,
		mux:             mux,
		requestQueue:    &mockRequestQueue,
		stopQueue:       &mockStopQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject{
		logger:          &mockLogger,
		requestQueue:    &mockRequestQueue,
		stopQueue:       &mockStopQueue,
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
	mockStopQueue := setupStopQueueForEmptyReturn()
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(nilError, hashSubmission, 0)
	mockWaiter := mocks.MockWaiter{0}
	mux := new(sync.Mutex)
	hasher := Hasher{
		config:          &verboseConfig,
		logger:          &mockLogger,
		mux:             mux,
		requestQueue:    &mockRequestQueue,
		stopQueue:       &mockStopQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject{
		logger:          &mockLogger,
		requestQueue:    &mockRequestQueue,
		stopQueue:       &mockStopQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
		hasher:          &hasher,
	}
}

func assertLoggerCalled(t *testing.T, testObject testObject) {
	expected := uint64(1)

	actual := testObject.logger.LogMessageCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertLoggerNotCalled(t *testing.T, testObject testObject) {
	expected := uint64(0)

	actual := testObject.logger.LogMessageCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertStopQueueGetCalled(t *testing.T, testObject testObject) {
	expected := uint64(1)

	actual := testObject.stopQueue.GetCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertStopQueueGetNotCalled(t *testing.T, testObject testObject) {
	expected := uint64(0)

	actual := testObject.stopQueue.GetCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertStopQueuePutCalledNTimes(t *testing.T, testObject testObject, n uint64) {
	expected := n

	actual := testObject.stopQueue.PutCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestHasher_Start_ProcessOrSleep_Error(t *testing.T) {
	testObject := setupHasherForSubmissionQueueError()

	err := testObject.hasher.Start()

	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHasher_Start_ProcessOrSleep_Error_StopQueueNotCalled(t *testing.T) {
	testObject := setupHasherForSubmissionQueueError()
	testObject.hasher.Start()
	assertStopQueueGetNotCalled(t, testObject)
}

func TestHasher_Start_StopQueue_Error(t *testing.T) {
	testObject := setupHasherForStopReason()
	err := testObject.hasher.Start()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHasher_Start_StopQueue_Error_StopQueueCalled(t *testing.T) {
	testObject := setupHasherForStopReason()
	testObject.hasher.Start()
	assertStopQueueGetCalled(t, testObject)
}

func TestHasher_ProcessOrSleep_Process_Success(t *testing.T) {
	testObject := setupHasherForSuccess()
	err := testObject.hasher.processOrSleep()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHasher_ProcessOrSleep_Sleep_Success(t *testing.T) {
	testObject := setupHasherForRequestQueueError()

	err := testObject.hasher.processOrSleep()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHasher_ProcessOrSleep_Sleep_WaiterCalled(t *testing.T) {
	testObject := setupHasherForRequestQueueError()

	testObject.hasher.processOrSleep()

	expected := uint64(1)
	actual := testObject.waiter.WaitCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestHasher_ProcessOrSleep_Error(t *testing.T) {
	testObject := setupHasherForSubmissionQueueError()

	err := testObject.hasher.processOrSleep()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHasher_HandleHashingRequest_Success(t *testing.T) {
	testObject := setupHasherForSuccess()

	err := testObject.hasher.handleHashingRequest(hashingRequest)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHasher_HandleHashingRequest_Success_LoggerCalled(t *testing.T) {
	testObject := setupHasherForSuccess()
	testObject.hasher.handleHashingRequest(hashingRequest)
	assertLoggerCalled(t, testObject)
}

func TestHasher_HandleHashingRequest_Success_LoggerNotCalled(t *testing.T) {
	testObject := setupHasherForSuccessNonVerbose()
	testObject.hasher.handleHashingRequest(hashingRequest)
	assertLoggerNotCalled(t, testObject)
}

func TestHasher_HandleHashingRequest_HashSubmissionError(t *testing.T) {
	testObject := setupHasherForSubmissionQueueError()

	err := testObject.hasher.handleHashingRequest(hashingRequest)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestHasher_HandleHashingRequest_SubmissionQueueError(t *testing.T) {
	testObject := setupHasherForSubmissionQueueError()

	err := testObject.hasher.handleHashingRequest(hashingRequest)
	if err == nil {
		t.Error("Expected error but nil returned.")
	}
}

func TestHasher_HandleHashingRequest_SubmissionQueueError_PutCalled(t *testing.T) {
	testObject := setupHasherForSubmissionQueueError()

	testObject.hasher.handleHashingRequest(hashingRequest)

	expected := uint64(1)
	actual := testObject.submissionQueue.PutCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestHasher_GetHashSubmission_CorrectResults(t *testing.T) {
	expected := models.HashSubmission{
		HashType: "sha256",
		Results: []string{
			"password123:ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f",
			"hunter2:f52fbd32b2b3b86ff88ef6c490628285f482af15ddcb29541f94bcf526a3f6c7",
		},
	}
	passwords := []string{
		"password123",
		"hunter2",
	}
	hashingRequest := models.HashingRequest{
		Hash:      sha256.New(),
		HashName:  "sha256",
		Passwords: passwords,
	}
	testObject := setupHasherForSuccess()

	actual := testObject.hasher.getHashSubmission(hashingRequest)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", expected, actual)
	}
}

func TestHasher_UpdateStopQueue_StopQueuePutCalls(t *testing.T) {
	testObject := setupHasherForSuccess()
	testObject.hasher.updateStopQueue(testError)
	assertStopQueuePutCalledNTimes(t, testObject, uint64(threads) - 1)
}

func TestHasher_GetPasswordHashes_CorrectResults(t *testing.T) {
	hashResults := []string{
		"f52fbd32b2b3b86ff88ef6c490628285f482af15ddcb29541f94bcf526a3f6c7",
		"ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f",
	}
	passwords := []string{
		"hunter2",
		"password123",
	}
	hashFunction := sha256.New()

	var expected []string
	for i, _ := range passwords {
		expectedResult := passwords[i] + ":" + hashResults[i]
		expected = append(expected, expectedResult)
	}

	testObject := setupHasherForSuccess()

	actual := testObject.hasher.getPasswordHashes(hashFunction, passwords)
	for i, _ := range expected {
		if strings.Compare(expected[i], actual[i]) != 0 {
			t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
		}
	}
}

func TestHasher_GetPasswordHash_CorrectResult(t *testing.T) {
	hashResult := "f52fbd32b2b3b86ff88ef6c490628285f482af15ddcb29541f94bcf526a3f6c7"
	password := "hunter2"
	hashFunction := sha256.New()

	testObject := setupHasherForSuccess()

	expected := password + ":" + hashResult
	actual := testObject.hasher.getPasswordHash(hashFunction, password)
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}
