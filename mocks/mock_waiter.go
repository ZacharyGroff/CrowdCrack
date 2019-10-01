package mocks

type MockWaiter struct {
	WaitCalls uint64
}

func NewMockWaiter() MockWaiter {
	return MockWaiter{0}
}

func (m *MockWaiter) Wait() {
	m.WaitCalls++
}
