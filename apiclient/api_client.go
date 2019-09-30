package apiclient

type ApiClient interface {
	GetHashName() (int, string)
	GetPasswords(int) (int, []string)
	SubmitHashes([]string) int
}