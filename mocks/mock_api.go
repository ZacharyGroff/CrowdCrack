package mocks

type MockApi struct {
	HandleRequestsCalls uint64
}

func (m *MockApi) HandleRequests() {
	m.HandleRequestsCalls++
}
