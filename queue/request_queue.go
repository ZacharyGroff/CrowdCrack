package queue

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type RequestQueue interface {
	Get() (models.HashingRequest, error)
	Put(models.HashingRequest) error
}
