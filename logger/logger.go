package logger

import (
	"time"
)

type Logger interface {
	LogPasswordsSent(uint64, time.Duration)
	LogHashesComputed(uint64, time.Duration)
	LogHashesCracked(uint64, time.Duration)
	LogMessage(string)
}
