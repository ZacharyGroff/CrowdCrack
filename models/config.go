package models

type Config struct {
	ApiPort                  uint16
	ComputedHashOverflowPath string
	FlushToFile              bool
	HashFunction             string
	HashlistPath             string
	HashQueueBuffer          uint64
	LogPath                  string
	LogFrequencyInSeconds    uint64
	PasswordQueueBuffer      uint64
	ServerAddress            string
	Verbose                  bool
	WordlistPath             string
}
