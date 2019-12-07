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
	PasswordRequestSize      uint64
	RequestBackupPath        string
	ServerAddress            string
	SubmissionBackupPath     string
	Threads                  uint16
	Verbose                  bool
	WordlistPath             string
}
