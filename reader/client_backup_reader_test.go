package reader

import (
	"crypto/sha256"
	"encoding/json"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"os"
	"testing"
)

var hash = sha256.New()
var hashName = "sha256"

var password1 = "testPassword1"
var password2 = "testPassword2"
var hash1 = "e3b42a03569a4cd0130c703c80885e0aaaadf110c521a941dbc664143a5b1cad"
var hash2 = "008c658a8d900d5a2c829157009b597cfca120d06434dfbca33eab8fd3bedf14"
var result1 = password1 + ":" + hash1
var result2 = password2 + ":" + hash2
var passwords = []string{password1, password2}
var results = []string{result1, result2}

var requestQueueSize = len(passwords)
var submissionQueueSize = len(results)

var hashingRequest = models.HashingRequest{
	Hash:      hash,
	HashName:  hashName,
	Passwords: passwords,
}

var hashSubmission = models.HashSubmission{
	HashType: hashName,
	Results:  nil,
}

type clientBackupReaderTestObject struct {
	clientBackupReader ClientBackupReader
	requestQueue       *mocks.MockRequestQueue
	submissionQueue    *mocks.MockSubmissionQueue
}

func setupRequestQueue() mocks.MockRequestQueue {
	requestQueue := mocks.NewMockRequestQueue(nilError, hashingRequest, requestQueueSize)
	return requestQueue
}

func setupSubmissionQueue() mocks.MockSubmissionQueue {
	requestQueue := mocks.NewMockSubmissionQueue(nilError, hashSubmission, submissionQueueSize)
	return requestQueue
}

func setupClientBackupReader() clientBackupReaderTestObject {
	config := setupConfig()
	requestQueue := setupRequestQueue()
	submissionQueue := setupSubmissionQueue()
	clientBackupReader := ClientBackupReader{
		config:          &config,
		requestQueue:    &requestQueue,
		submissionQueue: &submissionQueue,
	}

	return clientBackupReaderTestObject{
		clientBackupReader: clientBackupReader,
		requestQueue:       &requestQueue,
		submissionQueue:    &submissionQueue,
	}
}

func getHashingRequestsAsString() string {
	var hashingRequests = []models.HashingRequest{hashingRequest}
	bytes, _ := json.Marshal(hashingRequests)
	return string(bytes)
}

func getHashSubmissionsAsString() string {
	var hashSubmissions = []models.HashSubmission{hashSubmission}
	bytes, _ := json.Marshal(hashSubmissions)
	return string(bytes)
}

func assertRequestQueuePutCalledNTimes(t *testing.T, q *mocks.MockRequestQueue, expected uint64) {
	actual := q.PutCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertSubmissionQueuePutCalledNTimes(t *testing.T, q *mocks.MockSubmissionQueue, expected uint64) {
	actual := q.PutCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestClientBackupReader_BackupsExist_True_OneBackupExists(t *testing.T) {
	expected := true

	testObject := setupClientBackupReader()
	lines := []string{"test"}
	setupFile(testObject.clientBackupReader.config.RequestBackupPath, lines)

	actual := testObject.clientBackupReader.BackupsExist()
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}

	os.Remove(testObject.clientBackupReader.config.RequestBackupPath)
}

func TestClientBackupReader_BackupsExist_True_AllBackupsExist(t *testing.T) {
	expected := true

	testObject := setupClientBackupReader()
	lines := []string{"test"}
	setupFile(testObject.clientBackupReader.config.RequestBackupPath, lines)
	setupFile(testObject.clientBackupReader.config.SubmissionBackupPath, lines)

	actual := testObject.clientBackupReader.BackupsExist()
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}

	os.Remove(testObject.clientBackupReader.config.RequestBackupPath)
	os.Remove(testObject.clientBackupReader.config.SubmissionBackupPath)
}

func TestClientBackupReader_BackupsExist_False(t *testing.T) {
	expected := false

	testObject := setupClientBackupReader()

	actual := testObject.clientBackupReader.BackupsExist()
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}

func TestClientBackupReader_LoadBackups_BackupsLoadedReturnsNil(t *testing.T) {
	testObject := setupClientBackupReader()

	lines := []string{"test"}
	setupFile(testObject.clientBackupReader.config.RequestBackupPath, lines)
	setupFile(testObject.clientBackupReader.config.SubmissionBackupPath, lines)

	err := testObject.clientBackupReader.LoadBackups()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestClientBackupReader_LoadBackups_RequestQueuePutCalled(t *testing.T) {
	testObject := setupClientBackupReader()

	hashingRequests := getHashingRequestsAsString()
	lines := []string{hashingRequests}

	expected := uint64(len(lines))
	setupFile(testObject.clientBackupReader.config.RequestBackupPath, lines)

	testObject.clientBackupReader.LoadBackups()
	assertRequestQueuePutCalledNTimes(t, testObject.requestQueue, expected)

	os.Remove(testObject.clientBackupReader.config.RequestBackupPath)
}

func TestClientBackupReader_LoadBackups_SubmissionQueuePutCalled(t *testing.T) {
	testObject := setupClientBackupReader()

	hashSubmissions := getHashSubmissionsAsString()
	lines := []string{hashSubmissions}

	expected := uint64(len(lines))
	setupFile(testObject.clientBackupReader.config.SubmissionBackupPath, lines)

	testObject.clientBackupReader.LoadBackups()
	assertSubmissionQueuePutCalledNTimes(t, testObject.submissionQueue, expected)

	os.Remove(testObject.clientBackupReader.config.SubmissionBackupPath)
}

func TestClientBackupReader_LoadBackups_NoBackupsReturnsNilError(t *testing.T) {
	testObject := setupClientBackupReader()
	err := testObject.clientBackupReader.LoadBackups()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}
