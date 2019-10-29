package waiter

import (
	"fmt"
	"runtime"
	"time"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/logger"
)

type Sleeper struct {
	logger logger.Logger
	sleepDuration time.Duration
	isLogging bool
}

func NewSleeper(config *models.Config, logger logger.Logger) Sleeper {
	sleepDuration := time.Duration(60) * time.Second
	return Sleeper {
		logger:        logger,
		sleepDuration: sleepDuration,
		isLogging:     config.Verbose,
	}
}

func (s Sleeper) Wait() {
	if s.isLogging {
		logMessage := s.getLogMessage()
		s.logger.LogMessage(logMessage)
	}
	time.Sleep(s.sleepDuration)
}

func (s Sleeper) getLogMessage() string {
	pc, _, _, _ := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	return fmt.Sprintf("%s sleeping for %d seconds", details.Name(), s.sleepDuration / time.Second)
}

