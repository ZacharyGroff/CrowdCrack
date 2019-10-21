package waiter

import (
	"time"
	"github.com/ZacharyGroff/CrowdCrack/logger"
)

type Sleeper struct {
	logger logger.Logger
	sleepDuration time.Duration
	isLogging bool
	logMessage string
}

func NewSleeper(sleepSeconds int, isLogging bool, logMessage string, logger logger.Logger) Sleeper {
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
