package models

import (
	"hash"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func GetSupportedHashFunctions() map[string]hash.Hash {
	return map[string]hash.Hash {
		"md4": md4.New(),
		"md5": md5.New(),
		"sha1": sha1.New(),
		"sha256": sha256.New(),
		"sha512": sha512.New(),
		"ripemd160": ripemd160.New(),
		"sha3_224": sha3.New224(),
		"sha3_256": sha3.New256(),
		"sha3_384": sha3.New384(),
		"sha3_512": sha3.New512(),
		"sha512_224": sha512.New512_224(),
		"sha512_256": sha512.New512_256(),
	}
}
