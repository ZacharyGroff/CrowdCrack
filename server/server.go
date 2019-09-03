package server

import (
	"log"
	"github.com/ZacharyGroff/CrowdCrack/api"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/reader"
)

type Server struct {
	Config *config.ServerConfig
	Api *api.Api
	Reader reader.PasswordReader
}

func NewServer(c *config.ServerConfig, a *api.Api, r *reader.WordlistReader) Server {
	return Server{c, a, r}
}

func (s Server) Start() {
	log.Println("Starting Server...")
	err := s.Reader.LoadPasswords()
	if err != nil {
		log.Fatal(err)
	}

	s.Api.HandleRequests()
}
