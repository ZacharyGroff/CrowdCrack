package reader

import (
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"os"
)

type ClientBackupReader struct {
	config          *models.Config
	requestQueue    interfaces.RequestQueue
	submissionQueue interfaces.SubmissionQueue
}

func NewClientBackupReader(p interfaces.ConfigProvider, r interfaces.RequestQueue, s interfaces.SubmissionQueue) interfaces.BackupReader {
	return &ClientBackupReader{
		config:          p.GetConfig(),
		requestQueue:    r,
		submissionQueue: s,
	}
}

func (c *ClientBackupReader) BackupsExist() bool {
	requestBackupExists := fileExists(c.config.RequestBackupPath)
	submissionBackupExists := fileExists(c.config.SubmissionBackupPath)

	return requestBackupExists || submissionBackupExists
}

func (c *ClientBackupReader) LoadBackups() error {
	panic("implement me")
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return false
}
