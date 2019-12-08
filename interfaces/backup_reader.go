package interfaces

type BackupReaderInterface interface {
	BackupsExist() bool
	LoadBackups() error
}
