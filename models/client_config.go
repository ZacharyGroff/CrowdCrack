package models

type ClientConfig struct {
	ServerAddress string
	HashQueueBuffer uint64
	PasswordQueueBuffer uint64
	FlushToFile bool
	ComputedHashOverflowPath string
}
