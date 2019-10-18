package verifier

import (
	"errors"
	"reflect"
	"strings"
	"testing"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
)

var hashMap = map[string]bool {
	fakeHash: true,
}
var fakePassword = "fakePassword"
var fakeHash = "fakeHash"
var fakePasswordHash = fakePassword + ":" + fakeHash
var testError = errors.New("test error")
var nilError error

type testObject struct {
	hashVerifier *HashVerifier
	mockFlushingQueue *mocks.MockFlushingQueue
	mockHashReader *mocks.MockHashReader
	mockLogger *mocks.MockLogger
	mockTracker *mocks.MockTracker
}

func setupHashVerifierForSuccess() testObject {
	mockFlushingQueue := mocks.NewMockFlushingQueue(fakePasswordHash, nilError)
	mockHashReader := mocks.NewMockHashReader(hashMap, nilError)
	mockLogger := mocks.NewMockLogger(nilError)
	mockTracker := mocks.NewMockTracker(42)
	hashVerifier := HashVerifier{
		computedHashes:     &mockFlushingQueue,
		hashReader:         &mockHashReader,
		logger:             &mockLogger,
		tracker:            &mockTracker,
		userProvidedHashes: hashMap,
	}

	return testObject {
		hashVerifier:      &hashVerifier,
		mockFlushingQueue: &mockFlushingQueue,
		mockHashReader:    &mockHashReader,
		mockLogger:        &mockLogger,
		mockTracker:       &mockTracker,
	}
}

func setupHashVerifierForNoMatch() testObject {
	mockFlushingQueue := mocks.NewMockFlushingQueue(fakePasswordHash, nilError)
	mockHashReader := mocks.NewMockHashReader(nil, nilError)
	mockLogger := mocks.NewMockLogger(nilError)
	mockTracker := mocks.NewMockTracker(42)
	hashVerifier := HashVerifier{
		computedHashes:     &mockFlushingQueue,
		hashReader:         &mockHashReader,
		logger:             &mockLogger,
		tracker:            &mockTracker,
		userProvidedHashes: nil,
	}

	return testObject {
		hashVerifier:      &hashVerifier,
		mockFlushingQueue: &mockFlushingQueue,
		mockHashReader:    &mockHashReader,
		mockLogger:        &mockLogger,
		mockTracker:       &mockTracker,
	}
}

func setupHashVerifierForHashReaderError() testObject {
	mockFlushingQueue := mocks.NewMockFlushingQueue(fakePasswordHash, nilError)
	mockHashReader := mocks.NewMockHashReader(hashMap, testError)
	mockLogger := mocks.NewMockLogger(nilError)
	mockTracker := mocks.NewMockTracker(42)
	hashVerifier := HashVerifier{
		computedHashes:     nil,
		hashReader:         &mockHashReader,
		logger:             &mockLogger,
		tracker:            &mockTracker,
		userProvidedHashes: nil,
	}

	return testObject {
		hashVerifier:      &hashVerifier,
		mockFlushingQueue: &mockFlushingQueue,
		mockHashReader:    &mockHashReader,
		mockLogger:        &mockLogger,
		mockTracker:       &mockTracker,
	}
}

func assertTrackerCalled(t *testing.T, m *mocks.MockTracker) {
	expected := uint64(1)
	actual := m.TrackHashesCrackedCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertLoggerCalled(t *testing.T, m *mocks.MockLogger) {
	expected := uint64(1)
	actual := m.LogMessageCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertTrackerNotCalled(t *testing.T, m *mocks.MockTracker) {
	expected := uint64(0)
	actual := m.TrackHashesCrackedCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestHashVerifierLoadUserProvidedHashesCorrectHashes(t *testing.T) {
	testObject := setupHashVerifierForSuccess()
	testObject.hashVerifier.loadUserProvidedHashes()

	actual := testObject.hashVerifier.userProvidedHashes
	if !reflect.DeepEqual(hashMap, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", hashMap, actual)
	}
}

func TestHashVerifierLoadUserProvidedHashesSuccess(t *testing.T) {
	testObject := setupHashVerifierForSuccess()

	err := testObject.hashVerifier.loadUserProvidedHashes()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHashVerifierLoadUserProvidedHashesError(t *testing.T) {
	testObject := setupHashVerifierForHashReaderError()

	err := testObject.hashVerifier.loadUserProvidedHashes()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHashVerifierVerifyNextPasswordHashIsMatch(t *testing.T) {
	expected := true

	testObject := setupHashVerifierForSuccess()

	actual := testObject.hashVerifier.verifyNextPasswordHash()
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}

func TestHashVerifier_VerifyNextPassword_HashIsMatch_TrackerCalled(t *testing.T) {
	testObject := setupHashVerifierForSuccess()
	testObject.hashVerifier.verifyNextPasswordHash()
	assertTrackerCalled(t, testObject.mockTracker)
}

func TestHashVerifierVerifyNextPasswordHashIsNotMatch(t *testing.T) {
	expected := false

	testObject := setupHashVerifierForNoMatch()

	actual := testObject.hashVerifier.verifyNextPasswordHash()
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}

func TestHashVerifier_VerifyNextPassword_HashIsNotMatch_TrackerNotCalled(t *testing.T) {
	testObject := setupHashVerifierForNoMatch()
	testObject.hashVerifier.verifyNextPasswordHash()
	assertTrackerNotCalled(t, testObject.mockTracker)
}

func TestHashVerifierGetNextPasswordHashCorrectHash(t *testing.T) {
	testObject := setupHashVerifierForSuccess()

	actual := testObject.hashVerifier.getNextPasswordHash()
	if strings.Compare(fakePasswordHash, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", fakeHash, actual)
	}
}

func TestHashVerifierParsePasswordHashCorrectPassword(t *testing.T) {
	testObject := setupHashVerifierForSuccess()

	actual, _ := testObject.hashVerifier.parsePasswordHash(fakePasswordHash)
	if strings.Compare(fakePassword, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", fakePassword, actual)
	}
}

func TestHashVerifierParsePasswordHashCorrectHash(t *testing.T) {
	testObject := setupHashVerifierForSuccess()

	_, actual := testObject.hashVerifier.parsePasswordHash(fakePasswordHash)
	if strings.Compare(fakeHash, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", fakeHash, actual)
	}
}

func TestHashVerifierIsMatchTrue(t *testing.T) {
	expected := true

	testObject := setupHashVerifierForSuccess()

	actual := testObject.hashVerifier.isMatch(fakeHash)
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}

func TestHashVerifierIsMatchFalse(t *testing.T) {
	expected := false

	testObject := setupHashVerifierForNoMatch()

	actual := testObject.hashVerifier.isMatch("sha256")
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}

func TestHashVerifier_Inform_LoggerCalled(t *testing.T) {
	testObject := setupHashVerifierForSuccess()
	testObject.hashVerifier.inform(fakePassword, fakeHash)
	assertLoggerCalled(t, testObject.mockLogger)
}
