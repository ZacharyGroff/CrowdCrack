package models

import (
	"hash"
)

type HashingRequest struct {
	Hash      hash.Hash `json:"hash"`
	HashName  string    `json:"hashName"`
	Passwords []string  `json:"passwords"`
}
