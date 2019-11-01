package queue

import (
	"errors"
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type HashingSubmissionQueue struct {
	submissions chan models.HashSubmission
}

func NewHashingSubmissionQueue() *HashingSubmissionQueue{
	s := make(chan models.HashSubmission, 100)
	return &HashingSubmissionQueue{s}
}

func (q HashingSubmissionQueue) Size() int {
	return len(q.submissions)
}

func (q HashingSubmissionQueue) Get() (models.HashSubmission, error) {
	for {
		select{
		case submission := <- q.submissions:
			return submission, nil
		default:
			err := errors.New("No submissions in queue.")
			return models.HashSubmission{}, err
		}
	}
}

func (q HashingSubmissionQueue) Put(submission models.HashSubmission) error {
	select {
	case q.submissions <- submission:
		return nil
	default:
		err := fmt.Errorf("No room in buffer. Discarding submission: %+v\n", submission)
		return err
	}
}
