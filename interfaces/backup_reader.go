package interfaces

type BackupReader interface {
	BackupsExist() bool
	LoadBackups() error
}
