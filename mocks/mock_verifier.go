package mocks

type mockVerifier struct {
    verifyCalls uint64
}

func (m *mockVerifier) Verify() {
    m.verifyCalls++
}
