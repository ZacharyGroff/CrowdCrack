package config

import (
	"encoding/json"
	"os"
	"log"
	"io/ioutil"
)

type ServerConfig struct {
	WordlistPath string `json:"wordlistPath"`
	HashFunction string `json:"hashFunction"`
	ApiPort uint16 `json:"apiPort"`
}

func NewServerConfig() *ServerConfig {
	config := ServerConfig{}
	err := config.parseConfig("config/server_config.json")
	
	if err != nil {
		log.Fatal("Error while parsing config")
	}
	
	return &config
}

func (c *ServerConfig) parseConfig(path string) error {
	file, err := os.Open(path)
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &c)
	
	return err
}
