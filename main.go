package main

import (
	"github.com/ZacharyGroff/CrowdCrack/userinput"
)

func main() {
	if userinput.IsClient() {
		client := InitializeClient()
		client.Start()
	}

	server := InitializeServer()
	server.Start()
}
