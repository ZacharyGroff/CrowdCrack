package waiter

import (
	"fmt"
	"runtime"
	"time"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
)

type Sleeper struct {
	logger logger.Logger
	sleepDuration time.Duration
	isLogging bool
}

func NewSleeper(u userinput.CmdLineConfigProvider, l logger.Logger) Sleeper {
	config := u.GetConfig()
	sleepDuration := time.Duration(60) * time.Second
	return Sleeper {
		logger:        l,
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

