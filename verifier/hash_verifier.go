package verifier

import (
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/reader"
	"github.com/ZacharyGroff/CrowdCrack/tracker"
	"strings"
)

type HashVerifier struct {
	computedHashes     interfaces.FlushingQueue
	hashReader         interfaces.HashReader
	logger             interfaces.Logger
	tracker            interfaces.Tracker
	userProvidedHashes map[string]bool
}

func NewHashVerifier(q *queue.HashQueue, r *reader.HashlistReader, l *logger.ConcurrentLogger, t *tracker.StatsTracker) *HashVerifier {
	return &HashVerifier{
		computedHashes: q,
		hashReader:     r,
		logger:         l,
		tracker:        t,
	}
}

func (v HashVerifier) Start() {
	err := v.loadUserProvidedHashes()
	if err != nil {
		panic(err)
	}

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
		v.tracker.TrackHashesCracked(1)
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
	v.tracker.TrackHashMatchAttempt()
	if v.userProvidedHashes[hash] {
		return true
	}

	return false
}

func (v HashVerifier) inform(password string, hash string) {
	logMessage := fmt.Sprintf("Hash Cracked: %s Result: %s", hash, password)
	v.logger.LogMessage(logMessage)
}
