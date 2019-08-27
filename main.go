package main

func main() {
	client := InitializeClient()
	client.Start()
	
	server := InitializeServer()
	server.Start()
}
