package userinput

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type ConfigProvider interface {
	GetClientConfig() *models.ClientConfig
	GetServerConfig() *models.ServerConfig
}
