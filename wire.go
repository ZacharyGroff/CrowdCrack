//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ZacharyGroff/CrowdCrack/api"
	"github.com/ZacharyGroff/CrowdCrack/client"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/encoder"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/reader"
	"github.com/ZacharyGroff/CrowdCrack/requester"
	"github.com/ZacharyGroff/CrowdCrack/server"
	"github.com/ZacharyGroff/CrowdCrack/submitter"
	"github.com/ZacharyGroff/CrowdCrack/verifier"
)

func InitializeClient() client.Client {
	wire.Build(client.NewClient, encoder.NewHasher, requester.NewPasswordRequester, submitter.NewHashSubmitter, queue.NewClientPasswordQueue, queue.NewClientHashQueue, config.NewClientConfig)
	return client.Client{}
}

func InitializeServer() server.Server {
	wire.Build(server.NewServer, api.NewApi, verifier.NewVerifier, reader.NewHashlistReader, reader.NewWordlistReader, queue.NewServerPasswordQueue, queue.NewServerHashQueue, config.NewServerConfig)
	return server.Server{}
}
