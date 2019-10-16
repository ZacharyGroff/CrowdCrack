package server

import (
	"github.com/ZacharyGroff/CrowdCrack/observer"
	"log"
	"github.com/ZacharyGroff/CrowdCrack/api"
	"github.com/ZacharyGroff/CrowdCrack/reader"
	"github.com/ZacharyGroff/CrowdCrack/verifier"
)

type Server struct {
	Api api.Api
	Reader reader.PasswordReader
	Verifier verifier.Verifier
	Observer observer.Observer
}

func NewServer(a *api.HashApi, r *reader.WordlistReader, v *verifier.HashVerifier, o *observer.StatsObserver) Server {
	return Server{
		Api:      a,
		Reader:   r,
		Verifier: v,
		Observer: o,
	}
}

func (s Server) Start() {
	log.Println("Starting Server...")
	err := s.Reader.LoadPasswords()
	if err != nil {
		panic(err)
	}

	go s.Verifier.Verify()
	go s.Observer.Start()

	s.Api.HandleRequests()
}
