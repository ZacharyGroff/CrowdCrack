package mocks

type MockEncoder struct {
	StartCalls uint64
	errorToReturn error
}

func NewMockEncoder(e error) MockEncoder {
	return MockEncoder {
		StartCalls:    0,
		errorToReturn: e,
	}
}

func (m *MockEncoder) Start() error {
	m.StartCalls++
	return m.errorToReturn
}

