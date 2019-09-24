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

func NewSleeper(s int, i bool, l string) Sleeper {
	d := time.Duration(s) * time.Second
	return Sleeper{d, i, l}
}

func (s Sleeper) Wait() {
	if s.isLogging {
		log.Println(s.logMessage)
	}
	time.Sleep(s.sleepDuration)
}
