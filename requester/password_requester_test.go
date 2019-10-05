package requester

import (
	"hash"
	"reflect"
	"strings"
	"testing"
	"crypto/sha256"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
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
var expectedHashName = "sha256"
var expectedHash = sha256.New()
var successCode = 200
var errorCode = 500

func setupApiClientForSuccess() mocks.MockApiClient {
	return mocks.NewMockApiClient(successCode, successCode, successCode, expectedHashName, expectedPasswords)
}

func setupApiClientForError() mocks.MockApiClient {
	return mocks.NewMockApiClient(errorCode, errorCode, errorCode, expectedHashName, expectedPasswords)
}

func setupApiClientForGetHashNameError() mocks.MockApiClient {
	return mocks.NewMockApiClient(errorCode, successCode, successCode, expectedHashName, expectedPasswords)
}

func setupApiClientForGetPasswordsError() mocks.MockApiClient {
	return mocks.NewMockApiClient(successCode, errorCode, successCode, expectedHashName, expectedPasswords)
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

func setupSupportedHashes() map[string]hash.Hash {
	return map[string]hash.Hash {
		"sha256": sha256.New(),
	}
}

func setupNoSupportedHashes() map[string]hash.Hash {
	return map[string]hash.Hash {}
}

func setupPasswordRequestForSuccess() testObject {
	apiClient := setupApiClientForSuccess()
	requestQueue := setupRequestQueueForSuccess()
	supportedHashes := setupSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{client: &apiClient, requestQueue: &requestQueue, supportedHashes: supportedHashes, waiter: &waiter}

	return testObject{&passwordRequester, &apiClient, &requestQueue, &waiter}
}

func setupPasswordRequestForApiClientError() testObject {
	apiClient := setupApiClientForError()
	requestQueue := setupRequestQueueForSuccess()
	supportedHashes := setupSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{client: &apiClient, requestQueue: &requestQueue, supportedHashes: supportedHashes, waiter: &waiter}

	return testObject{&passwordRequester, &apiClient, &requestQueue, &waiter}
}

func setupPasswordRequestForGetHashNameError() testObject {
	apiClient := setupApiClientForGetHashNameError()
	requestQueue := setupRequestQueueForSuccess()
	supportedHashes := setupSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{client: &apiClient, requestQueue: &requestQueue, supportedHashes: supportedHashes, waiter: &waiter}

	return testObject{&passwordRequester, &apiClient, &requestQueue, &waiter}
}

func setupPasswordRequestForGetPasswordsError() testObject {
	apiClient := setupApiClientForGetPasswordsError()
	requestQueue := setupRequestQueueForSuccess()
	supportedHashes := setupSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{client: &apiClient, requestQueue: &requestQueue, supportedHashes: supportedHashes, waiter: &waiter}

	return testObject{&passwordRequester, &apiClient, &requestQueue, &waiter}
}

func setupPasswordRequestForFullRequestQueue() testObject {
	apiClient := setupApiClientForSuccess()
	requestQueue := setupRequestQueueFull()
	supportedHashes := setupSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{client: &apiClient, requestQueue: &requestQueue, supportedHashes: supportedHashes, waiter: &waiter}

	return testObject{&passwordRequester, &apiClient, &requestQueue, &waiter}
}

func setupPasswordRequestFoNoSupportedHashes() testObject {
	apiClient := setupApiClientForSuccess()
	requestQueue := setupRequestQueueForSuccess()
	supportedHashes := setupNoSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{client: &apiClient, requestQueue: &requestQueue, supportedHashes: supportedHashes, waiter: &waiter}

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

func TestGetSupportedHashesCorrectHashes(t *testing.T) {
	expected := map[string]hash.Hash {
		"sha256": sha256.New(),
	}
	actual := getSupportedHashes()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", expectedHash, actual)
	}
}

func TestStartError(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	err := testObject.passwordRequester.Start()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestProcessOrWaitAddRequestToQueueNoError(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	err := testObject.passwordRequester.processOrWait()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestProcessOrWaitAddRequestToQueueNoErrorSizeCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.processOrWait()
	assertRequestQueueSizeCalled(t, testObject)
}

func TestProcessOrWaitAddRequestToQueueNoErrorWaitNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.processOrWait()
	assertWaiterNotCalled(t, testObject)
}

func TestProcessOrWaitAddRequestToQueueError(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	err := testObject.passwordRequester.processOrWait()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestProcessOrWaitAddRequestToQueueErrorSizeCalled(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	testObject.passwordRequester.processOrWait()
	assertRequestQueueSizeCalled(t, testObject)
}

func TestProcessOrWaitAddRequestToQueueErrorWaitNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	testObject.passwordRequester.processOrWait()
	assertWaiterNotCalled(t, testObject)
}

func TestProcessOrWaitRequestQueueFullNoError(t *testing.T) {
	testObject := setupPasswordRequestForFullRequestQueue()
	err := testObject.passwordRequester.processOrWait()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestProcessOrWaitRequestQueueFullSizeCalled(t *testing.T) {
	testObject := setupPasswordRequestForFullRequestQueue()
	testObject.passwordRequester.processOrWait()
	assertRequestQueueSizeCalled(t, testObject)
}

func TestProcessOrWaitRequestQueueFullWaitCalled(t *testing.T) {
	testObject := setupPasswordRequestForFullRequestQueue()
	testObject.passwordRequester.processOrWait()
	assertWaiterCalled(t, testObject)
}

func TestAddRequestToQueueNoError(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	err := testObject.passwordRequester.addRequestToQueue()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestAddRequestToQueueNoErrorPutCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.addRequestToQueue()
	assertRequestQueuePutCalled(t, testObject)
}

func TestAddRequestToQueueGetHashError(t *testing.T) {
	expected := "Unexpected response from api on hash name request with status code: 500\n"
	testObject := setupPasswordRequestForGetHashNameError()
	err := testObject.passwordRequester.addRequestToQueue()

	actual := err.Error()
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestAddRequestToQueueGetHashErrorPutNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForGetHashNameError()
	testObject.passwordRequester.addRequestToQueue()
	assertRequestQueuePutNotCalled(t, testObject)
}

func TestAddRequestToQueueGetPasswordsError(t *testing.T) {
	expected := "Unexpected response from api on password request with status code: 500\n"
	testObject := setupPasswordRequestForGetPasswordsError()
	err := testObject.passwordRequester.addRequestToQueue()

	actual := err.Error()
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestAddRequestToQueueGetPasswordsErrorPutNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForGetPasswordsError()
	testObject.passwordRequester.addRequestToQueue()
	assertRequestQueuePutNotCalled(t, testObject)
}

func TestGetHashNoError(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	_, _, err := testObject.passwordRequester.getHash()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestGetHashRequestHashNameError(t *testing.T) {
	expected := "Unexpected response from api on hash name request with status code: 500\n"
	testObject := setupPasswordRequestForGetHashNameError()
	_, _, err := testObject.passwordRequester.getHash()

	actual := err.Error()
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestGetHashGetHashFunctionError(t *testing.T) {
	expected := "Current hash: sha256 is unsupported\n"
	testObject := setupPasswordRequestFoNoSupportedHashes()
	_, _, err := testObject.passwordRequester.getHash()

	actual := err.Error()
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestGetHashFunctionNoError(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	_, err := testObject.passwordRequester.getHashFunction(expectedHashName)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestGetHashFunctionCorrectHash(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	actual, _ := testObject.passwordRequester.getHashFunction(expectedHashName)
	if !reflect.DeepEqual(expectedHash, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", expectedHash, actual)
	}
}

func TestGetHashFunctionError(t *testing.T) {
	testObject := setupPasswordRequestFoNoSupportedHashes()
	_, err := testObject.passwordRequester.getHashFunction(expectedHashName)
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestRequestHashNameNoError(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	_, err := testObject.passwordRequester.requestHashName()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestRequestHashClientCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.requestHashName()
	assertApiClientGetHashNameCalled(t, testObject)
}

func TestRequestHashNameCorrectHashName(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	actual, _ := testObject.passwordRequester.requestHashName()
	if strings.Compare(expectedHashName, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expectedHashName, actual)
	}
}

func TestRequestHashNameError(t *testing.T) {
	testObject := setupPasswordRequestForGetHashNameError()
	_, err := testObject.passwordRequester.requestHashName()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestGetPasswordsNoError(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	_, err := testObject.passwordRequester.getPasswords()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestGetPasswordsClientCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.getPasswords()
	assertApiClientGetPasswordsCalled(t, testObject)
}

func TestGetPasswordsCorrectPasswords(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	actual, _ := testObject.passwordRequester.getPasswords()
	if !reflect.DeepEqual(expectedPasswords, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", expectedHash, actual)
	}
}

func TestGetPasswordsError(t *testing.T) {
	testObject := setupPasswordRequestForGetPasswordsError()
	_, err := testObject.passwordRequester.getPasswords()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}
