package logger

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	file, err := os.OpenFile(s.config.LogPath, os.O_WRONLY | os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, logMessage)

	return writer.Flush()
}
