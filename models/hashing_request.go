package models

import (
	"hash"
)

type HashingRequest struct {
	Hash hash.Hash
	Passwords []string
}
