package interfaces

type FlushingQueue interface {
	Size() int
	Get() (string, error)
	Put(string) error
	Flush() error
}
