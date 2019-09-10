package requester

import (
	"hash"
)

type Requester interface {
	Request() (hash.Hash, error)
}
