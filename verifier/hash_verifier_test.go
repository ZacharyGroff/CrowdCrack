package verifier

import (
	"errors"
	"reflect"
	"strings"
	"testing"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
)

func TestHashVerifierLoadUserProvidedHashesCorrectHashes(t *testing.T) {
	expected := map[string]bool {
		"fakeHash": true,
	}
	var errorToReturn error
	mockHashReader := mocks.NewMockHashReader(expected, errorToReturn)
	hashVerifier := HashVerifier{hashReader: &mockHashReader}

	hashVerifier.loadUserProvidedHashes()

	actual := hashVerifier.userProvidedHashes 
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", expected, actual)
	}
}


func TestHashVerifierLoadUserProvidedHashesSuccess(t *testing.T) {
	var mapToReturn map[string]bool
	var errorToReturn error
	mockHashReader := mocks.NewMockHashReader(mapToReturn, errorToReturn)
	hashVerifier := HashVerifier{hashReader: &mockHashReader}

	err := hashVerifier.loadUserProvidedHashes()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestHashVerifierLoadUserProvidedHashesError(t *testing.T) {
	var mapToReturn map[string]bool
	errorToReturn := errors.New("test error")
	mockHashReader := mocks.NewMockHashReader(mapToReturn, errorToReturn)
	hashVerifier := HashVerifier{hashReader: &mockHashReader}

	err := hashVerifier.loadUserProvidedHashes()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestHashVerifierGetNextPasswordHashCorrectHash(t *testing.T) {
	expected := "fakeHash"
	var errToReturn error
	mockFlushingQueue := mocks.NewMockFlushingQueue(expected, errToReturn)
	hashVerifier := HashVerifier{computedHashes: &mockFlushingQueue}

	actual := hashVerifier.getNextPasswordHash()
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestHashVerifierParsePasswordHashCorrectPassword(t *testing.T) {
	expected := "fakePassword"
	hash := "fakeHash"
	passwordHash := expected + ":" + hash
	hashVerifier := HashVerifier{}

	actual, _ := hashVerifier.parsePasswordHash(passwordHash)
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestHashVerifierParsePasswordHashCorrectHash(t *testing.T) {
	password := "fakePassword"
	expected := "fakeHash"
	passwordHash := password + ":" + expected
	hashVerifier := HashVerifier{}

	_, actual := hashVerifier.parsePasswordHash(passwordHash)
	if strings.Compare(expected, actual) != 0 {
		t.Errorf("Expected: %s\nActual: %s\n", expected, actual)
	}
}

func TestHashVerifierIsMatchTrue(t *testing.T) {
	expected := true
	hash := "fakeHash"
	userProvidedHashes := map[string]bool {
		hash: true,
	}
	hashVerifier := HashVerifier{userProvidedHashes: userProvidedHashes}

	actual := hashVerifier.isMatch(hash)
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}

func TestHashVerifierIsMatchFalse(t *testing.T) {
	expected := false
	userProvidedHashes := map[string]bool {
		"sha256": true,
	}
	hashVerifier := HashVerifier{userProvidedHashes: userProvidedHashes}

	actual := hashVerifier.isMatch("fakeHash")
	if expected != actual {
		t.Errorf("Expected: %t\nActual: %t\n", expected, actual)
	}
}
