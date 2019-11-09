package interfaces

type Queue interface {
	Size() int
	Get() (string, error)
	Put(string) error
}
