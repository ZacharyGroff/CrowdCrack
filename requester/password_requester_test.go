package requester

import (
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"testing"
)

type testObject struct {
	passwordRequester *PasswordRequester
	apiClient *mocks.MockApiClient
	requestQueue *mocks.MockRequestQueue
	waiter *mocks.MockWaiter
}

var expectedPasswords = []string {
	"hunter2",
	"password123",
}

var expectedHashName = "testHashName"

func setupApiClientForSuccess() mocks.MockApiClient {
	statusCode := 200
	return mocks.NewMockApiClient(statusCode, expectedHashName, expectedPasswords)
}

func setupApiClientForError() mocks.MockApiClient {
	statusCode := 500
	return mocks.NewMockApiClient(statusCode, expectedHashName, expectedPasswords)
}

func setupRequestQueueForSuccess() mocks.MockRequestQueue {
	err := error(nil)
	hashingRequest := models.HashingRequest{}
	return mocks.NewMockRequestQueue(err, hashingRequest, 0)
}

func setupRequestQueueFull() mocks.MockRequestQueue {
	err := error(nil)
	hashingRequest := models.HashingRequest{}
	return mocks.NewMockRequestQueue(err, hashingRequest, 2)
}

func setupPasswordRequestForSuccess() testObject {
	apiClient := setupApiClientForSuccess()
	requestQueue := setupRequestQueueForSuccess()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{client: &apiClient, requestQueue: &requestQueue, waiter: &waiter}

	return testObject{&passwordRequester, &apiClient, &requestQueue, &waiter}
}

func setupPasswordRequestForApiClientError() testObject {
	apiClient := setupApiClientForError()
	requestQueue := setupRequestQueueForSuccess()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{client: &apiClient, requestQueue: &requestQueue, waiter: &waiter}

	return testObject{&passwordRequester, &apiClient, &requestQueue, &waiter}
}

func setupPasswordRequestForFullRequestQueue() testObject {
	apiClient := setupApiClientForSuccess()
	requestQueue := setupRequestQueueFull()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{client: &apiClient, requestQueue: &requestQueue, waiter: &waiter}

	return testObject{&passwordRequester, &apiClient, &requestQueue, &waiter}
}

func assertRequestQueueSizeCalled(t *testing.T, testObject testObject) {
	expected := uint64(1)
	actual := testObject.requestQueue.SizeCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertRequestQueuePutCalled(t *testing.T, testObject testObject) {
	expected := uint64(1)
	actual := testObject.requestQueue.PutCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertRequestQueuePutNotCalled(t *testing.T, testObject testObject) {
	expected := uint64(0)
	actual := testObject.requestQueue.PutCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertWaiterCalled(t *testing.T, testObject testObject) {
	expected := uint64(1)
	actual := testObject.waiter.WaitCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertWaiterNotCalled(t *testing.T, testObject testObject) {
	expected := uint64(0)
	actual := testObject.waiter.WaitCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertApiClientGetHashNameCalled(t *testing.T, testObject testObject) {
	expected := uint64(1)
	actual := testObject.apiClient.GetHashNameCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertApiClientGetPasswordsCalled(t *testing.T, testObject testObject) {
	expected := uint64(1)
	actual := testObject.apiClient.GetPasswordsCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}
