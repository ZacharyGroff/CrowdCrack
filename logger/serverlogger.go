package logger

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type ServerLogger struct {
	config *models.ServerConfig
}

func NewServerLogger(c *models.ServerConfig) *ServerLogger {
	return &ServerLogger{c}
}

func (s *ServerLogger) LogMessage(logMessage string) error {
	if s.config.Verbose {
		log.Println(logMessage)
	}
	err := s.logToFile(logMessage)

	return err
}

func (s *ServerLogger) logToFile(logMessage string) error {
	file, err := os.OpenFile(s.config.LogPath, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, logMessage)

	return writer.Flush()
}
