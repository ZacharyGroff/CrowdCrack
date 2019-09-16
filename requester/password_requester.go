package requester

import (
	"fmt"
	"hash"
	"strings"
	"crypto/sha256"
	"encoding/json"
	"net/http"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
)

type PasswordRequester struct {
	config *config.ClientConfig
	supportedHashes map[string]hash.Hash
	requestQueue queue.RequestQueue
}

func NewPasswordRequester(c *config.ClientConfig, r *queue.HashingRequestQueue) *PasswordRequester {
	s := getSupportedHashes()
	return &PasswordRequester{c, s, r}
}

func getSupportedHashes() map[string]hash.Hash {
	return map[string]hash.Hash {
		"sha256": sha256.New(),
	}
}

func (p PasswordRequester) Request() error {
	hash, err := p.getHash()

	if err != nil {
		return err
	}

	passwords, err := p.getPasswords()

	hashingRequest := models.HashingRequest{hash, passwords}
	p.requestQueue.Put(hashingRequest)

	return nil
}

func (p PasswordRequester) getHash() (hash.Hash, error) {
	hashName, err := p.requestHashName()
	if err != nil {
		return nil, err
	}

	currentHash, isSupported := p.supportedHashes[hashName]
	if !isSupported {
		return nil, fmt.Errorf("Current hash: %s is unsupported.", hashName)
	}

	return currentHash, nil
}

func (p PasswordRequester) requestHashName() (string, error) {
	url := p.config.ServerAddress + "/current-hash"
	r, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	var hashName string
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&hashName)

	return hashName, err
}

func (p PasswordRequester) getPasswords() ([]string, error) {
	url := p.config.ServerAddress + "/passwords"
	numPasswords := strings.NewReader("1000")

	var passwords []string
	r, err := http.Post(url, "text/plain", numPasswords)
	if err != nil {
		return passwords, err
	}
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode (&passwords)

	return passwords, err
}
