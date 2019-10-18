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

func setupHashVerifierForSuccess() *HashVerifier {
	mockFlushingQueue := mocks.NewMockFlushingQueue(fakePasswordHash, nilError)
	mockHashReader := mocks.NewMockHashReader(hashMap, nilError)
	mockLogger := mocks.NewMockLogger(nilError)
	mockTracker := mocks.NewMockTracker(42)
	return &HashVerifier{
		computedHashes:     &mockFlushingQueue,
		hashReader:         &mockHashReader,
		logger:             &mockLogger,
		tracker:            &mockTracker,
		userProvidedHashes: hashMap,
	}
}

func setupHashVerifierForNoMatch() *HashVerifier {
	mockFlushingQueue := mocks.NewMockFlushingQueue(fakePasswordHash, nilError)
	mockHashReader := mocks.NewMockHashReader(nil, nilError)
	mockLogger := mocks.NewMockLogger(nilError)
	mockTracker := mocks.NewMockTracker(42)
	return &HashVerifier{
		computedHashes:     &mockFlushingQueue,
		hashReader:         &mockHashReader,
		logger:             &mockLogger,
		tracker:            &mockTracker,
		userProvidedHashes: nil,
	}
}

func setupHashVerifierForHashReaderError() *HashVerifier {
	mockHashReader := mocks.NewMockHashReader(hashMap, testError)
	mockLogger := mocks.NewMockLogger(nilError)
	mockTracker := mocks.NewMockTracker(42)
	return &HashVerifier{
		computedHashes:     nil,
		hashReader:         &mockHashReader,
		logger:             &mockLogger,
		tracker:            &mockTracker,
		userProvidedHashes: nil,
	}
}

func TestHashVerifierLoadUserProvidedHashesCorrectHashes(t *testing.T) {
	hashVerifier := setupHashVerifierForSuccess()
	hashVerifier.loadUserProvidedHashes()

	actual := hashVerifier.userProvidedHashes 
	if !reflect.DeepEqual(hashMap, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", hashMap, actual)
	}
}

func TestHashVerifierLoadUserProvidedHashesSuccess(t *testing.T) {
	hashVerifier := setupHashVerifierForSuccess()

	err := hashVerifier.loadUserProvidedHashes()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHashVerifierLoadUserProvidedHashesError(t *testing.T) {
	hashVerifier := setupHashVerifierForHashReaderError()

	err := hashVerifier.loadUserProvidedHashes()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHashVerifierVerifyNextPasswordHashIsMatch(t *testing.T) {
	expected := true

	hashVerifier := setupHashVerifierForSuccess()

	actual := hashVerifier.verifyNextPasswordHash()
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}

func TestHashVerifierVerifyNextPasswordHashIsNotMatch(t *testing.T) {
	expected := false

	hashVerifier := setupHashVerifierForNoMatch()

	actual := hashVerifier.verifyNextPasswordHash()
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}

func TestHashVerifierGetNextPasswordHashCorrectHash(t *testing.T) {
	hashVerifier := setupHashVerifierForSuccess()

	actual := hashVerifier.getNextPasswordHash()
	if strings.Compare(fakePasswordHash, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", fakeHash, actual)
	}
}

func TestHashVerifierParsePasswordHashCorrectPassword(t *testing.T) {
	hashVerifier := setupHashVerifierForSuccess()

	actual, _ := hashVerifier.parsePasswordHash(fakePasswordHash)
	if strings.Compare(fakePassword, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", fakePassword, actual)
	}
}

func TestHashVerifierParsePasswordHashCorrectHash(t *testing.T) {
	hashVerifier := setupHashVerifierForSuccess()

	_, actual := hashVerifier.parsePasswordHash(fakePasswordHash)
	if strings.Compare(fakeHash, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", fakeHash, actual)
	}
}

func TestHashVerifierIsMatchTrue(t *testing.T) {
	expected := true

	hashVerifier := setupHashVerifierForSuccess()

	actual := hashVerifier.isMatch(fakeHash)
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}

func TestHashVerifierIsMatchFalse(t *testing.T) {
	expected := false

	hashVerifier := setupHashVerifierForNoMatch()

	actual := hashVerifier.isMatch("sha256")
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}
