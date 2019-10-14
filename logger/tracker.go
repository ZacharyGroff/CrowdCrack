package logger

type Tracker interface {
	TrackPasswordsSent(uint64)
	TrackHashesComputed(uint64)
	TrackHashesCracked(uint64)
}
