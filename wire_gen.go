// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
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

// Injectors from wire.go:

func InitializeClient() client.Client {
	clientConfig := config.NewClientConfig()
	hashingRequestQueue := queue.NewHashingRequestQueue()
	hashingSubmissionQueue := queue.NewHashingSubmissionQueue()
	hasher := encoder.NewHasher(clientConfig, hashingRequestQueue, hashingSubmissionQueue)
	passwordRequester := requester.NewPasswordRequester(clientConfig, hashingRequestQueue)
	hashSubmitter := submitter.NewHashSubmitter(clientConfig, hashingSubmissionQueue)
	clientClient := client.NewClient(clientConfig, hasher, passwordRequester, hashSubmitter)
	return clientClient
}

func InitializeServer() server.Server {
	serverConfig := config.NewServerConfig()
	passwordQueue := queue.NewServerPasswordQueue(serverConfig)
	hashQueue := queue.NewServerHashQueue(serverConfig)
	hashApi := api.NewHashApi(serverConfig, passwordQueue, hashQueue)
	wordlistReader := reader.NewWordlistReader(serverConfig, passwordQueue)
	hashlistReader := reader.NewHashlistReader(serverConfig)
	hashVerifier := verifier.NewHashVerifier(serverConfig, hashQueue, hashlistReader)
	serverServer := server.NewServer(hashApi, wordlistReader, hashVerifier)
	return serverServer
}
