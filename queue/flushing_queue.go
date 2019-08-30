package queue

type FlushingQueue interface {
	Size() int
	Get() (string, error)
	Put(string) error
	Flush() error
}
