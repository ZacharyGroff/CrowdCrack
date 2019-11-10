package mocks

type MockFlushingQueue struct {
	SizeCalls      uint64
	GetCalls       uint64
	PutCalls       uint64
	FlushCalls     uint64
	stringToReturn string
	errorToReturn  error
}

func NewMockFlushingQueue(s string, e error) MockFlushingQueue {
	return MockFlushingQueue{0, 0, 0, 0, s, e}
}

func (m *MockFlushingQueue) Size() int {
	m.SizeCalls++
	return int(m.PutCalls) - int(m.GetCalls)
}

func (m *MockFlushingQueue) Get() (string, error) {
	m.GetCalls++
	return m.stringToReturn, m.errorToReturn
}

func (m *MockFlushingQueue) Put(s string) error {
	m.PutCalls++
	return m.errorToReturn
}

func (m *MockFlushingQueue) Flush() error {
	m.FlushCalls++
	return m.errorToReturn
}
