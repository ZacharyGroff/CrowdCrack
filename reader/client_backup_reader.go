package reader

import (
	"encoding/json"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"io/ioutil"
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
	if fileExists(c.config.RequestBackupPath) {
		err := c.loadRequests()
		if err != nil {
			return err
		}
	}

	if fileExists(c.config.SubmissionBackupPath) {
		err := c.loadSubmissions()
		if err != nil {
			return err
		}
	}

	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return false
}

func (c *ClientBackupReader) loadRequests() error {
	jsonBytes, err := getJsonAsBytes(c.config.RequestBackupPath)
	if err != nil {
		return err
	}

	var requests []models.HashingRequest
	json.Unmarshal(jsonBytes, &requests)

	err = c.addRequestsToQueue(requests)
	if err != nil {
		return err
	}

	return nil
}

func (c *ClientBackupReader) loadSubmissions() error {
	jsonBytes, err := getJsonAsBytes(c.config.SubmissionBackupPath)
	if err != nil {
		return err
	}

	var submissions []models.HashSubmission
	json.Unmarshal(jsonBytes, &submissions)

	err = c.addSubmissionsToQueue(submissions)
	if err != nil {
		return err
	}

	return nil
}

func (c *ClientBackupReader) addRequestsToQueue(requests []models.HashingRequest) error {
	for _, request := range requests {
		err := c.requestQueue.Put(request)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ClientBackupReader) addSubmissionsToQueue(submissions []models.HashSubmission) error {
	for _, submission := range submissions {
		err := c.submissionQueue.Put(submission)
		if err != nil {
			return err
		}
	}

	return nil
}

func getJsonAsBytes(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
