package waiter

import (
	"log"
	"time"
)

type Sleeper struct {
	sleepDuration time.Duration
	isLogging bool
	logMessage string
}

func NewSleeper(sleepSeconds int, isLogging bool, logMessage string) Sleeper {
	sleepDuration := time.Duration(sleepSeconds) * time.Second
	return Sleeper{sleepDuration, isLogging, logMessage}
}

func (s Sleeper) Wait() {
	if s.isLogging {
		log.Println(s.logMessage)
	}
	time.Sleep(s.sleepDuration)
}
