package flusher

import (
	"bufio"
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/models"
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
	file, err := os.OpenFile("request_queue_overflow.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	requests, err := c.emptyRequestQueue()
	if err != nil {
		return err
	}

	for _, request := range requests {
		fmt.Fprintln(writer, request)
	}

	return writer.Flush()
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

func (c *ClientQueueFlusher) flushSubmissionQueue() error {
	file, err := os.OpenFile("submission_queue_overflow.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	submissions, err := c.emptySubmissionQueue()
	if err != nil {
		return err
	}

	for _, submission := range submissions {
		fmt.Fprintln(writer, submission)
	}

	return writer.Flush()
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
