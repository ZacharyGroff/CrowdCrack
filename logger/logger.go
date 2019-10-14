package logger

type Logger interface {
	LogMessage(string) error
}
