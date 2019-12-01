package submitter

import (
	"errors"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"reflect"
	"testing"
)

var testError = errors.New("test error")
var nilError = error(nil)
var threads = uint16(43)
var verboseConfig = models.Config{Verbose: true, Threads: threads}
var nonVerboseConfig = models.Config{Verbose: false, Threads: threads}

type testObject struct {
	hashSubmitter       *HashSubmitter
	mockApiClient       *mocks.MockApiClient
	mockLogger          *mocks.MockLogger
	mockStopQueue       *mocks.MockClientStopQueue
	mockSubmissionQueue *mocks.MockSubmissionQueue
	mockWaiter          *mocks.MockWaiter
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
	stopReason := models.ClientStopReason{}
	return mocks.NewMockClientStopQueue(stopReason, testError, testError)
}

func setupHashSubmitterForNoError() testObject {
	mockApiClient := mocks.NewMockApiClient(200, 200, 200, "fakeHash", []string{})
	mockLogger := mocks.NewMockLogger(nilError)
	mockStopQueue := setupStopQueueForEmptyReturn()
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(nilError, models.HashSubmission{}, 1)
	mockWaiter := mocks.NewMockWaiter()
	hashSubmitter := HashSubmitter{
		config:          &verboseConfig,
		client:          &mockApiClient,
		logger:          &mockLogger,
		stopQueue:       &mockStopQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject{
		hashSubmitter:       &hashSubmitter,
		mockApiClient:       &mockApiClient,
		mockLogger:          &mockLogger,
		mockStopQueue:       &mockStopQueue,
		mockSubmissionQueue: &mockSubmissionQueue,
		mockWaiter:          &mockWaiter,
	}
}

func setupHashSubmitterForStopReasonReturn() testObject {
	mockApiClient := mocks.NewMockApiClient(200, 200, 200, "fakeHash", []string{})
	mockLogger := mocks.NewMockLogger(nilError)
	mockStopQueue := setupStopQueueForStopReasonReturn()
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(nilError, models.HashSubmission{}, 1)
	mockWaiter := mocks.NewMockWaiter()
	hashSubmitter := HashSubmitter{
		config:          &verboseConfig,
		client:          &mockApiClient,
		logger:          &mockLogger,
		stopQueue:       &mockStopQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject{
		hashSubmitter:       &hashSubmitter,
		mockApiClient:       &mockApiClient,
		mockLogger:          &mockLogger,
		mockStopQueue:       &mockStopQueue,
		mockSubmissionQueue: &mockSubmissionQueue,
		mockWaiter:          &mockWaiter,
	}
}

func setupHashSubmitterForNoErrorEmptySubmissionQueue() testObject {
	mockApiClient := mocks.NewMockApiClient(200, 200, 200, "fakeHash", []string{})
	mockLogger := mocks.NewMockLogger(nilError)
	mockStopQueue := setupStopQueueForEmptyReturn()
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(nilError, models.HashSubmission{}, 0)
	mockWaiter := mocks.NewMockWaiter()
	hashSubmitter := HashSubmitter{
		config:          &verboseConfig,
		client:          &mockApiClient,
		logger:          &mockLogger,
		stopQueue:       &mockStopQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject{
		hashSubmitter:       &hashSubmitter,
		mockApiClient:       &mockApiClient,
		mockLogger:          &mockLogger,
		mockStopQueue:       &mockStopQueue,
		mockSubmissionQueue: &mockSubmissionQueue,
		mockWaiter:          &mockWaiter,
	}
}

func setupHashSubmitterForClientError() testObject {
	mockApiClient := mocks.NewMockApiClient(500, 500, 500, "fakeHash", []string{})
	mockLogger := mocks.NewMockLogger(nilError)
	mockStopQueue := setupStopQueueForEmptyReturn()
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(nilError, models.HashSubmission{}, 1)
	mockWaiter := mocks.NewMockWaiter()
	hashSubmitter := HashSubmitter{
		config:          &verboseConfig,
		client:          &mockApiClient,
		logger:          &mockLogger,
		stopQueue:       &mockStopQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject{
		hashSubmitter:       &hashSubmitter,
		mockApiClient:       &mockApiClient,
		mockLogger:          &mockLogger,
		mockStopQueue:       &mockStopQueue,
		mockSubmissionQueue: &mockSubmissionQueue,
		mockWaiter:          &mockWaiter,
	}
}

func setupHashSubmitterForSubmissionQueueError() testObject {
	mockApiClient := mocks.NewMockApiClient(200, 0, 0, "fakeHash", []string{})
	mockLogger := mocks.NewMockLogger(nilError)
	mockStopQueue := setupStopQueueForEmptyReturn()
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(testError, models.HashSubmission{}, 1)
	mockWaiter := mocks.NewMockWaiter()
	hashSubmitter := HashSubmitter{
		config:          &verboseConfig,
		client:          &mockApiClient,
		logger:          &mockLogger,
		stopQueue:       &mockStopQueue,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject{
		hashSubmitter:       &hashSubmitter,
		mockApiClient:       &mockApiClient,
		mockLogger:          &mockLogger,
		mockStopQueue:       &mockStopQueue,
		mockSubmissionQueue: &mockSubmissionQueue,
		mockWaiter:          &mockWaiter,
	}
}

func assertClientSubmitHashesCalled(t *testing.T, testObject testObject) {
	expected := uint64(1)
	actual := testObject.mockApiClient.SubmitHashesCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func assertSubmissionQueueGetCalled(t *testing.T, testObject testObject) {
	expected := uint64(1)
	actual := testObject.mockSubmissionQueue.GetCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func assertSubmissionQueueSizeCalled(t *testing.T, testObject testObject) {
	expected := uint64(1)
	actual := testObject.mockSubmissionQueue.SizeCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func assertWaiterWaitCalled(t *testing.T, testObject testObject) {
	expected := uint64(1)
	actual := testObject.mockWaiter.WaitCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func assertClientSubmitHashesNotCalled(t *testing.T, testObject testObject) {
	expected := uint64(0)
	actual := testObject.mockApiClient.SubmitHashesCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func assertWaiterWaitNotCalled(t *testing.T, testObject testObject) {
	expected := uint64(0)
	actual := testObject.mockWaiter.WaitCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func assertLoggerCalledOnce(t *testing.T, testObject testObject) {
	expected := uint64(1)
	actual := testObject.mockLogger.LogMessageCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func assertLoggerCalledNTimes(t *testing.T, testObject testObject, n uint64) {
	expected := n
	actual := testObject.mockLogger.LogMessageCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func assertLoggerNotCalled(t *testing.T, testObject testObject) {
	expected := uint64(0)
	actual := testObject.mockLogger.LogMessageCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func assertStopQueuePutNotCalled(t *testing.T, testObject testObject) {
	expected := uint64(0)
	actual := testObject.mockStopQueue.PutCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func assertStopQueuePutCalledNTimes(t *testing.T, testObject testObject, n uint64) {
	expected := n
	actual := testObject.mockStopQueue.PutCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func TestNewHashSubmitter(t *testing.T) {
	testObject := setupHashSubmitterForNoError()
	expected := testObject.hashSubmitter

	configProvider := mocks.NewMockConfigProvider(expected.config)

	actual := NewHashSubmitter(&configProvider, expected.client, expected.logger, expected.submissionQueue, expected.stopQueue, expected.waiter)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", expected, actual)
	}
}

func TestHashSubmitter_Start_ProcessOrSleepError(t *testing.T) {
	testObject := setupHashSubmitterForClientError()
	err := testObject.hashSubmitter.Start()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHashSubmitter_Start_ProcessOrSleepError_StopQueuePutCalled(t *testing.T) {
	testObject := setupHashSubmitterForClientError()
	testObject.hashSubmitter.Start()
	assertStopQueuePutCalledNTimes(t, testObject, uint64(threads - 1))
}

func TestHashSubmitter_Start_StopQueueError(t *testing.T) {
	testObject := setupHashSubmitterForStopReasonReturn()
	err := testObject.hashSubmitter.Start()
	if err == nil {
		t.Error("Expected error but nil returned")
	}}


func TestHashSubmitter_Start_StopQueueError_StopQueuePutNotCalled(t *testing.T) {
	testObject := setupHashSubmitterForStopReasonReturn()
	testObject.hashSubmitter.Start()
	assertStopQueuePutNotCalled(t, testObject)
}

func TestHashSubmitter_ProcessOrSleep_Success(t *testing.T) {
	testObject := setupHashSubmitterForNoError()
	err := testObject.hashSubmitter.processOrSleep()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHashSubmitter_ProcessOrSleep_Success_CorrectSizeCalls(t *testing.T) {
	testObject := setupHashSubmitterForNoError()
	testObject.hashSubmitter.processOrSleep()
	assertSubmissionQueueSizeCalled(t, testObject)
}

func TestHashSubmitter_ProcessOrSleep_Success_CorrectWaiterCalls(t *testing.T) {
	testObject := setupHashSubmitterForNoError()
	testObject.hashSubmitter.processOrSleep()
	assertWaiterWaitNotCalled(t, testObject)
}

func TestHashSubmitter_ProcessOrSleep_EmptySubmissionQueue_Success(t *testing.T) {
	testObject := setupHashSubmitterForNoErrorEmptySubmissionQueue()
	err := testObject.hashSubmitter.processOrSleep()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHashSubmitter_ProcessOrSleep_EmptySubmissionQueue_Success_CorrectSizeCalls(t *testing.T) {
	testObject := setupHashSubmitterForNoErrorEmptySubmissionQueue()
	testObject.hashSubmitter.processOrSleep()
	assertSubmissionQueueSizeCalled(t, testObject)
}

func TestHashSubmitter_ProcessOrSleep_EmptySubmissionQueue_Success_CorrectWaiterCalls(t *testing.T) {
	testObject := setupHashSubmitterForNoErrorEmptySubmissionQueue()
	testObject.hashSubmitter.processOrSleep()
	assertWaiterWaitCalled(t, testObject)
}

func TestHashSubmitter_ProcessOrSleep_ProcessSubmissionError(t *testing.T) {
	testObject := setupHashSubmitterForClientError()
	err := testObject.hashSubmitter.processOrSleep()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHashSubmitter_ProcessOrSleep_ProcessSubmissionError_CorrectSizeCalls(t *testing.T) {
	testObject := setupHashSubmitterForClientError()
	testObject.hashSubmitter.processOrSleep()
	assertSubmissionQueueSizeCalled(t, testObject)
}

func TestHashSubmitter_ProcessOrSleep_ProcessSubmissionError_CorrectWaiterCalls(t *testing.T) {
	testObject := setupHashSubmitterForClientError()
	testObject.hashSubmitter.processOrSleep()
	assertWaiterWaitNotCalled(t, testObject)
}

func TestHashSubmitter_ProcessSubmission_Success(t *testing.T) {
	testObject := setupHashSubmitterForNoError()

	err := testObject.hashSubmitter.processSubmission()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHashSubmitter_ProcessSubmission_Success_CorrectSubmissionQueueCalls(t *testing.T) {
	testObject := setupHashSubmitterForNoError()
	testObject.hashSubmitter.processSubmission()
	assertSubmissionQueueGetCalled(t, testObject)
}

func TestHashSubmitter_ProcessSubmission_Success_CorrectClientCalls(t *testing.T) {
	testObject := setupHashSubmitterForNoError()
	testObject.hashSubmitter.processSubmission()
	assertClientSubmitHashesCalled(t, testObject)
}

func TestHashSubmitter_ProcessSubmission_Success_LoggerCalledTwice(t *testing.T) {
	testObject := setupHashSubmitterForNoError()
	testObject.hashSubmitter.processSubmission()
	assertLoggerCalledNTimes(t, testObject, 2)
}

func TestHashSubmitter_ProcessSubmission_Error_BadStatusCodeReturnedFromClient(t *testing.T) {
	testObject := setupHashSubmitterForClientError()

	err := testObject.hashSubmitter.processSubmission()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHashSubmitter_ProcessSubmission_BadStatusCodeReturnedFromClient_CorrectSubmissionQueueCalls(t *testing.T) {
	testObject := setupHashSubmitterForClientError()
	testObject.hashSubmitter.processSubmission()
	assertSubmissionQueueGetCalled(t, testObject)
}

func TestHashSubmitter_ProcessSubmission_BadStatusCodeReturnedFromClient_CorrectClientCalls(t *testing.T) {
	testObject := setupHashSubmitterForClientError()
	testObject.hashSubmitter.processSubmission()
	assertClientSubmitHashesCalled(t, testObject)
}

func TestHashSubmitter_ProcessSubmission_BadStatusCodeReturnedFromClient_LoggerCalledOnce(t *testing.T) {
	testObject := setupHashSubmitterForClientError()
	testObject.hashSubmitter.processSubmission()
	assertLoggerCalledOnce(t, testObject)
}

func TestHashSubmitter_ProcessSubmission_SubmissionQueue_Error(t *testing.T) {
	testObject := setupHashSubmitterForSubmissionQueueError()

	err := testObject.hashSubmitter.processSubmission()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHashSubmitter_ProcessSubmission_SubmissionQueueError_CorrectSubmissionQueueCalls(t *testing.T) {
	testObject := setupHashSubmitterForSubmissionQueueError()
	testObject.hashSubmitter.processSubmission()
	assertSubmissionQueueGetCalled(t, testObject)
}

func TestHashSubmitter_ProcessSubmission_SubmissionQueueError_CorrectClientCalls(t *testing.T) {
	testObject := setupHashSubmitterForSubmissionQueueError()
	testObject.hashSubmitter.processSubmission()
	assertClientSubmitHashesNotCalled(t, testObject)
}

func TestHashSubmitter_ProcessSubmission_SubmissionQueueError_LoggerNotCalled(t *testing.T) {
	testObject := setupHashSubmitterForSubmissionQueueError()
	testObject.hashSubmitter.processSubmission()
	assertLoggerNotCalled(t, testObject)
}
