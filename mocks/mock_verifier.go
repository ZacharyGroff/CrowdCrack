package mocks

type MockVerifier struct {
    VerifyCalls uint64
}

func (m *MockVerifier) Start() {
    m.VerifyCalls++
}
