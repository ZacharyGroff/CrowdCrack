package verifier

import (
	"log"
	"strings"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/reader"
)

type Verifier struct {
	config *config.ServerConfig
	computedHashes queue.FlushingQueue
	userProvidedHashes map[string]bool
}

func NewVerifier(c *config.ServerConfig, q *queue.HashQueue, r *reader.HashlistReader) *Verifier {
	u, err := r.GetHashes()
	if err != nil {
		panic(err)
	}

	return &Verifier{c, q, u}
}

func (v Verifier) Verify() {
	for {
		passwordHash := v.getNextPasswordHash()
		password, hash := v.parsePasswordHash(passwordHash)
		if v.isMatch(hash) {
			v.inform(password, hash)
		}
	}
}

func (v Verifier) getNextPasswordHash() string {
	for {
		hash, err := v.computedHashes.Get()
		if err == nil {
			return hash
		}
	}	
}

func (v Verifier) parsePasswordHash(passwordHash string) (string, string) {
	passwordHashArray := strings.Split(passwordHash, ":")
	return passwordHashArray[0], passwordHashArray[1]
}

func (v Verifier) isMatch(hash string) bool {
	if v.userProvidedHashes[hash] {
		return true
	}

	return false
}

func (v Verifier) inform(password string, hash string) {
	log.Printf("Hash Cracked: %s\nResult: %s\n", hash, password)
}
