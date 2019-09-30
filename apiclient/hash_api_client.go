package apiclient

import (
	"bytes"
	"strconv"
	"strings"
	"encoding/json"
	"net/http"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type HashApiClient struct {
	config *config.ClientConfig
}

func NewHashApiClient(c *config.ClientConfig) *HashApiClient {
	return &HashApiClient{c}
}

func (h HashApiClient) GetHashName() (int, string) {
	url := h.config.ServerAddress + "/current-hash"
	var hashName string

	response, err := http.Get(url)
	if err != nil {
		return 500, hashName
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&hashName)
	if err != nil {
		panic(err)
	}

	return response.StatusCode, hashName
}

func (h HashApiClient) GetPasswords(numPasswords int) (int, []string) {
	url := h.config.ServerAddress + "/passwords"
	requestBody := strings.NewReader(strconv.Itoa(numPasswords))
	var passwords []string

	response, err := http.Post(url, "text/plain", requestBody)
	if err != nil {
		return 500, passwords
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&passwords)
	if err != nil {
		panic(err)
	}

	return response.StatusCode, passwords
}

func (h HashApiClient) SubmitHashes(hashSubmission models.HashSubmission) int {
	url := h.config.ServerAddress + "/hashes"

	jsonHashSubmission, err := json.Marshal(hashSubmission)
	if err != nil {
		panic(err)
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonHashSubmission))
	if err != nil {
		return 500
	}

	return response.StatusCode
}
