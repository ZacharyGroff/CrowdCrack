package requester

import (
	"crypto/sha256"
	"errors"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"hash"
	"reflect"
	"strings"
	"testing"
)

type testObject struct {
	passwordRequester *PasswordRequester
	apiClient         *mocks.MockApiClient
	logger            *mocks.MockLogger
	requestQueue      *mocks.MockRequestQueue
	stopQueue         *mocks.MockClientStopQueue
	waiter            *mocks.MockWaiter
}

var expectedPasswords = []string{
	"hunter2",
	"password123",
}
var threads = 42
var emptyPasswords []string
var expectedHashName = "sha256"
var expectedHash = sha256.New()
var successCode = 200
var errorCode = 500
var nilError error
var testError = errors.New("testError")

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
	hashingRequest := models.HashingRequest{}
	return mocks.NewMockRequestQueue(nilError, hashingRequest, 0)
}

func setupRequestQueueFull() mocks.MockRequestQueue {
	hashingRequest := models.HashingRequest{}
	return mocks.NewMockRequestQueue(testError, hashingRequest, 10)
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
	return mocks.NewMockClientStopQueue(stopReason, testError, testError)
}

func setupSupportedHashes() map[string]hash.Hash {
	return models.GetSupportedHashFunctions()
}

func setupNoSupportedHashes() map[string]hash.Hash {
	return map[string]hash.Hash{}
}

func setupVerboseConfig() models.Config {
	return models.Config{
		PasswordRequestSize: 1,
		Verbose:             true,
		Threads:             uint16(threads),
	}
}

func setupNonVerboseConfig() models.Config {
	return models.Config{
		PasswordRequestSize: 1,
		Verbose:             false,
	}
}

func setupLogger() mocks.MockLogger {
	return mocks.NewMockLogger(nilError)
}

func setupPasswordRequestForSuccess() testObject {
	apiClient := setupApiClientForSuccess()
	config := setupVerboseConfig()
	logger := setupLogger()
	requestQueue := setupRequestQueueForSuccess()
	stopQueue := setupStopQueueForEmptyReturn()
	supportedHashes := setupSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{
		config:          &config,
		client:          &apiClient,
		logger:          &logger,
		requestQueue:    &requestQueue,
		stopQueue:       &stopQueue,
		supportedHashes: supportedHashes,
		waiter:          &waiter,
	}

	return testObject{
		passwordRequester: &passwordRequester,
		apiClient:         &apiClient,
		logger:            &logger,
		requestQueue:      &requestQueue,
		stopQueue:         &stopQueue,
		waiter:            &waiter,
	}
}

func setupPasswordRequestForStopQueueMessage() testObject {
	apiClient := setupApiClientForSuccess()
	config := setupVerboseConfig()
	logger := setupLogger()
	requestQueue := setupRequestQueueForSuccess()
	stopQueue := setupStopQueueForStopReasonReturn()
	supportedHashes := setupSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{
		config:          &config,
		client:          &apiClient,
		logger:          &logger,
		requestQueue:    &requestQueue,
		stopQueue:       &stopQueue,
		supportedHashes: supportedHashes,
		waiter:          &waiter,
	}

	return testObject{
		passwordRequester: &passwordRequester,
		apiClient:         &apiClient,
		logger:            &logger,
		requestQueue:      &requestQueue,
		stopQueue:         &stopQueue,
		waiter:            &waiter,
	}
}

func setupPasswordRequestForSuccessNonVerbose() testObject {
	apiClient := setupApiClientForSuccess()
	config := setupNonVerboseConfig()
	logger := setupLogger()
	requestQueue := setupRequestQueueForSuccess()
	stopQueue := setupStopQueueForEmptyReturn()
	supportedHashes := setupSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{
		config:          &config,
		client:          &apiClient,
		logger:          &logger,
		requestQueue:    &requestQueue,
		stopQueue:       &stopQueue,
		supportedHashes: supportedHashes,
		waiter:          &waiter,
	}

	return testObject{
		passwordRequester: &passwordRequester,
		apiClient:         &apiClient,
		logger:            &logger,
		requestQueue:      &requestQueue,
		stopQueue:         &stopQueue,
		waiter:            &waiter,
	}
}

func setupPasswordRequestForApiClientError() testObject {
	apiClient := setupApiClientForError()
	config := setupVerboseConfig()
	logger := setupLogger()
	requestQueue := setupRequestQueueForSuccess()
	stopQueue := setupStopQueueForEmptyReturn()
	supportedHashes := setupSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{
		config:          &config,
		client:          &apiClient,
		logger:          &logger,
		requestQueue:    &requestQueue,
		stopQueue:       &stopQueue,
		supportedHashes: supportedHashes,
		waiter:          &waiter,
	}

	return testObject{
		passwordRequester: &passwordRequester,
		apiClient:         &apiClient,
		logger:            &logger,
		requestQueue:      &requestQueue,
		stopQueue:         &stopQueue,
		waiter:            &waiter,
	}
}

func setupPasswordRequestForGetHashNameError() testObject {
	apiClient := setupApiClientForGetHashNameError()
	config := setupVerboseConfig()
	logger := setupLogger()
	requestQueue := setupRequestQueueForSuccess()
	stopQueue := setupStopQueueForEmptyReturn()
	supportedHashes := setupSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{
		config:          &config,
		client:          &apiClient,
		logger:          &logger,
		requestQueue:    &requestQueue,
		stopQueue:       &stopQueue,
		supportedHashes: supportedHashes,
		waiter:          &waiter,
	}

	return testObject{
		passwordRequester: &passwordRequester,
		apiClient:         &apiClient,
		logger:            &logger,
		requestQueue:      &requestQueue,
		stopQueue:         &stopQueue,
		waiter:            &waiter,
	}
}

func setupPasswordRequestForGetPasswordsError() testObject {
	apiClient := setupApiClientForGetPasswordsError()
	config := setupVerboseConfig()
	logger := setupLogger()
	requestQueue := setupRequestQueueForSuccess()
	stopQueue := setupStopQueueForEmptyReturn()
	supportedHashes := setupSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{
		config:          &config,
		client:          &apiClient,
		logger:          &logger,
		requestQueue:    &requestQueue,
		stopQueue:       &stopQueue,
		supportedHashes: supportedHashes,
		waiter:          &waiter,
	}

	return testObject{
		passwordRequester: &passwordRequester,
		apiClient:         &apiClient,
		logger:            &logger,
		requestQueue:      &requestQueue,
		stopQueue:         &stopQueue,
		waiter:            &waiter,
	}
}

func setupPasswordRequestForFullRequestQueue() testObject {
	apiClient := setupApiClientForSuccess()
	config := setupVerboseConfig()
	logger := setupLogger()
	requestQueue := setupRequestQueueFull()
	stopQueue := setupStopQueueForEmptyReturn()
	supportedHashes := setupSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{
		config:          &config,
		client:          &apiClient,
		logger:          &logger,
		requestQueue:    &requestQueue,
		stopQueue:       &stopQueue,
		supportedHashes: supportedHashes,
		waiter:          &waiter,
	}

	return testObject{
		passwordRequester: &passwordRequester,
		apiClient:         &apiClient,
		logger:            &logger,
		requestQueue:      &requestQueue,
		stopQueue:         &stopQueue,
		waiter:            &waiter,
	}
}

func setupPasswordRequestFoNoSupportedHashes() testObject {
	apiClient := setupApiClientForSuccess()
	config := setupVerboseConfig()
	logger := setupLogger()
	requestQueue := setupRequestQueueForSuccess()
	stopQueue := setupStopQueueForEmptyReturn()
	supportedHashes := setupNoSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{
		config:          &config,
		client:          &apiClient,
		logger:          &logger,
		requestQueue:    &requestQueue,
		stopQueue:       &stopQueue,
		supportedHashes: supportedHashes,
		waiter:          &waiter,
	}

	return testObject{
		passwordRequester: &passwordRequester,
		apiClient:         &apiClient,
		logger:            &logger,
		requestQueue:      &requestQueue,
		stopQueue:         &stopQueue,
		waiter:            &waiter,
	}
}

func setupPasswordRequestFoNoPasswordsReturned() testObject {
	apiClient := setupApiClientForNoPasswordsReturned()
	config := setupVerboseConfig()
	logger := setupLogger()
	requestQueue := setupRequestQueueForSuccess()
	stopQueue := setupStopQueueForEmptyReturn()
	supportedHashes := setupSupportedHashes()
	waiter := mocks.NewMockWaiter()
	passwordRequester := PasswordRequester{
		config:          &config,
		client:          &apiClient,
		logger:          &logger,
		requestQueue:    &requestQueue,
		stopQueue:       &stopQueue,
		supportedHashes: supportedHashes,
		waiter:          &waiter,
	}

	return testObject{
		passwordRequester: &passwordRequester,
		apiClient:         &apiClient,
		logger:            &logger,
		requestQueue:      &requestQueue,
		stopQueue:         &stopQueue,
		waiter:            &waiter,
	}
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

func assertLoggerCalledOnce(t *testing.T, testObject testObject) {
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

func TestPasswordRequester_Start_ProcessOrWait_Error(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	err := testObject.passwordRequester.Start()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestPasswordRequester_Start_ProcessOrWait_Error_LoggerCalled(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	testObject.passwordRequester.Start()
	assertLoggerCalledOnce(t, testObject)
}

func TestPasswordRequester_Start_ProcessOrWait_Error_StopQueueGetCalled(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	testObject.passwordRequester.Start()
	assertStopQueueGetCalled(t, testObject)
}

func TestPasswordRequester_Start_ProcessOrWait_Error_StopQueuePutCalled(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	testObject.passwordRequester.Start()
	assertStopQueuePutCalledNTimes(t, testObject, uint64(threads - 1))
}

func TestPasswordRequester_Start_StopQueue_Error(t *testing.T) {
	testObject := setupPasswordRequestForStopQueueMessage()
	err := testObject.passwordRequester.Start()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestPasswordRequester_Start_StopQueue_Error_StopQueueCalled(t *testing.T) {
	testObject := setupPasswordRequestForStopQueueMessage()
	testObject.passwordRequester.Start()
	assertStopQueueGetCalled(t, testObject)
}

func TestPasswordRequester_Start_StopQueue_Error_LoggerCalled(t *testing.T) {
	testObject := setupPasswordRequestForStopQueueMessage()
	testObject.passwordRequester.Start()
	assertLoggerCalledOnce(t, testObject)
}

func TestPasswordRequester_ProcessOrWait_Process_Success(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	err := testObject.passwordRequester.processOrWait()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPasswordRequester_ProcessOrWait_Process_Success_SizeCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.processOrWait()
	assertRequestQueueSizeCalled(t, testObject)
}

func TestPasswordRequester_ProcessOrWait_Process_Success_WaitNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.processOrWait()
	assertWaiterNotCalled(t, testObject)
}

func TestPasswordRequester_ProcessOrWait_Process_Error(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	err := testObject.passwordRequester.processOrWait()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestPasswordRequester_ProcessOrWait_Process_Error_SizeCalled(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	testObject.passwordRequester.processOrWait()
	assertRequestQueueSizeCalled(t, testObject)
}

func TestPasswordRequester_ProcessOrWait_Process_Error_WaitNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForApiClientError()
	testObject.passwordRequester.processOrWait()
	assertWaiterNotCalled(t, testObject)
}

func TestPasswordRequester_ProcessOrWait_RequestQueueFull_Success(t *testing.T) {
	testObject := setupPasswordRequestForFullRequestQueue()
	err := testObject.passwordRequester.processOrWait()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPasswordRequester_ProcessOrWait_RequestQueueFull_SizeCalled(t *testing.T) {
	testObject := setupPasswordRequestForFullRequestQueue()
	testObject.passwordRequester.processOrWait()
	assertRequestQueueSizeCalled(t, testObject)
}

func TestPasswordRequester_ProcessOrWait_RequestQueueFull_WaitCalled(t *testing.T) {
	testObject := setupPasswordRequestForFullRequestQueue()
	testObject.passwordRequester.processOrWait()
	assertWaiterCalled(t, testObject)
}

func TestPasswordRequester_Process_Success(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	err := testObject.passwordRequester.process()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPasswordRequester_Process_Success_PutCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.process()
	assertRequestQueuePutCalled(t, testObject)
}

func TestPasswordRequester_Process_Success_WaitNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.process()
	assertWaiterNotCalled(t, testObject)
}

func TestPasswordRequester_Process_Success_Verbose_LoggerCalledOnce(t *testing.T) {
	testObject := setupPasswordRequestForSuccess()
	testObject.passwordRequester.process()
	assertLoggerCalledOnce(t, testObject)
}

func TestPasswordRequester_Process_Success_NonVerbose_LoggerNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForSuccessNonVerbose()
	testObject.passwordRequester.process()
	assertLoggerNotCalled(t, testObject)
}

func TestPasswordRequester_Process_GetHashError(t *testing.T) {
	expected := "Unexpected response from api on hash name request with status code: 500\n"
	testObject := setupPasswordRequestForGetHashNameError()
	err := testObject.passwordRequester.process()

	actual := err.Error()
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestPasswordRequester_Process_GetHashError_PutNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForGetHashNameError()
	testObject.passwordRequester.process()
	assertRequestQueuePutNotCalled(t, testObject)
}

func TestPasswordRequester_Process_GetHashError_WaitNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForGetHashNameError()
	testObject.passwordRequester.process()
	assertWaiterNotCalled(t, testObject)
}

func TestPasswordRequester_Process_GetPasswords_Error(t *testing.T) {
	expected := "Unexpected response from api on password request with status code: 500\n"
	testObject := setupPasswordRequestForGetPasswordsError()
	err := testObject.passwordRequester.process()

	actual := err.Error()
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestPasswordRequester_Process_GetPasswordsError_PutNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForGetPasswordsError()
	testObject.passwordRequester.process()
	assertRequestQueuePutNotCalled(t, testObject)
}

func TestPasswordRequester_Process_GetPasswordsError_WaitNotCalled(t *testing.T) {
	testObject := setupPasswordRequestForGetPasswordsError()
	testObject.passwordRequester.process()
	assertWaiterNotCalled(t, testObject)
}

func TestPasswordRequester_Process_GetPasswords_NoPasswordsReturned_Success(t *testing.T) {
	testObject := setupPasswordRequestFoNoPasswordsReturned()
	err := testObject.passwordRequester.process()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestPasswordRequester_Process_GetPasswords_NoPasswordsReturned_WaitCalled(t *testing.T) {
	testObject := setupPasswordRequestFoNoPasswordsReturned()
	testObject.passwordRequester.process()
	assertWaiterCalled(t, testObject)
}

func TestPasswordRequester_Process_GetPasswords_NoPasswordsReturned_LoggerCalled(t *testing.T) {
	testObject := setupPasswordRequestFoNoPasswordsReturned()
	testObject.passwordRequester.process()
	assertLoggerCalledOnce(t, testObject)
}

func TestPasswordRequester_Process_GetPasswords_NoPasswordsReturned_PutNotCalled(t *testing.T) {
	testObject := setupPasswordRequestFoNoPasswordsReturned()
	testObject.passwordRequester.process()
	assertRequestQueuePutNotCalled(t, testObject)
}

func TestPasswordRequester_Process_RequestQueuePut_Error(t *testing.T) {
	testObject := setupPasswordRequestForFullRequestQueue()
	err := testObject.passwordRequester.process()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
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
