package logger

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
)

type GenericLogger struct {
	config *models.Config
}

func NewGenericLogger(p userinput.CmdLineConfigProvider) *GenericLogger {
	c := p.GetConfig()
	return &GenericLogger{
		config: c,
	}
}

func (s *GenericLogger) LogMessage(logMessage string) error {
	if s.config.Verbose {
		log.Println(logMessage)
	}
	err := s.logToFile(logMessage)

	return err
}

func (s *GenericLogger) logToFile(logMessage string) error {
	timeFormattedMessage := getTimeFormattedMessage(time.Now(), logMessage)
	file, err := os.OpenFile(s.config.LogPath, os.O_WRONLY | os.O_CREATE | os.O_APPEND, os.ModePerm)
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
