package waiter

import (
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"time"
)

type Sleeper struct {
	logger logger.Logger
	sleepDuration time.Duration
	isLogging bool
	logMessage string
}

func NewSleeper(sleepSeconds int, isLogging bool, logMessage string, logger *logger.GenericLogger) Sleeper {
	sleepDuration := time.Duration(sleepSeconds) * time.Second
	return Sleeper {
		logger:        logger,
		sleepDuration: sleepDuration,
		isLogging:     isLogging,
		logMessage:    logMessage,
	}
}

func (s Sleeper) Wait() {
	if s.isLogging {
		s.logger.LogMessage(s.logMessage)
	}
	time.Sleep(s.sleepDuration)
}
