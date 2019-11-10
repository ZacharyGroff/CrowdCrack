package queue

import (
	"errors"
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type HashingRequestQueue struct {
	requests chan models.HashingRequest
}

func NewHashingRequestQueue() *HashingRequestQueue {
	r := make(chan models.HashingRequest, 10)
	return &HashingRequestQueue{r}
}

func (q HashingRequestQueue) Size() int {
	return len(q.requests)
}

func (q HashingRequestQueue) Get() (models.HashingRequest, error) {
	for {
		select {
		case request := <-q.requests:
			return request, nil
		default:
			err := errors.New("No requests in queue.")
			return models.HashingRequest{}, err
		}
	}
}

func (q HashingRequestQueue) Put(request models.HashingRequest) error {
	select {
	case q.requests <- request:
		return nil
	default:
		err := fmt.Errorf("No room in buffer. Discarding request: %+v\n", request)
		return err
	}
}
