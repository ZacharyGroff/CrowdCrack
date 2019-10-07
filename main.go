package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	if isClient(args) {
		client := InitializeClient()
		client.Start()
	}

	server := InitializeServer()
	server.Start()
}

func isClient(args []string) bool {
	for _, b := range args {
		if b == "--client" {
			return true
		}
	}
	return false
}
