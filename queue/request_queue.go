package queue

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type RequestQueue interface {
	Size() int
	Get() (models.HashingRequest, error)
	Put(models.HashingRequest) error
}
