package server

import (
	"testing"
	"time"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
)

type testObject struct {
	server *Server
	mockApi *mocks.MockApi
	mockPasswordReader *mocks.MockPasswordReader
	mockVerifier *mocks.MockVerifier	
}

func setupServerForNoError() testObject {
	mockApi := mocks.MockApi{0}
	mockPasswordReader := mocks.NewMockPasswordReader(false)
	mockVerifier := mocks.MockVerifier{0}
	server := Server{&mockApi, &mockPasswordReader, &mockVerifier}
	
	return testObject{&server, &mockApi, &mockPasswordReader, &mockVerifier}
}

func setupServerForError() testObject {
	mockApi := mocks.MockApi{0}
	mockPasswordReader := mocks.NewMockPasswordReader(true)
	mockVerifier := mocks.MockVerifier{0}
	server := Server{&mockApi, &mockPasswordReader, &mockVerifier}

	return testObject{&server, &mockApi, &mockPasswordReader, &mockVerifier}
}

func assertLoadPasswordsCalled(t *testing.T, p *mocks.MockPasswordReader) {
	expected := uint64(1)
	actual := p.LoadPasswordsCalls 
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertVerifyCalled(t *testing.T, v *mocks.MockVerifier) {
	time.Sleep(100 * time.Millisecond)
	expected := uint64(1)
	actual := v.VerifyCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertHandleRequestsCalled(t *testing.T, a *mocks.MockApi) {
	expected := uint64(1)
	actual := a.HandleRequestsCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertVerifyNotCalled(t *testing.T, v *mocks.MockVerifier) {
	time.Sleep(100 * time.Millisecond)
	expected := uint64(0)
	actual := v.VerifyCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertHandleRequestsNotCalled(t *testing.T, a *mocks.MockApi) {
	expected := uint64(0)
	actual := a.HandleRequestsCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertNoError(t *testing.T, testObject testObject) {
	assertLoadPasswordsCalled(t, testObject.mockPasswordReader)
	assertVerifyCalled(t, testObject.mockVerifier)
	assertHandleRequestsCalled(t, testObject.mockApi)
}

func assertError(t *testing.T, testObject testObject) {
	assertLoadPasswordsCalled(t, testObject.mockPasswordReader)
	assertVerifyNotCalled(t, testObject.mockVerifier)
	assertHandleRequestsNotCalled(t, testObject.mockApi)
}

func recoverAndAssertError(t *testing.T, testObject testObject) {
	recover()
	assertError(t, testObject)
}

func TestServerStartNoError(t *testing.T) {
	testObject := setupServerForNoError()
	testObject.server.Start()
	assertNoError(t, testObject)
}

func TestServerStartLoadPasswordsError(t *testing.T) {
	testObject := setupServerForError()
	defer recoverAndAssertError(t, testObject)
	testObject.server.Start()
}
