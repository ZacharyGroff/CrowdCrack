package reader

import (
	"bufio"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"os"
	"time"
)

type WordlistReader struct {
	config    *models.Config
	passwords interfaces.Queue
}

func NewWordlistReader(p interfaces.ConfigProvider, q interfaces.Queue) interfaces.PasswordReader {
	return &WordlistReader{
		config:    p.GetConfig(),
		passwords: q,
	}
}

func (w WordlistReader) LoadPasswords() error {
	file, err := os.Open(w.config.WordlistPath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	return w.populateQueueFromScanner(scanner)
}

func (w *WordlistReader) populateQueueFromScanner(scanner *bufio.Scanner) error {
	for scanner.Scan() {
		password := scanner.Text()
		w.putInQueueOrWait(password)
	}

	return scanner.Err()
}

func (w *WordlistReader) putInQueueOrWait(password string) {
	err := w.passwords.Put(password)
	if err != nil {
		time.Sleep(15 * time.Second)
		w.putInQueueOrWait(password)
	}
}
