package encoder

import (
	"hash"
)

type Encoder interface {
	Encode() error
}
