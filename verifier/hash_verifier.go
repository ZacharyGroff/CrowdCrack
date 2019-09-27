package verifier

import (
	"log"
	"strings"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/reader"
)

type HashVerifier struct {
	computedHashes queue.FlushingQueue
	hashReader reader.HashReader
	userProvidedHashes map[string]bool
}

func NewHashVerifier(q *queue.HashQueue, r *reader.HashlistReader) *HashVerifier {
	hashVerifier := HashVerifier{computedHashes: q, hashReader: r}
	hashVerifier.loadUserProvidedHashes()

	return &hashVerifier
}

func (v HashVerifier) Verify() {
	for {
		v.verifyNextPasswordHash()
	}
}

func (v HashVerifier) loadUserProvidedHashes() {
	userProvidedHashes, err := v.hashReader.GetHashes()
	if err != nil {
		panic(err)
	}
	v.userProvidedHashes = userProvidedHashes
}

func (v HashVerifier) verifyNextPasswordHash() {
	passwordHash := v.getNextPasswordHash()
	password, hash := v.parsePasswordHash(passwordHash)
	if v.isMatch(hash) {
		v.inform(password, hash)
	}
}

func (v HashVerifier) getNextPasswordHash() string {
	for {
		hash, err := v.computedHashes.Get()
		if err == nil {
			return hash
		}
	}	
}

func (v HashVerifier) parsePasswordHash(passwordHash string) (string, string) {
	passwordHashArray := strings.Split(passwordHash, ":")
	return passwordHashArray[0], passwordHashArray[1]
}

func (v HashVerifier) isMatch(hash string) bool {
	if v.userProvidedHashes[hash] {
		return true
	}

	return false
}

func (v HashVerifier) inform(password string, hash string) {
	log.Printf("Hash Cracked: %s\nResult: %s\n", hash, password)
}
