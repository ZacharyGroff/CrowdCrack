package queue

import (
	"errors"
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/config"
)

type PasswordQueue struct {
	passwords chan string
	config *config.ServerConfig
}

func NewPasswordQueue(config *config.ServerConfig) *PasswordQueue {
	passwords := make(chan string, config.PasswordQueueBuffer)
	return &PasswordQueue{passwords, config}
}

func (q PasswordQueue) Size() int {
	return len(q.passwords)
}

func (q PasswordQueue) Get() (string, error) {
	for {
		select {
		case password := <- q.passwords:
			return password, nil
		default:
			err := errors.New("No Urls in queue.")
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
