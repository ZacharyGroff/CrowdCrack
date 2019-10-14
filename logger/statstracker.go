package logger

type StatsTracker struct {
	PasswordsSent uint64
	HashesComputed uint64
	HashesCracked uint64
}

func (s *StatsTracker) TrackPasswordsSent(passwordsSent uint64) {
	s.PasswordsSent += passwordsSent
}

func (s *StatsTracker) TrackHashesComputed(hashesComputed uint64) {
	s.HashesComputed += hashesComputed
}

func (s *StatsTracker) TrackHashesCracked(hashesCracked uint64) {
	s.HashesCracked += hashesCracked
}



