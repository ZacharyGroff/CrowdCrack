package reader

import (
	"bufio"
	"os"
	"github.com/ZacharyGroff/CrowdCrack/config"
)

type HashlistReader struct {
	config *config.ServerConfig
}

func NewHashlistReader(c *config.ServerConfig) *HashlistReader {
	return &HashlistReader{c}
}

func (h HashlistReader) GetHashes() (map[string]bool, error) {
	hashes := make(map[string]bool)
	file, err := os.Open(h.config.HashlistPath)
	if err != nil {
		return hashes, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hash := scanner.Text()
		hashes[hash] = true
		if err != nil {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return hashes, err
	}
	
	return hashes, nil
}
