package interfaces

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type ClientStopQueue interface {
	Get() (models.ClientStopReason, error)
	Put(models.ClientStopReason) error
}
