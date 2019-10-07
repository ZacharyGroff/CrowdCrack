package config

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
