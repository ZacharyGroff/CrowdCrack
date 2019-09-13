package queue

type SubmissionQueue interface {
	Get() (uint64, error)
	Put(uint64) error
}
