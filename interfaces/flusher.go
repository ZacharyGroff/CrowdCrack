package interfaces

type Flusher interface {
	NeedsFlushed() bool
	Flush() error
}
