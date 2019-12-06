package flusher

import (
	"encoding/json"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"io/ioutil"
	"os"
)

type ClientQueueFlusher struct {
	config          *models.Config
	requestQueue    interfaces.RequestQueue
	submissionQueue interfaces.SubmissionQueue
}

func NewClientQueueFlusher(p interfaces.ConfigProvider, r interfaces.RequestQueue, s interfaces.SubmissionQueue) interfaces.Flusher {
	return &ClientQueueFlusher{
		config:          p.GetConfig(),
		requestQueue:    r,
		submissionQueue: s,
	}
}

func (c *ClientQueueFlusher) NeedsFlushed() bool {
	return c.requestQueue.Size() > 0  || c.submissionQueue.Size() > 0
}

func (c *ClientQueueFlusher) Flush() error {
	if c.requestQueue.Size() > 0 {
		err := c.flushRequestQueue()
		if err != nil {
			return err
		}
	}

	if c.submissionQueue.Size() > 0 {
		err := c.flushSubmissionQueue()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ClientQueueFlusher) flushRequestQueue() error {
	requests, err := c.emptyRequestQueue()
	if err != nil {
		return err
	}

	requestsJson, err := json.Marshal(&requests)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("request_queue_backup.json", requestsJson, os.ModePerm)

	return err
}

func (c *ClientQueueFlusher) flushSubmissionQueue() error {
	submissions, err := c.emptySubmissionQueue()
	if err != nil {
		return err
	}

	submissionsJson, err := json.Marshal(&submissions)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("submission_queue_backup.json", submissionsJson, os.ModePerm)

	return err
}

func (c *ClientQueueFlusher) emptyRequestQueue() ([]models.HashingRequest, error) {
	initialSize := c.requestQueue.Size()
	var hashingRequests []models.HashingRequest
	for i := 0; i < initialSize; i++ {
		hashingRequest, err := c.requestQueue.Get()
		if err != nil {
			return nil, err
		}

		hashingRequests = append(hashingRequests, hashingRequest)
	}

	return hashingRequests, nil
}

func (c *ClientQueueFlusher) emptySubmissionQueue() ([]models.HashSubmission, error) {
	var hashSubmissions []models.HashSubmission
	for i := 0; i < c.submissionQueue.Size(); i++ {
		hashSubmission, err := c.submissionQueue.Get()
		if err != nil {
			return nil, err
		}

		hashSubmissions = append(hashSubmissions, hashSubmission)
	}

	return hashSubmissions, nil
}
