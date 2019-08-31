package api

import (
	"log"
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/ZacharyGroff/CrowdCrack/config"
	"github.com/ZacharyGroff/CrowdCrack/queue"
)

type Api struct {
	Config *config.ServerConfig
	Passwords queue.Queue
	Hashes queue.FlushingQueue
}

func NewApi(c *config.ServerConfig, p *queue.PasswordQueue, h *queue.HashQueue) *Api {
	return &Api{c, p, h}
}

func (a Api) HandleRequests() {
	log.Printf("Api listening to requests on port %d", a.Config.ApiPort)
	http.HandleFunc("/current-hash", a.getHash)
	port := fmt.Sprintf(":%d", a.Config.ApiPort)
	log.Fatal(http.ListenAndServe(port, nil))
}

func (a Api) getHash(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(a.Config.HashFunction)
}
