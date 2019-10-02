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

	err := hashVerifier.loadUserProvidedHashes()
	if err != nil {
		panic(err)
	}

	return &hashVerifier
}

func (v HashVerifier) Verify() {
	for {
		v.verifyNextPasswordHash()
	}
}

func (v *HashVerifier) loadUserProvidedHashes() error {
	userProvidedHashes, err := v.hashReader.GetHashes()
	if err != nil {
		return err
	}
	v.userProvidedHashes = userProvidedHashes

	return nil
}

func (v HashVerifier) verifyNextPasswordHash() bool {
	passwordHash := v.getNextPasswordHash()
	password, hash := v.parsePasswordHash(passwordHash)

	isMatch := v.isMatch(hash)
	if isMatch {
		v.inform(password, hash)
	}

	return isMatch
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
