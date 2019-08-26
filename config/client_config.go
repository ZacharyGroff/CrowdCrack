package config

import (
	"encoding/json"
	"os"
	"log"
	"io/ioutil"
)

type ClientConfig struct {
	ServerAddress string `json:"serverAddress"`
}

func (c *ClientConfig) parseConfig(path string) error {
	file, err := os.Open(path)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &c)
	
	return err
}

func NewClientConfig() *ClientConfig {
	config := ClientConfig{}
	err := config.parseConfig("config/client_config.json")
	
	if err != nil {
		log.Fatal("Error while parsing config")
	}
	
	return &config
}
