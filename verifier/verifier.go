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
		hash := v.getNextHash()
		if v.isMatch(hash) {
			v.inform(hash)
		}
	}
}

func (v Verifier) getNextHash() string {
	for {
		hash, err := v.computedHashes.Get()
		if err == nil {
			return hash
		}
	}	
}

func (v Verifier) isMatch(hash string) bool {
	hashAndPassword := strings.Split(hash, ":")
	hashValue := hashAndPassword[0]
	if v.userProvidedHashes[hashValue] {
		return true
	}

	return false
}

func (v Verifier) inform(hash string) {
	log.Printf("Password Cracked: %s\n", hash)
}
