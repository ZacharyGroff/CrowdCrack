package submitter

import (
	"errors"
	"testing"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

var testError = errors.New("test error")
var nilError = error(nil)

type testObject struct {
	hashSubmitter *HashSubmitter
	mockApiClient *mocks.MockApiClient
	mockLogger *mocks.MockLogger
	mockSubmissionQueue *mocks.MockSubmissionQueue
	mockWaiter *mocks.MockWaiter
}

func setupHashSubmitterForNoError() testObject {
	mockApiClient := mocks.NewMockApiClient(200, 200, 200, "fakeHash", []string{})
	mockLogger := mocks.NewMockLogger(nilError)
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(nilError, models.HashSubmission{}, 1)
	mockWaiter := mocks.NewMockWaiter()
	hashSubmitter := HashSubmitter {
		config:          nil,
		client:          &mockApiClient,
		logger:          &mockLogger,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject {
		hashSubmitter:       &hashSubmitter,
		mockApiClient:       &mockApiClient,
		mockLogger:          &mockLogger,
		mockSubmissionQueue: &mockSubmissionQueue,
		mockWaiter:          &mockWaiter,
	}
}

func setupHashSubmitterForNoErrorEmptySubmissionQueue() testObject {
	mockApiClient := mocks.NewMockApiClient(200, 200, 200, "fakeHash", []string{})
	mockLogger := mocks.NewMockLogger(nilError)
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(nilError, models.HashSubmission{}, 0)
	mockWaiter := mocks.NewMockWaiter()
	hashSubmitter := HashSubmitter {
		config:          nil,
		client:          &mockApiClient,
		logger:          &mockLogger,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject {
		hashSubmitter:       &hashSubmitter,
		mockApiClient:       &mockApiClient,
		mockLogger:          &mockLogger,
		mockSubmissionQueue: &mockSubmissionQueue,
		mockWaiter:          &mockWaiter,
	}
}

func setupHashSubmitterForClientError() testObject {
	mockApiClient := mocks.NewMockApiClient(500, 500, 500, "fakeHash", []string{})
	mockLogger := mocks.NewMockLogger(nilError)
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(nilError, models.HashSubmission{}, 1)
	mockWaiter := mocks.NewMockWaiter()
	hashSubmitter := HashSubmitter {
		config:          nil,
		client:          &mockApiClient,
		logger:          &mockLogger,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject {
		hashSubmitter:       &hashSubmitter,
		mockApiClient:       &mockApiClient,
		mockLogger:          &mockLogger,
		mockSubmissionQueue: &mockSubmissionQueue,
		mockWaiter:          &mockWaiter,
	}
}

func setupHashSubmitterForSubmissionQueueError() testObject {
	mockApiClient := mocks.NewMockApiClient(200, 0, 0, "fakeHash", []string{})
	mockLogger := mocks.NewMockLogger(nilError)
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(testError, models.HashSubmission{}, 1)
	mockWaiter := mocks.NewMockWaiter()
	hashSubmitter := HashSubmitter {
		config:          nil,
		client:          &mockApiClient,
		logger:          &mockLogger,
		submissionQueue: &mockSubmissionQueue,
		waiter:          &mockWaiter,
	}

	return testObject {
		hashSubmitter:       &hashSubmitter,
		mockApiClient:       &mockApiClient,
		mockLogger:          &mockLogger,
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

func TestStartError(t *testing.T) {
	testObject := setupHashSubmitterForClientError()
	err := testObject.hashSubmitter.Start()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestProcessOrSleepSuccess(t *testing.T) {
	testObject := setupHashSubmitterForNoError()
	err := testObject.hashSubmitter.processOrSleep()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestProcessOrSleepSuccessCorrectSizeCalls(t *testing.T) {
	testObject := setupHashSubmitterForNoError()
	testObject.hashSubmitter.processOrSleep()
	assertSubmissionQueueSizeCalled(t, testObject)
}

func TestProcessOrSleepSuccessCorrectWaiterCalls(t *testing.T) {
	testObject := setupHashSubmitterForNoError()
	testObject.hashSubmitter.processOrSleep()
	assertWaiterWaitNotCalled(t, testObject)
}

func TestProcessOrSleepEmptySubmissionQueueSuccess(t *testing.T) {
	testObject := setupHashSubmitterForNoErrorEmptySubmissionQueue()
	err := testObject.hashSubmitter.processOrSleep()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestProcessOrSleepEmptySubmissionQueueSuccessCorrectSizeCalls(t *testing.T) {
	testObject := setupHashSubmitterForNoErrorEmptySubmissionQueue()
	testObject.hashSubmitter.processOrSleep()
	assertSubmissionQueueSizeCalled(t, testObject)
}

func TestProcessOrSleepEmptySubmissionQueueSuccessCorrectWaiterCalls(t *testing.T) {
	testObject := setupHashSubmitterForNoErrorEmptySubmissionQueue()
	testObject.hashSubmitter.processOrSleep()
	assertWaiterWaitCalled(t, testObject)
}

func TestProcessOrSleepProcessSubmissionError(t *testing.T) {
	testObject := setupHashSubmitterForClientError()
	err := testObject.hashSubmitter.processOrSleep()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestProcessOrSleepProcessSubmissionErrorCorrectSizeCalls(t *testing.T) {
	testObject := setupHashSubmitterForClientError()
	testObject.hashSubmitter.processOrSleep()
	assertSubmissionQueueSizeCalled(t, testObject)
}

func TestProcessOrSleepProcessSubmissionErrorCorrectWaiterCalls(t *testing.T) {
	testObject := setupHashSubmitterForClientError()
	testObject.hashSubmitter.processOrSleep()
	assertWaiterWaitNotCalled(t, testObject)
}

func TestProcessSubmissionSuccess(t *testing.T) {
	testObject := setupHashSubmitterForNoError()

	err := testObject.hashSubmitter.processSubmission()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestProcessSubmissionSuccessCorrectSubmissionQueueCalls(t *testing.T) {
	testObject := setupHashSubmitterForNoError()
	testObject.hashSubmitter.processSubmission()
	assertSubmissionQueueGetCalled(t, testObject)
}

func TestProcessSubmissionSuccessCorrectClientCalls(t *testing.T) {
	testObject := setupHashSubmitterForNoError()
	testObject.hashSubmitter.processSubmission()
	assertClientSubmitHashesCalled(t, testObject)
}

func TestProcessSubmissionClientBadStatusCodeReturned(t *testing.T) {
	testObject := setupHashSubmitterForClientError()

	err := testObject.hashSubmitter.processSubmission()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestProcessSubmissionClientBadStatusCodeReturnedCorrectSubmissionQueueCalls(t *testing.T) {
	testObject := setupHashSubmitterForClientError()
	testObject.hashSubmitter.processSubmission()
	assertSubmissionQueueGetCalled(t, testObject)
}

func TestProcessSubmissionClientBadStatusCodeReturnedCorrectClientCalls(t *testing.T) {
	testObject := setupHashSubmitterForClientError()
	testObject.hashSubmitter.processSubmission()
	assertClientSubmitHashesCalled(t, testObject)
}

func TestProcessSubmissionSubmissionQueueErrorReturned(t *testing.T) {
	testObject := setupHashSubmitterForSubmissionQueueError()

	err := testObject.hashSubmitter.processSubmission()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestProcessSubmissionSubmissionQueueErrorCorrectSubmissionQueueCalls(t *testing.T) {
	testObject := setupHashSubmitterForSubmissionQueueError()
	testObject.hashSubmitter.processSubmission()
	assertSubmissionQueueGetCalled(t, testObject)
}

func TestProcessSubmissionSubmissionQueueErrorCorrectClientCalls(t *testing.T) {
	testObject := setupHashSubmitterForSubmissionQueueError()
	testObject.hashSubmitter.processSubmission()
	assertClientSubmitHashesNotCalled(t, testObject)
}
