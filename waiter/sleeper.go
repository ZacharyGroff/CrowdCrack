package waiter

import (
	"fmt"
	"runtime"
	"strings"
	"time"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
)

type Sleeper struct {
	logger logger.Logger
	sleepDuration time.Duration
	isLogging bool
}

func NewSleeper(u userinput.CmdLineConfigProvider, l *logger.GenericLogger) Sleeper {
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
	caller := s.getCaller()
	return fmt.Sprintf("%s sleeping for %d seconds", caller, s.sleepDuration / time.Second)
}

func (s Sleeper) getCaller() string {
	pc, _, _, _ := runtime.Caller(3)
	details := runtime.FuncForPC(pc)
	directories := strings.Split(details.Name(), "/")

	return directories[len(directories)-1]
}

