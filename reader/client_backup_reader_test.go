package reader

import (
	"crypto/sha256"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
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

var request = models.HashingRequest{
	Hash:      hash,
	HashName:  hashName,
	Passwords: passwords,
}

var submission = models.HashSubmission{
	HashType: hashName,
	Results:  nil,
}

type clientBackupReaderTestObject struct {
	clientBackupReader ClientBackupReader
	requestQueue       *interfaces.RequestQueue
	submissionQueue    *interfaces.SubmissionQueue
}

func setupRequestQueue() interfaces.RequestQueue {
	requestQueue := mocks.NewMockRequestQueue(nilError, request, requestQueueSize)
	return &requestQueue
}

func setupSubmissionQueue() interfaces.SubmissionQueue {
	requestQueue := mocks.NewMockSubmissionQueue(nilError, submission, submissionQueueSize)
	return &requestQueue
}

func setupClientBackupReader() clientBackupReaderTestObject {
	config := setupConfig()
	requestQueue := setupRequestQueue()
	submissionQueue := setupSubmissionQueue()
	clientBackupReader := ClientBackupReader{
		config:          &config,
		requestQueue:    requestQueue,
		submissionQueue: submissionQueue,
	}

	return clientBackupReaderTestObject{
		clientBackupReader: clientBackupReader,
		requestQueue:       &requestQueue,
		submissionQueue:    &submissionQueue,
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
