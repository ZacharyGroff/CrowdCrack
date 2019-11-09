package interfaces

type Logger interface {
	LogMessage(string) error
}
