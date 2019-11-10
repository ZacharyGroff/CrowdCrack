package mocks

type MockHashReader struct {
	GetHashesCalls uint64
	mapToReturn    map[string]bool
	errorToReturn  error
}

func NewMockHashReader(m map[string]bool, e error) MockHashReader {
	return MockHashReader{0, m, e}
}

func (m *MockHashReader) GetHashes() (map[string]bool, error) {
	m.GetHashesCalls++
	return m.mapToReturn, m.errorToReturn
}
