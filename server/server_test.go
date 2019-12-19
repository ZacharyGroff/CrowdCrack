package server

import (
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"reflect"
	"testing"
	"time"
)

type testObject struct {
	server             Server
	mockApi            *mocks.MockApi
	mockLogger         *mocks.MockLogger
	mockPasswordReader *mocks.MockPasswordReader
	mockObserver       *mocks.MockObserver
	mockVerifier       *mocks.MockVerifier
}

func setupServerForNoError() testObject {
	mockApi := mocks.MockApi{0}
	mockLogger := mocks.NewMockLogger(nil)
	mockPasswordReader := mocks.NewMockPasswordReader(false)
	mockObserver := mocks.NewMockObserver()
	mockVerifier := mocks.MockVerifier{0}

	server := Server{&mockApi, &mockLogger, &mockPasswordReader, &mockObserver, &mockVerifier}

	return testObject{server, &mockApi, &mockLogger, &mockPasswordReader, &mockObserver, &mockVerifier}
}

func setupServerForPasswordReaderError() testObject {
	mockApi := mocks.MockApi{0}
	mockLogger := mocks.NewMockLogger(nil)
	mockPasswordReader := mocks.NewMockPasswordReader(true)
	mockObserver := mocks.NewMockObserver()
	mockVerifier := mocks.MockVerifier{0}
	server := Server{&mockApi, &mockLogger, &mockPasswordReader, &mockObserver, &mockVerifier}

	return testObject{server, &mockApi, &mockLogger, &mockPasswordReader, &mockObserver, &mockVerifier}
}

func assertLoadPasswordsCalled(t *testing.T, p *mocks.MockPasswordReader) {
	time.Sleep(100 * time.Millisecond)
	expected := uint64(1)
	actual := p.LoadPasswordsCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertLoggerCalled(t *testing.T, l *mocks.MockLogger) {
	expected := uint64(1)
	actual := l.LogMessageCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertLoggerCalledTimes(t *testing.T, l *mocks.MockLogger, expected uint64) {
	actual := l.LogMessageCalls
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func assertObserverCalled(t *testing.T, o *mocks.MockObserver) {
	time.Sleep(100 * time.Millisecond)
	expected := uint64(1)
	actual := o.StartCalls
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

func assertNoError(t *testing.T, testObject testObject) {
	assertLoggerCalled(t, testObject.mockLogger)
	assertLoadPasswordsCalled(t, testObject.mockPasswordReader)
	assertObserverCalled(t, testObject.mockObserver)
	assertVerifyCalled(t, testObject.mockVerifier)
	assertHandleRequestsCalled(t, testObject.mockApi)
}

func assertError(t *testing.T, testObject testObject) {
	time.Sleep(100 * time.Millisecond)
	assertLoggerCalledTimes(t, testObject.mockLogger, 2)
	assertLoadPasswordsCalled(t, testObject.mockPasswordReader)
}

func TestNewServer(t *testing.T) {
	testObject := setupServerForNoError()
	expected := testObject.server

	actual := NewServer(expected.Api, expected.Logger, expected.Reader, expected.Observer, expected.Verifier)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %+v\nActual: %+v\n", expected, actual)
	}
}

func TestServer_Start_Success(t *testing.T) {
	testObject := setupServerForNoError()
	testObject.server.Start()
	assertNoError(t, testObject)
}

func TestServer_Start_LoadPasswords_Error(t *testing.T) {
	testObject := setupServerForPasswordReaderError()
	testObject.server.Start()
	assertError(t, testObject)
}
