package interfaces

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type SubmissionQueue interface {
	Size() int
	Get() (models.HashSubmission, error)
	Put(models.HashSubmission) error
}
