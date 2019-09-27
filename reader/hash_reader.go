package reader

type HashReader interface {
	GetHashes() (map[string]bool, error)
}
