package server

import (
	"github.com/ZacharyGroff/CrowdCrack/api"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/logger"
	"github.com/ZacharyGroff/CrowdCrack/observer"
	"github.com/ZacharyGroff/CrowdCrack/reader"
	"github.com/ZacharyGroff/CrowdCrack/verifier"
)

type Server struct {
	Api      interfaces.Api
	Logger   interfaces.Logger
	Reader   interfaces.PasswordReader
	Observer interfaces.Observer
	Verifier interfaces.Verifier
}

func NewServer(a *api.HashApi, l *logger.ConcurrentLogger, r *reader.WordlistReader, o *observer.StatsObserver, v *verifier.HashVerifier) Server {
	return Server{
		Api:      a,
		Logger:   l,
		Reader:   r,
		Observer: o,
		Verifier: v,
	}
}

func (s Server) Start() {
	s.Logger.LogMessage("Starting Server...")
	err := s.Reader.LoadPasswords()
	if err != nil {
		panic(err)
	}

	go s.Verifier.Start()
	go s.Observer.Start()

	s.Api.HandleRequests()
}
