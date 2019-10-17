package server

import (
	"github.com/ZacharyGroff/CrowdCrack/api"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/observer"
	"github.com/ZacharyGroff/CrowdCrack/reader"
	"github.com/ZacharyGroff/CrowdCrack/verifier"
)

type Server struct {
	Api api.Api
	Logger logger.Logger
	Reader reader.PasswordReader
	Verifier verifier.Verifier
	Observer observer.Observer
}

func NewServer(a *api.HashApi, l *logger.ServerLogger, r *reader.WordlistReader, v *verifier.HashVerifier, o *observer.StatsObserver) Server {
	return Server{
		Api:      a,
		Logger:   l,
		Reader:   r,
		Verifier: v,
		Observer: o,
	}
}

func (s Server) Start() {
	s.Logger.LogMessage("Starting Server...")
	err := s.Reader.LoadPasswords()
	if err != nil {
		panic(err)
	}

	go s.Verifier.Verify()
	go s.Observer.Start()

	s.Api.HandleRequests()
}
