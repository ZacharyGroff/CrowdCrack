package queue

import (
	"github.com/ZacharyGroff/CrowdCrack/mocks"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"testing"
)

var threads = uint16(3)
var stopReason = models.ClientStopReason{
	Requester:"",
	Encoder:"testReason",
	Submitter:"",
}

func setupConfig() models.Config {
	return models.Config{
		Threads: threads,
	}
}

func setupClientStopReasonQueue() *ClientStopReasonQueue {
	config := setupConfig()
	configProvider := mocks.NewMockConfigProvider(&config)
	return NewClientStopReasonQueue(&configProvider)
}

func fillQueueToCapacity(c *ClientStopReasonQueue) {
	var i uint16
	for i = 0; i < threads - 1; i++ {
		c.Put(stopReason)
	}
}

func TestNewClientStopReasonQueue_CorrectChannelBufferSize(t *testing.T) {
	expected := threads - 1
	queue := setupClientStopReasonQueue()

	actual := uint16(cap(queue.stopReasons))
	if expected != actual {
		t.Errorf("Expected: %d\nActual: %d\n", expected, actual)
	}
}

func TestClientStopReasonQueue_Get_Success(t *testing.T) {
	queue := setupClientStopReasonQueue()
	queue.stopReasons <- stopReason

	_, err := queue.Get()
	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestClientStopReasonQueue_Get_Error(t *testing.T) {
	queue := setupClientStopReasonQueue()

	_, err := queue.Get()
	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestClientStopReasonQueue_Get_CorrectResult(t *testing.T) {
	expected := stopReason
	queue := setupClientStopReasonQueue()
	queue.stopReasons <- expected

	actual, _ := queue.Get()
	if expected != actual {
		t.Errorf("Expected: %+v\nActual: %+v\n", expected, actual)
	}
}

func TestClientStopReasonQueue_Put_Success(t *testing.T) {
	queue := setupClientStopReasonQueue()
	err := queue.Put(stopReason)

	if err != nil {
		t.Errorf("Unexpected error returned: %s\n", err.Error())
	}
}

func TestClientStopReasonQueue_Put_Error(t *testing.T) {
	queue := setupClientStopReasonQueue()

	fillQueueToCapacity(queue)
	err := queue.Put(stopReason)

	if err == nil {
		t.Error("Expected error but nil returned")
	}
}

func TestClientStopReasonQueue_Put_CorrectResult(t *testing.T) {
	expected := stopReason

	queue := setupClientStopReasonQueue()
	queue.Put(expected)

	actual := <- queue.stopReasons
	if expected != actual {
		t.Errorf("Expected: %+v\nActual: %+v\n", expected, actual)
	}
}
