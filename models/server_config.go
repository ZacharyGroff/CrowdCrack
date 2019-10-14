package models

type ServerConfig struct {
	WordlistPath string
	HashlistPath string
	LogPath string
	LogFrequencyInSeconds uint64
	HashFunction string
	ApiPort uint16
	PasswordQueueBuffer uint64
	HashQueueBuffer uint64
	FlushToFile bool
	ComputedHashOverflowPath string
}
