package verifier

import (
	"errors"
	"reflect"
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
