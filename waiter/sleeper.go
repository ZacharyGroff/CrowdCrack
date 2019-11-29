package waiter

import (
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"runtime"
	"strings"
	"time"
)

type Sleeper struct {
	logger        interfaces.Logger
	sleepDuration time.Duration
	isLogging     bool
}

func NewSleeper(u interfaces.ConfigProvider, l interfaces.Logger) interfaces.Waiter {
	config := u.GetConfig()
	sleepDuration := time.Duration(5) * time.Second
	return Sleeper{
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
	return fmt.Sprintf("%s sleeping for %d seconds", caller, s.sleepDuration/time.Second)
}

func (s Sleeper) getCaller() string {
	pc, _, _, _ := runtime.Caller(3)
	details := runtime.FuncForPC(pc)
	directories := strings.Split(details.Name(), "/")

	return directories[len(directories)-1]
}
