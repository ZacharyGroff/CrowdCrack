package encoder

import (
	"hash"
)

type Encoder interface {
	Encode(hash.Hash) error
}
