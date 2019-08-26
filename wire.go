//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ZacharyGroff/CrowdCrack/client"
	"github.com/ZacharyGroff/CrowdCrack/config"
)

func InitializeClient() client.Client {
	wire.Build(client.NewClient, config.NewClientConfig)
	return client.Client{}
}
