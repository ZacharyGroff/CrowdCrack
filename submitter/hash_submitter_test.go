package submitter

import (
	"errors"
	"testing"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type testObject struct {
	hashSubmitter *HashSubmitter
	mockSubmissionQueue *mocks.MockSubmissionQueue
	mockApiClient *mocks.MockApiClient
	mockWaiter *mocks.MockWaiter
}

func setupHashSubmitterForNoError() testObject {
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(error(nil), models.HashSubmission{}, 1)
	mockApiClient := mocks.NewMockApiClient(200, "fakeHash", []string{})
	mockWaiter := mocks.NewMockWaiter()
	hashSubmitter := HashSubmitter{submissionQueue: &mockSubmissionQueue, client: &mockApiClient, waiter: &mockWaiter}

	return testObject{&hashSubmitter, &mockSubmissionQueue, &mockApiClient, &mockWaiter}
}

func setupHashSubmitterForNoErrorEmptySubmissionQueue() testObject {
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(error(nil), models.HashSubmission{}, 0)
	mockApiClient := mocks.NewMockApiClient(200, "fakeHash", []string{})
	mockWaiter := mocks.NewMockWaiter()
	hashSubmitter := HashSubmitter{submissionQueue: &mockSubmissionQueue, client: &mockApiClient, waiter: &mockWaiter}

	return testObject{&hashSubmitter, &mockSubmissionQueue, &mockApiClient, &mockWaiter}
}

func setupHashSubmitterForClientError() testObject {
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(error(nil), models.HashSubmission{}, 1)
	mockApiClient := mocks.NewMockApiClient(500, "fakeHash", []string{})
	mockWaiter := mocks.NewMockWaiter()
	hashSubmitter := HashSubmitter{submissionQueue: &mockSubmissionQueue, client: &mockApiClient, waiter: &mockWaiter}

	return testObject{&hashSubmitter, &mockSubmissionQueue, &mockApiClient, &mockWaiter}
}

func setupHashSubmitterForSubmissionQueueError() testObject {
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(errors.New("test error"), models.HashSubmission{}, 1)
	mockApiClient := mocks.NewMockApiClient(200, "fakeHash", []string{})
	mockWaiter := mocks.NewMockWaiter()
	hashSubmitter := HashSubmitter{submissionQueue: &mockSubmissionQueue, client: &mockApiClient, waiter: &mockWaiter}

	return testObject{&hashSubmitter, &mockSubmissionQueue, &mockApiClient, &mockWaiter}
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