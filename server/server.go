package server

import (
	"log"
	"github.com/ZacharyGroff/CrowdCrack/api"
	"github.com/ZacharyGroff/CrowdCrack/config"
)

type Server struct {
	Config *config.ServerConfig
	Api *api.Api
}

func NewServer(c *config.ServerConfig, a *api.Api) Server {
	return Server{c, a}
}

func (s Server) Start() {
	log.Println("Starting Server...")
	s.Api.HandleRequests()
}
