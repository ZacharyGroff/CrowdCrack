package requester

import (
	"hash"
)

type Requester interface {
	Request() error
}
