package queue

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"github.com/ZacharyGroff/CrowdCrack/config"
)

type HashQueue struct {
	hashes chan string
	config config.QueueConfig
}

func NewServerHashQueue(config *config.ServerConfig) *HashQueue {
	hashes := make(chan string, config.GetHashQueueBuffer())
	return &HashQueue{hashes, *config}
}

func NewClientHashQueue(config *config.ClientConfig) *HashQueue {
	hashes := make(chan string, config.GetHashQueueBuffer())
	return &HashQueue{hashes, *config}
}

func (q HashQueue) Size() int {
	return len(q.hashes)
}

func (q HashQueue) Get() (string, error) {
	for {
		select {
		case hash := <- q.hashes:
			return hash, nil
		default:
			err := errors.New("No Urls in queue.")
			return "", err
		}
	}
}

func (q HashQueue) Put(hash string) error {
	select {
	case q.hashes <- hash:
		return nil
	default:
		err := fmt.Errorf("No room in buffer. Discarding hash: %+v\n", hash)
		return err
	}
}

func (q HashQueue) Flush() error {
	if q.config.GetFlushToFile() {
		return q.flushToFile()
	}
	_, err := q.emptyChannel()
	
	return err
}

func (q HashQueue) flushToFile() error {
	file, err := os.OpenFile(q.config.GetComputedHashOverflowPath(), os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	hashes, err := q.emptyChannel()
	if err != nil {
		return err
	}

	for _, hash := range hashes {
		fmt.Fprintln(writer, hash) 
	}

	return writer.Flush()
}

func (q HashQueue) emptyChannel() ([]string, error) {
	initialSize := len(q.hashes)
	var hashes []string
	for i := 0; i < initialSize; i++ {
		hash, err := q.Get()
		if err != nil {
			return nil, err
		}

		hashes = append(hashes, hash)
	}

	return hashes, nil
}
