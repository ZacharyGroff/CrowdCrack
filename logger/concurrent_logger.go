package logger

import (
	"bufio"
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
	"log"
	"os"
	"sync"
	"time"
)

type ConcurrentLogger struct {
	config *models.Config
	mux    sync.Mutex
}

func NewConcurrentLogger(p userinput.CmdLineConfigProvider) *ConcurrentLogger {
	c := p.GetConfig()
	return &ConcurrentLogger{
		config: c,
	}
}

func (s *ConcurrentLogger) LogMessage(logMessage string) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.config.Verbose {
		log.Println(logMessage)
	}
	err := s.logToFile(logMessage)

	return err
}

func (s *ConcurrentLogger) logToFile(logMessage string) error {
	timeFormattedMessage := getTimeFormattedMessage(time.Now(), logMessage)
	file, err := os.OpenFile(s.config.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, timeFormattedMessage)

	return writer.Flush()
}

func getTimeFormattedMessage(currentTime time.Time, logMessage string) string {
	timeFormatted := currentTime.Format(time.RFC822)
	return fmt.Sprintf("%s: %s", timeFormatted, logMessage)
}
