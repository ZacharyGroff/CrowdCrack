package tracker

type Tracker interface {
	TrackPasswordsSent(uint64)
	TrackHashesComputed(uint64)
	TrackHashesCracked(uint64)
	TrackHashMatchAttempt()
	GetPasswordsSent() uint64
	GetHashesComputed() uint64
	GetHashesCracked() uint64
	GetHashMatchAttempts() uint64
}
