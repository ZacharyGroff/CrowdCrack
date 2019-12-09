package mocks

type MockBackupReader struct {
	BackupExistsCalls uint64
	LoadBackupsCalls  uint64
	boolToReturn      bool
	errorToReturn     error
}

func (m *MockBackupReader) BackupsExist() bool {
	m.BackupExistsCalls++
	return m.boolToReturn
}

func (m *MockBackupReader) LoadBackups() error {
	m.LoadBackupsCalls++
	return m.errorToReturn
}
