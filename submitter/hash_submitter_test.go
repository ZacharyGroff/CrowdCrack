package submitter

import (
	"errors"
	"testing"
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

func TestProcessSubmissionSuccess(t *testing.T) {
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(error(nil), models.HashSubmission{})
	mockApiClient := mocks.NewMockApiClient(200, "fakeHash", []string{})
	hashSubmitter := HashSubmitter{submissionQueue: &mockSubmissionQueue, client: &mockApiClient}

	err := hashSubmitter.processSubmission()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestProcessSubmissionSuccessCorrectSubmissionQueueCalls(t *testing.T) {
	expected := uint64(1)

	mockSubmissionQueue := mocks.NewMockSubmissionQueue(error(nil), models.HashSubmission{})
	mockApiClient := mocks.NewMockApiClient(200, "fakeHash", []string{})
	hashSubmitter := HashSubmitter{submissionQueue: &mockSubmissionQueue, client: &mockApiClient}

	hashSubmitter.processSubmission()

	actual := mockSubmissionQueue.GetCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func TestProcessSubmissionSuccessCorrectClientCalls(t *testing.T) {
	expected := uint64(1)

	mockSubmissionQueue := mocks.NewMockSubmissionQueue(error(nil), models.HashSubmission{})
	mockApiClient := mocks.NewMockApiClient(200, "fakeHash", []string{})
	hashSubmitter := HashSubmitter{submissionQueue: &mockSubmissionQueue, client: &mockApiClient}

	hashSubmitter.processSubmission()

	actual := mockApiClient.SubmitHashesCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func TestProcessSubmissionClientBadStatusCodeReturned(t *testing.T) {
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(error(nil), models.HashSubmission{})
	mockApiClient := mocks.NewMockApiClient(500, "fakeHash", []string{})
	hashSubmitter := HashSubmitter{submissionQueue: &mockSubmissionQueue, client: &mockApiClient}

	err := hashSubmitter.processSubmission()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestProcessSubmissionClientBadStatusCodeReturnedCorrectSubmissionQueueCalls(t *testing.T) {
	expected := uint64(1)

	mockSubmissionQueue := mocks.NewMockSubmissionQueue(error(nil), models.HashSubmission{})
	mockApiClient := mocks.NewMockApiClient(500, "fakeHash", []string{})
	hashSubmitter := HashSubmitter{submissionQueue: &mockSubmissionQueue, client: &mockApiClient}

	hashSubmitter.processSubmission()

	actual := mockSubmissionQueue.GetCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func TestProcessSubmissionClientBadStatusCodeReturnedCorrectClientCalls(t *testing.T) {
	expected := uint64(1)

	mockSubmissionQueue := mocks.NewMockSubmissionQueue(error(nil), models.HashSubmission{})
	mockApiClient := mocks.NewMockApiClient(500, "fakeHash", []string{})
	hashSubmitter := HashSubmitter{submissionQueue: &mockSubmissionQueue, client: &mockApiClient}

	hashSubmitter.processSubmission()

	actual := mockApiClient.SubmitHashesCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func TestProcessSubmissionSubmissionQueueErrorReturned(t *testing.T) {
	mockSubmissionQueue := mocks.NewMockSubmissionQueue(errors.New("test error"), models.HashSubmission{})
	mockApiClient := mocks.NewMockApiClient(500, "fakeHash", []string{})
	hashSubmitter := HashSubmitter{submissionQueue: &mockSubmissionQueue, client: &mockApiClient}

	err := hashSubmitter.processSubmission()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestProcessSubmissionSubmissionQueueErrorCorrectSubmissionQueueCalls(t *testing.T) {
	expected := uint64(1)

	mockSubmissionQueue := mocks.NewMockSubmissionQueue(errors.New("test error"), models.HashSubmission{})
	mockApiClient := mocks.NewMockApiClient(500, "fakeHash", []string{})
	hashSubmitter := HashSubmitter{submissionQueue: &mockSubmissionQueue, client: &mockApiClient}

	hashSubmitter.processSubmission()

	actual := mockSubmissionQueue.GetCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}

func TestProcessSubmissionSubmissionQueueErrorCorrectClientCalls(t *testing.T) {
	expected := uint64(0)

	mockSubmissionQueue := mocks.NewMockSubmissionQueue(errors.New("test error"), models.HashSubmission{})
	mockApiClient := mocks.NewMockApiClient(500, "fakeHash", []string{})
	hashSubmitter := HashSubmitter{submissionQueue: &mockSubmissionQueue, client: &mockApiClient}

	hashSubmitter.processSubmission()

	actual := mockApiClient.SubmitHashesCalls
	if expected != actual {
		t.Errorf("Expected %d\nActual: %d\n", expected, actual)
	}
}