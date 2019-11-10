package interfaces

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type ConfigProvider interface {
	GetConfig() *models.Config
}
