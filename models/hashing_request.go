package models

import (
	"hash"
)

type HashingRequest struct {
	Hash      hash.Hash
	HashName  string
	Passwords []string
}
