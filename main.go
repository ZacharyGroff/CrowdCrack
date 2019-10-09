package main

import (
	"os"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
)

func main() {
	args := os.Args[1:]
	if userinput.IsClient(args) {
		client := InitializeClient()
		client.Start()
	}

	server := InitializeServer()
	server.Start()
}
