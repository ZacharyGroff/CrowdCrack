package config

type ClientConfig struct {
	ServerAddress string `json:"serverAddress"`
	HashQueueBuffer uint64 `json:"hashQueueBuffer"`
	PasswordQueueBuffer uint64 `json:"passwordQueueBuffer"`
	FlushToFile bool `json:"flushToFile"`
	ComputedHashOverflowPath string `json:"computedHashOverflowPath"`	
}
