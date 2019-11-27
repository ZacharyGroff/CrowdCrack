package reader

import (
	"bufio"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"os"
)

type HashlistReader struct {
	config *models.Config
}

func NewHashlistReader(p interfaces.ConfigProvider) interfaces.HashReader {
	c := p.GetConfig()
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
