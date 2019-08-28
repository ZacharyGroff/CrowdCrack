package api

import (
	"log"
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/ZacharyGroff/CrowdCrack/config"
)

type Api struct {
	Config *config.ServerConfig
}

func NewApi(c *config.ServerConfig) *Api {
	return &Api{c}
}

func (a Api) HandleRequests() {
	log.Printf("Api listening to requests on port %d", a.Config.ApiPort)
	http.HandleFunc("/hash", a.getHash)
	port := fmt.Sprintf(":%d", a.Config.ApiPort)
	log.Fatal(http.ListenAndServe(port, nil))
}

func (a Api) getHash(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(a.Config.HashFunction)
}
