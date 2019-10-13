package logger

type Logger interface {
	TrackPasswordsSent(int)
	TrackHashesReceived(int)
	TrackHashCracked(int)
	LogMessage(string)
}
