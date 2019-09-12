package models

import (
	"hash"
)

type HashingRequest struct {
	Hash hash.Hash
	NumPasswords uint64
}
