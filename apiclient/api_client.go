package apiclient

type ApiClient interface {
	GetHashName() (int, string)
	SubmitHashes([]string) int
	GetPasswords(uint64) (int, []string)
}