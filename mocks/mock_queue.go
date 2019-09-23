package mocks

type MockQueue struct {
    PutCalls uint64
}

func (m MockQueue) Size() int {
    return -1
}

func (m MockQueue) Get() (string, error) {
    return "", nil
}

func (m *MockQueue) Put(password string) error {
    m.PutCalls++
    return nil
}
