package requester

import (
	"hash"
	"reflect"
	"strings"
	"testing"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
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
var emptyPasswords []string
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

func setupApiClientForNoPasswordsReturned() mocks.MockApiClient {
	return mocks.NewMockApiClient(successCode, successCode, successCode, expectedHashName, emptyPasswords)
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
		"md4": md4.New(),
		"md5": md5.New(),
		"sha1": sha1.New(),
		"sha256": sha256.New(),
		"sha512": sha512.New(),
		"ripemd160": ripemd160.New(),
		"sha3_224": sha3.New224(),
		"sha3_256": sha3.New256(),
		"sha3_384": sha3.New384(),
		"sha3_512": sha3.New512(),
		"sha512_224": sha512.New512_224(),
		"sha512_256": sha512.New512_256(),
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

func setupPasswordRequestFoNoPasswordsReturned() testObject {
	apiClient := setupApiClientForNoPasswordsReturned()
	requestQueue := setupRequestQueueForSuccess()
	supportedHashes := setupSupportedHashes()
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

func TestPasswordRequester_GetSupportedHashes_CorrectResults(t *testing.T) {
	expected := setupSupportedHashes()
	actual := getSupportedHashes()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", expectedHash, actual)
	}
}

func TestPasswordRequester_Start_Error(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	err := testObject.passwordRequester.Start()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestPasswordRequester_ProcessOrSleep_AddRequestToQueue_Success(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	err := testObject.passwordRequester.processOrSleep()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPasswordRequester_ProcessOrSleep_AddRequestToQueue_Success_SizeCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.processOrSleep()
	assertRequestQueueSizeCalled(t, testObject)
}

func TestPasswordRequester_ProcessOrSleep_AddRequestToQueue_Success_WaitNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.processOrSleep()
	assertWaiterNotCalled(t, testObject)
}

func TestPasswordRequester_ProcessOrSleep_AddRequestToQueue_Error(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	err := testObject.passwordRequester.processOrSleep()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestPasswordRequester_ProcessOrSleep_AddRequestToQueue_Error_SizeCalled(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	testObject.passwordRequester.processOrSleep()
	assertRequestQueueSizeCalled(t, testObject)
}

func TestPasswordRequester_ProcessOrSleep_AddRequestToQueue_Error_WaitNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	testObject.passwordRequester.processOrSleep()
	assertWaiterNotCalled(t, testObject)
}

func TestPasswordRequester_ProcessOrSleep_RequestQueueFull_Success(t *testing.T) {
	testObject := setupPasswordRequestForFullRequestQueue()
	err := testObject.passwordRequester.processOrSleep()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPasswordRequester_ProcessOrSleep_RequestQueueFull_SizeCalled(t *testing.T) {
	testObject := setupPasswordRequestForFullRequestQueue()
	testObject.passwordRequester.processOrSleep()
	assertRequestQueueSizeCalled(t, testObject)
}

func TestPasswordRequester_ProcessOrSleep_RequestQueueFull_WaitCalled(t *testing.T) {
	testObject := setupPasswordRequestForFullRequestQueue()
	testObject.passwordRequester.processOrSleep()
	assertWaiterCalled(t, testObject)
}

func TestPasswordRequester_AddRequestToQueue_Success(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	err := testObject.passwordRequester.addRequestToQueue()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPasswordRequester_AddRequestToQueue_Success_PutCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.addRequestToQueue()
	assertRequestQueuePutCalled(t, testObject)
}

func TestPasswordRequester_AddRequestToQueue_Success_WaitNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.addRequestToQueue()
	assertWaiterNotCalled(t, testObject)
}

func TestPasswordRequester_AddRequestToQueue_GetHashError(t *testing.T) {
	expected := "Unexpected response from api on hash name request with status code: 500\n"
	testObject := setupPasswordRequestForGetHashNameError()
	err := testObject.passwordRequester.addRequestToQueue()

	actual := err.Error()
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestPasswordRequester_AddRequestToQueue_GetHashError_PutNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForGetHashNameError()
	testObject.passwordRequester.addRequestToQueue()
	assertRequestQueuePutNotCalled(t, testObject)
}

func TestPasswordRequester_AddRequestToQueue_GetHashError_WaitNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForGetHashNameError()
	testObject.passwordRequester.addRequestToQueue()
	assertWaiterNotCalled(t, testObject)
}

func TestPasswordRequester_AddRequestToQueue_GetPasswords_Error(t *testing.T) {
	expected := "Unexpected response from api on password request with status code: 500\n"
	testObject := setupPasswordRequestForGetPasswordsError()
	err := testObject.passwordRequester.addRequestToQueue()

	actual := err.Error()
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestPasswordRequester_AddRequestToQueue_GetPasswordsError_PutNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForGetPasswordsError()
	testObject.passwordRequester.addRequestToQueue()
	assertRequestQueuePutNotCalled(t, testObject)
}

func TestPasswordRequester_AddRequestToQueue_GetPasswordsError_WaitNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForGetPasswordsError()
	testObject.passwordRequester.addRequestToQueue()
	assertWaiterNotCalled(t, testObject)
}

func TestPasswordRequester_AddRequestToQueue_GetPasswords_NoPasswordsReturned_Success(t *testing.T) {
	testObject := setupPasswordRequestFoNoPasswordsReturned()
	testObject.passwordRequester.addRequestToQueue()
	assertWaiterCalled(t, testObject)
}

func TestPasswordRequester_AddRequestToQueue_GetPasswords_NoPasswordsReturned_PutNotCalled(t *testing.T) {
	testObject := setupPasswordRequestFoNoPasswordsReturned()
	testObject.passwordRequester.addRequestToQueue()
	assertRequestQueuePutNotCalled(t, testObject)
}

func TestPasswordRequester_GetHash_Success(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	_, _, err := testObject.passwordRequester.getHash()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPasswordRequester_GetHash_RequestHashNameError(t *testing.T) {
	expected := "Unexpected response from api on hash name request with status code: 500\n"
	testObject := setupPasswordRequestForGetHashNameError()
	_, _, err := testObject.passwordRequester.getHash()

	actual := err.Error()
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestPasswordRequester_GetHash_GetHashFunctionError(t *testing.T) {
	expected := "Current hash: sha256 is unsupported\n"
	testObject := setupPasswordRequestFoNoSupportedHashes()
	_, _, err := testObject.passwordRequester.getHash()

	actual := err.Error()
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestPasswordRequester_GetHashFunction_Success(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	_, err := testObject.passwordRequester.getHashFunction(expectedHashName)
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPasswordRequester_GetHashFunction_CorrectResults(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	actual, _ := testObject.passwordRequester.getHashFunction(expectedHashName)
	if !reflect.DeepEqual(expectedHash, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", expectedHash, actual)
	}
}

func TestPasswordRequester_GetHashFunction_Error(t *testing.T) {
	testObject := setupPasswordRequestFoNoSupportedHashes()
	_, err := testObject.passwordRequester.getHashFunction(expectedHashName)
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestPasswordRequester_RequestHashName_Success(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	_, err := testObject.passwordRequester.requestHashName()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPasswordRequester_RequestHash_Success_ClientCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.requestHashName()
	assertApiClientGetHashNameCalled(t, testObject)
}

func TestPasswordRequester_RequestHashName_CorrectHashName(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	actual, _ := testObject.passwordRequester.requestHashName()
	if strings.Compare(expectedHashName, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expectedHashName, actual)
	}
}

func TestPasswordRequester_RequestHashName_Error(t *testing.T) {
	testObject := setupPasswordRequestForGetHashNameError()
	_, err := testObject.passwordRequester.requestHashName()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestPasswordRequester_GetPasswords_Success(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	_, err := testObject.passwordRequester.getPasswords()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPasswordRequester_GetPasswords_Success_ClientCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.getPasswords()
	assertApiClientGetPasswordsCalled(t, testObject)
}

func TestPasswordRequester_GetPasswords_CorrectResults(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	actual, _ := testObject.passwordRequester.getPasswords()
	if !reflect.DeepEqual(expectedPasswords, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", expectedHash, actual)
	}
}

func TestPasswordRequester_GetPasswords_Error(t *testing.T) {
	testObject := setupPasswordRequestForGetPasswordsError()
	_, err := testObject.passwordRequester.getPasswords()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}
