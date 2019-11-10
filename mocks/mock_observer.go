package mocks

type MockObserver struct {
	StartCalls uint64
	StopCalls  uint64
}

func NewMockObserver() MockObserver {
	return MockObserver{
		StartCalls: 0,
		StopCalls:  0,
	}
}

func (m *MockObserver) Start() {
	m.StartCalls++
}

func (m *MockObserver) Stop() {
	m.StopCalls++
}
