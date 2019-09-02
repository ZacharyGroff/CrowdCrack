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
	http.HandleFunc("/current-hash", a.getHashName)
	http.HandleFunc("/hashes", a.retrieveHashes)
	port := fmt.Sprintf(":%d", a.Config.ApiPort)
	log.Fatal(http.ListenAndServe(port, nil))
}

func (a Api) getHashName(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(a.Config.HashFunction)
}

func (a Api) retrieveHashes(w http.ResponseWriter, r *http.Request) {
	var hashSubmission HashSubmission
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&hashSubmission)
	if err != nil {
		fmt.Println(err)
	}

	for _, hash := range hashSubmission.Results {
		for a.Hashes.Put(hash) != nil {}
	}
	
	json.NewEncoder(w).Encode("Submission Successful")
}

func (a Api) sendPasswords(w http.ResponseWriter, r *http.Request) {
	var passwordRequest PasswordRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	var numPasswords uint64
	if uint64(a.Passwords.Size()) < passwordRequest.NumPasswords {
		numPasswords = uint64(a.Passwords.Size())
	} else {
		numPasswords = passwordRequest.NumPasswords
	}

	var passwords []string
	var i uint64
	for i = 0; i < numPasswords; i++ {
		password, _ := a.Passwords.Get()
		passwords = append(passwords, password)
	}

	json.NewEncoder(w).Encode(passwords)
}
