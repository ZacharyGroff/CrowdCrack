package mocks

type mockQueue struct {
    PutCalls uint64
}

func (m mockQueue) Size() int {
    return -1
}

func (m mockQueue) Get() (string, error) {
    return "", nil
}

func (m *mockQueue) Put(password string) error {
    m.PutCalls++
    return nil
}
