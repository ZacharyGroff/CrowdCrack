package mocks

type MockVerifier struct {
    VerifyCalls uint64
}

func (m *MockVerifier) Verify() {
    m.VerifyCalls++
}
