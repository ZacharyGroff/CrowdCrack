package tracker

type StatsTracker struct {
	passwordsSent uint64
	hashesComputed uint64
	hashesCracked uint64
	hashMatchAttempts uint64
}

func NewStatsTracker() *StatsTracker {
	return &StatsTracker{
		passwordsSent:     0,
		hashesComputed:    0,
		hashesCracked:     0,
		hashMatchAttempts: 0,
	}
}

func (s *StatsTracker) TrackPasswordsSent(passwordsSent uint64) {
	s.passwordsSent += passwordsSent
}

func (s *StatsTracker) TrackHashesComputed(hashesComputed uint64) {
	s.hashesComputed += hashesComputed
}

func (s *StatsTracker) TrackHashesCracked(hashesCracked uint64) {
	s.hashesCracked += hashesCracked
}

func (s *StatsTracker) TrackHashMatchAttempt() {
	s.hashesCracked++
}

func (s *StatsTracker) GetPasswordsSent() uint64 {
	return s.passwordsSent
}

func (s *StatsTracker) GetHashesComputed() uint64 {
	return s.hashesComputed
}

func (s *StatsTracker) GetHashesCracked() uint64 {
	return s.hashesCracked
}

func (s *StatsTracker) GetHashMatchAttempts() uint64 {
	return s.hashesCracked
}
