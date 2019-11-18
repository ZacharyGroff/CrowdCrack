package queue

import (
	"errors"
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type ClientStopReasonQueue struct {
	stopReasons chan models.ClientStopReason
}

func NewClientStopReasonQueue(p interfaces.ConfigProvider) *ClientStopReasonQueue {
	config := p.GetConfig()
	stopReasons := getStopReasonsChannel(config)
	return &ClientStopReasonQueue{
		stopReasons: stopReasons,
	}
}

func getStopReasonsChannel(config *models.Config) chan models.ClientStopReason {
	buffer := config.Threads - 1
	stopReasons := make(chan models.ClientStopReason, buffer)
	return stopReasons
}

func (c ClientStopReasonQueue) Get() (models.ClientStopReason, error) {
	select {
	case reason := <-c.stopReasons:
		return reason, nil
	default:
		err := errors.New("No reasons in queue.")
		return models.ClientStopReason{}, err
	}
}

func (c ClientStopReasonQueue) Put(reason models.ClientStopReason) error {
	select {
	case c.stopReasons <- reason:
		return nil
	default:
		err := fmt.Errorf("No room in buffer. Discarding reason: %+v\n", reason)
		return err
	}
}
