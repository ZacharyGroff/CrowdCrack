package server

import (
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
)

type Server struct {
	Api      interfaces.Api
	Logger   interfaces.Logger
	Reader   interfaces.PasswordReader
	Observer interfaces.Observer
	Verifier interfaces.Verifier
}

func NewServer(a interfaces.Api, l interfaces.Logger, r interfaces.PasswordReader, o interfaces.Observer, v interfaces.Verifier) Server {
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
