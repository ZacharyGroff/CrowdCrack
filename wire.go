//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ZacharyGroff/CrowdCrack/api"
	"github.com/ZacharyGroff/CrowdCrack/apiclient"
	"github.com/ZacharyGroff/CrowdCrack/client"
	"github.com/ZacharyGroff/CrowdCrack/encoder"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/observer"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/reader"
	"github.com/ZacharyGroff/CrowdCrack/requester"
	"github.com/ZacharyGroff/CrowdCrack/server"
	"github.com/ZacharyGroff/CrowdCrack/tracker"
	"github.com/ZacharyGroff/CrowdCrack/submitter"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
	"github.com/ZacharyGroff/CrowdCrack/verifier"
	"github.com/ZacharyGroff/CrowdCrack/waiter"
)

func InitializeClient() client.Client {
	wire.Build(client.NewClient, encoder.NewHasher, requester.NewPasswordRequester, submitter.NewHashSubmitter, apiclient.NewHashApiClient, queue.NewHashingRequestQueue, queue.NewHashingSubmissionQueue, waiter.NewSleeper, logger.NewConcurrentLogger, userinput.NewCmdLineConfigProvider)
	return client.Client{}
}

func InitializeServer() server.Server {
	wire.Build(server.NewServer, api.NewHashApi, verifier.NewHashVerifier, reader.NewHashlistReader, reader.NewWordlistReader, queue.NewPasswordQueue, queue.NewHashQueue, observer.NewStatsObserver, logger.NewConcurrentLogger, tracker.NewStatsTracker, userinput.NewCmdLineConfigProvider)
	return server.Server{}
}
