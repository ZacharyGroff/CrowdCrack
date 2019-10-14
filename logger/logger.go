package logger

import (
	"time"
)

type Logger interface {
	LogPasswordsSent(int, time.Duration)
	LogHashesComputed(int, time.Duration)
	LogHashCracked(string)
	LogMessage(string)
}
