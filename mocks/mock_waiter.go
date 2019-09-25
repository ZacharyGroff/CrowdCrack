package mocks

type MockWaiter struct {
	WaitCalls uint64
}

func (m *MockWaiter) Wait() {
	m.WaitCalls++
}
