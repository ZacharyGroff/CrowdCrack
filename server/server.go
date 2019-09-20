package server

import (
	"log"
	"github.com/ZacharyGroff/CrowdCrack/api"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/reader"
	"github.com/ZacharyGroff/CrowdCrack/verifier"
)

type Server struct {
	Config *config.ServerConfig
	Api *api.Api
	Reader reader.PasswordReader
	Verifier verifier.Verifier
}

func NewServer(c *config.ServerConfig, a *api.Api, r *reader.WordlistReader, v *verifier.HashVerifier) Server {
	return Server{c, a, r, v}
}

func (s Server) Start() {
	log.Println("Starting Server...")
	err := s.Reader.LoadPasswords()
	if err != nil {
		log.Fatal(err)
	}

	go s.Verifier.Verify()

	s.Api.HandleRequests()
}
