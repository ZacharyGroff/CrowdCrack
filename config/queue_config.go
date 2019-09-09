package config

type QueueConfig interface {
	GetHashQueueBuffer() uint64
	GetPasswordQueueBuffer() uint64
	GetFlushToFile() bool 
	GetComputedHashOverflowPath() string
}
