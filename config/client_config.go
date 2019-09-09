package config

import (
	"encoding/json"
	"os"
	"log"
	"io/ioutil"
)

type ClientConfig struct {
	ServerAddress string `json:"serverAddress"`
	HashQueueBuffer uint64 `json:"hashQueueBuffer"`
	PasswordQueueBuffer uint64 `json:"passwordQueueBuffer"`
	FlushToFile bool `json:"flushToFile"`
	ComputedHashOverflowPath string `json:"computedHashOverflowPath"`	
}

func NewClientConfig() *ClientConfig {
	config := ClientConfig{}
	err := config.parseConfig("config/client_config.json")
	
	if err != nil {
		log.Fatal("Error while parsing config")
	}
	
	return &config
}

func (c ClientConfig) GetHashQueueBuffer() uint64 {
	return c.HashQueueBuffer
}

func (c ClientConfig) GetPasswordQueueBuffer() uint64 {
	return c.PasswordQueueBuffer
}

func (c ClientConfig) GetFlushToFile() bool {
	return c.FlushToFile
}

func (c ClientConfig) GetComputedHashOverflowPath() string {
	return c.ComputedHashOverflowPath
}

func (c *ClientConfig) parseConfig(path string) error {
	file, err := os.Open(path)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &c)
	
	return err
}
