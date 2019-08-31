//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ZacharyGroff/CrowdCrack/api"
	"github.com/ZacharyGroff/CrowdCrack/client"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/server"
)

func InitializeClient() client.Client {
	wire.Build(client.NewClient, config.NewClientConfig)
	return client.Client{}
}

func InitializeServer() server.Server {
	wire.Build(server.NewServer, api.NewApi, queue.NewPasswordQueue, queue.NewHashQueue, config.NewServerConfig)
	return server.Server{}
}
