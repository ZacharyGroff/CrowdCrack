package reader

import (
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/queue"
)

type WordlistReader struct {
	config *config.ServerConfig
	passwords queue.Queue	
}

func NewWordlistReader(c *config.ServerConfig, p *queue.PasswordQueue) *WordlistReader {
	return &WordlistReader{c, p}
}

func (w WordlistReader) LoadPasswords() error {
	return nil
}
