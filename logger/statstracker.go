package logger

type StatsTracker interface {
	TrackPasswordsSent(int)
	TrackHashesComputed(int)
	TrackHashCracked(int)
}
