package mocks

type mockApi struct {
    handleRequestsCalls uint64
}

func (m *mockApi) HandleRequests() {
    m.handleRequestsCalls++
}
