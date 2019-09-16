package queue

type SubmissionQueue interface {
	Size() int
	Get() (uint64, error)
	Put(uint64) error
}
