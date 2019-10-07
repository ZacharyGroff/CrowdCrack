package userinput

import (
	"github.com/ZacharyGroff/CrowdCrack/config"
)

type InputParser interface {
	ParseClient() *config.ClientConfig
	ParseServer() *config.ServerConfig
}
