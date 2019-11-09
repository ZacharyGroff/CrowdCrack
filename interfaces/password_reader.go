package interfaces

type PasswordReader interface {
	LoadPasswords() error
}
