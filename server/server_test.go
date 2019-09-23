package server

import (
	"testing"
	"time"
)

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


func TestServerStartNoError(t *testing.T) {
	mockApi := mockApi{0}
	mockPasswordReader := mockPasswordReader{0, false}
	mockVerifier := mockVerifier{0}
	server := Server{&mockApi, &mockPasswordReader, &mockVerifier}

	server.Start()

	assertLoadPasswordsCalled(t, &mockPasswordReader)
	assertVerifyCalled(t, &mockVerifier)
	assertHandleRequestsCalled(t, &mockApi)
}

func TestServerStartLoadPasswordsError(t *testing.T) {
	mockApi := mockApi{0}
	mockPasswordReader := mockPasswordReader{0, true}
	mockVerifier := mockVerifier{0}
	server := Server{&mockApi, &mockPasswordReader, &mockVerifier}

	defer func() {
		if r := recover(); r != nil {
			assertLoadPasswordsCalled(t, &mockPasswordReader)
			assertVerifyNotCalled(t, &mockVerifier)
			assertHandleRequestsNotCalled(t, &mockApi)
		}
	}()
	
	server.Start()
}
