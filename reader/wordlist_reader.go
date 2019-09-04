package reader

import (
	"bufio"
	"os"
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
	file, err := os.Open(w.config.WordlistPath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		password := scanner.Text()
		err := w.passwords.Put(password)
		if err != nil {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	
	return nil
}
