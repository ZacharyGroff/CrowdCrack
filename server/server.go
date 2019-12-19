package server

import (
	"fmt"
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

	go s.tryLoadPasswords()
	go s.Verifier.Start()
	go s.Observer.Start()

	s.Api.HandleRequests()
}

func (s Server) tryLoadPasswords() {
	err := s.Reader.LoadPasswords()
	if err != nil {
		logMessage := fmt.Sprintf("Error when loading passwords: %s", err.Error())
		s.Logger.LogMessage(logMessage)
		return
	}
}
