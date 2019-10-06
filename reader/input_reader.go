package reader

import (
	"github.com/ZacharyGroff/CrowdCrack/config"
)

type InputReader interface {
	Read([]string) *config.ClientConfig
}
