package main

import (
	"fmt"
	"crypto/sha1"
)

func main() {
	hasher := sha1.New()
	password := []byte("Hello World")
	hasher.Write(password)
	fmt.Printf("%x\n", hasher.Sum(nil))
}
