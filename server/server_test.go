package server

import (
	"testing"
	"time"
)

type testObject struct {
	server *Server
	mockApi *mockApi
	mockPasswordReader *mockPasswordReader
	mockVerifier *mockVerifier	
}

func setupServerForNoError() testObject {
	mockApi := mockApi{0}
	mockPasswordReader := mockPasswordReader{0, false}
	mockVerifier := mockVerifier{0}
	server := Server{&mockApi, &mockPasswordReader, &mockVerifier}
	
	return testObject{&server, &mockApi, &mockPasswordReader, &mockVerifier}
}

func setupServerForError() testObject {
	mockApi := mockApi{0}
	mockPasswordReader := mockPasswordReader{0, true}
	mockVerifier := mockVerifier{0}
	server := Server{&mockApi, &mockPasswordReader, &mockVerifier}

	return testObject{&server, &mockApi, &mockPasswordReader, &mockVerifier}
}

func assertLoadPasswordsCalled(t *testing.T, p *mockPasswordReader) {
	expected := uint64(1)
	actual := p.loadPasswordsCalls 
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertVerifyCalled(t *testing.T, v *mockVerifier) {
	time.Sleep(100 * time.Millisecond)
	expected := uint64(1)
	actual := v.verifyCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertHandleRequestsCalled(t *testing.T, a *mockApi) {
	expected := uint64(1)
	actual := a.handleRequestsCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertVerifyNotCalled(t *testing.T, v *mockVerifier) {
	time.Sleep(100 * time.Millisecond)
	expected := uint64(0)
	actual := v.verifyCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertHandleRequestsNotCalled(t *testing.T, a *mockApi) {
	expected := uint64(0)
	actual := a.handleRequestsCalls
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
