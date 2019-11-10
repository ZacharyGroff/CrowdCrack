package mocks

type MockTracker struct {
	TrackPasswordsSentCalls    uint64
	TrackHashesComputedCalls   uint64
	TrackHashesCrackedCalls    uint64
	TrackHashMatchAttemptCalls uint64
	GetPasswordsSentCalls      uint64
	GetHashesComputedCalls     uint64
	GetHashesCrackedCalls      uint64
	GetHashMatchAttemptsCalls  uint64
	uint64ToReturn             uint64
}

func NewMockTracker(u uint64) MockTracker {
	return MockTracker{
		TrackPasswordsSentCalls:    0,
		TrackHashesComputedCalls:   0,
		TrackHashesCrackedCalls:    0,
		TrackHashMatchAttemptCalls: 0,
		GetPasswordsSentCalls:      0,
		GetHashesComputedCalls:     0,
		GetHashesCrackedCalls:      0,
		GetHashMatchAttemptsCalls:  0,
		uint64ToReturn:             u,
	}
}

func (m *MockTracker) TrackPasswordsSent(uint64) {
	m.TrackPasswordsSentCalls++
}

func (m *MockTracker) TrackHashesComputed(uint64) {
	m.TrackHashesComputedCalls++
}

func (m *MockTracker) TrackHashesCracked(uint64) {
	m.TrackHashesCrackedCalls++
}

func (m *MockTracker) TrackHashMatchAttempt() {
	m.TrackHashMatchAttemptCalls++
}

func (m *MockTracker) GetPasswordsSent() uint64 {
	m.GetPasswordsSentCalls++
	return m.uint64ToReturn
}

func (m *MockTracker) GetHashesComputed() uint64 {
	m.GetHashesComputedCalls++
	return m.uint64ToReturn
}

func (m *MockTracker) GetHashesCracked() uint64 {
	m.GetHashesCrackedCalls++
	return m.uint64ToReturn
}

func (m *MockTracker) GetHashMatchAttempts() uint64 {
	m.GetHashMatchAttemptsCalls++
	return m.uint64ToReturn
}
