package interfaces

type HashReader interface {
	GetHashes() (map[string]bool, error)
}
