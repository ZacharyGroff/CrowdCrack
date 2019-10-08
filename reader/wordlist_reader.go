package reader

import (
	"bufio"
	"os"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
)

type WordlistReader struct {
	config *models.ServerConfig
	passwords queue.Queue	
}

func NewWordlistReader(p userinput.CmdLineConfigProvider, q *queue.PasswordQueue) *WordlistReader {
	c := p.GetServerConfig()
	return &WordlistReader{c, q}
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
