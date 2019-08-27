package server

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/ZacharyGroff/CrowdCrack/config"
)

type Server struct {
	Config *config.ServerConfig
}

func NewServer(c *config.ServerConfig) Server {
	return Server{c}
}

func (s Server) Start() {
	log.Println("Starting Server...")
	s.handleRequests()
}

func (s Server) handleRequests() {
	http.HandleFunc("/hash", s.getHash)
	log.Fatal(http.ListenAndServe(":2725", nil))
}

func (s Server) getHash(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.Config.HashFunction)
}
