package server

import (
	"log"
	"github.com/ZacharyGroff/CrowdCrack/api"
	"github.com/ZacharyGroff/CrowdCrack/reader"
	"github.com/ZacharyGroff/CrowdCrack/verifier"
)

type Server struct {
	Api api.Api
	Reader reader.PasswordReader
	Verifier verifier.Verifier
}

func NewServer(a *api.HashApi, r *reader.WordlistReader, v *verifier.HashVerifier) Server {
	return Server{a, r, v}
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
