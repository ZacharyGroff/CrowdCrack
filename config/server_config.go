package config

import (
	"encoding/json"
	"os"
	"log"
	"io/ioutil"
)

type ServerConfig struct {
	WordlistPath string `json:"wordlistPath"`
	HashlistPath string `json:"hashlistPath"`
	HashFunction string `json:"hashFunction"`
	ApiPort uint16 `json:"apiPort"`
	PasswordQueueBuffer uint64 `json:"passwordQueueBuffer"`
	HashQueueBuffer uint64 `json:"hashQueueBuffer"`
	FlushToFile bool `json:"flushToFile"`
	ComputedHashOverflowPath string `json:"computedHashOverflowPath"`
}

func NewServerConfig() *ServerConfig {
	config := ServerConfig{}
	err := config.parseConfig("config/server_config.json")
	
	if err != nil {
		log.Fatal("Error while parsing config")
	}
	
	return &config
}

func (c ServerConfig) GetHashQueueBuffer() uint64 {
	return c.HashQueueBuffer
}

func (c ServerConfig) GetPasswordQueueBuffer() uint64 {
	return c.PasswordQueueBuffer
}

func (c ServerConfig) GetFlushToFile() bool {
	return c.FlushToFile
}

func (c ServerConfig) GetComputedHashOverflowPath() string {
	return c.ComputedHashOverflowPath
}

func (c *ServerConfig) parseConfig(path string) error {
	file, err := os.Open(path)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &c)
	
	return err
}
