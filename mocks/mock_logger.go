package mocks

type MockLogger struct {
	LogMessageCalls uint64
	errorToReturn   error
}

func NewMockLogger(e error) MockLogger {
	return MockLogger{
		LogMessageCalls: 0,
		errorToReturn:   e,
	}
}

func (m *MockLogger) LogMessage(string) error {
	m.LogMessageCalls++
	return m.errorToReturn
}
