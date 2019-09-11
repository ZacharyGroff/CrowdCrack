package requester

import (
	"fmt"
	"hash"
	"crypto/sha256"
	"encoding/json"
	"net/http"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/queue"
)

type PasswordRequester struct {
	config *config.ClientConfig
	passwords queue.Queue
	supportedHashes map[string]hash.Hash
}

func NewPasswordRequester(c *config.ClientConfig, q *queue.PasswordQueue) *PasswordRequester {
	s := getSupportedHashes()
	return &PasswordRequester{c, q, s}
}

func getSupportedHashes() map[string]hash.Hash {
	return map[string]hash.Hash {
		"sha256": sha256.New(),
	}
}

func (p PasswordRequester) Request() (hash.Hash, error) {
	hash, err := p.requestHash()

	if err != nil {
		return nil, err
	}

	passwords, err := 

	return hash, nil
}

func (p PasswordRequester) getHash() (hash.Hash, error) {
	hashName, err := p.requestHashName()
	if err != nil {
		return err
	}

	currentHash, isSupported := p.supportedHashes[hashName]
	if !isSupported {
		return nil, fmt.Errorf("Current hash: %s is unsupported.", hashName)
	}

	return currentHash, nil
}

func (p PasswordRequester) requestHashName() (string, error) {
	address := p.config.ServerAddress + "/current-hash"
	r, err := http.Get(address)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var hashName string
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&hashName)

	return hashName, err
}
