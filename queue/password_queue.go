package queue

import (
	"errors"
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type PasswordQueue struct {
	passwords chan string
	config    *models.Config
}

func NewPasswordQueue(p interfaces.ConfigProvider) interfaces.Queue {
	config := p.GetConfig()
	passwords := make(chan string, config.PasswordQueueBuffer)
	return &PasswordQueue{passwords, config}
}

func (q PasswordQueue) Size() int {
	return len(q.passwords)
}

func (q PasswordQueue) Get() (string, error) {
	for {
		select {
		case password := <-q.passwords:
			return password, nil
		default:
			err := errors.New("No passwords in queue.")
			return "", err
		}
	}
}

func (q PasswordQueue) Put(password string) error {
	select {
	case q.passwords <- password:
		return nil
	default:
		err := fmt.Errorf("No room in buffer. Discarding password: %q\n", password)
		return err
	}
}
