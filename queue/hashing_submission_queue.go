package queue

import (
	"errors"
	"fmt"
)

type HashingSubmissionQueue struct {
	submissions chan uint64
}

func NewHashingSubmissionQueue() *HashingSubmissionQueue{
	s := make(chan uint64, 2)
	return &HashingSubmissionQueue{s}
}

func (q HashingSubmissionQueue) Get() (uint64, error) {
	for {
		select{
		case submission := <- q.submissions:
			return submission, nil
		default:
			err := errors.New("No submissions in queue.")
			return 0, err
		}
	}
}

func (q HashingSubmissionQueue) Put(submission uint64) error {
	select {
	case q.submissions <- submission:
		return nil
	default:
		err := fmt.Errorf("No room in buffer. Discarding submission: %d\n", submission)
		return err
	}
}
