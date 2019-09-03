package reader

type PasswordReader interface {
	LoadPasswords() error
}
