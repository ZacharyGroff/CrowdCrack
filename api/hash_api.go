package api

import (
	"encoding/json"
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"github.com/ZacharyGroff/CrowdCrack/queue"
	"github.com/ZacharyGroff/CrowdCrack/tracker"
	"github.com/ZacharyGroff/CrowdCrack/userinput"
	"log"
	"net/http"
)

type HashApi struct {
	Config *models.ServerConfig
	Passwords queue.Queue
	Hashes queue.FlushingQueue
	Tracker tracker.Tracker
}

func NewHashApi(p userinput.CmdLineConfigProvider, q *queue.PasswordQueue, h *queue.HashQueue, t *tracker.StatsTracker) *HashApi {
	c := p.GetServerConfig()
	return &HashApi{
		Config:    c,
		Passwords: q,
		Hashes:    h,
		Tracker:   t,
	}
}

func (a HashApi) HandleRequests() {
	log.Printf("Api listening to requests on port %d", a.Config.ApiPort)
	http.HandleFunc("/current-hash", a.getHashName)
	http.HandleFunc("/hashes", a.retrieveHashes)
	http.HandleFunc("/passwords", a.sendPasswords)
	port := fmt.Sprintf(":%d", a.Config.ApiPort)
	log.Fatal(http.ListenAndServe(port, nil))
}

func (a HashApi) getHashName(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(a.Config.HashFunction)
}

func (a HashApi) retrieveHashes(w http.ResponseWriter, r *http.Request) {
	var hashSubmission models.HashSubmission
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&hashSubmission)
	if err != nil {
		fmt.Println(err)
	}

	for _, hash := range hashSubmission.Results {
		for a.Hashes.Put(hash) != nil {}
	}
	
	json.NewEncoder(w).Encode("Submission Successful")

	numHashesComputed := uint64(len(hashSubmission.Results))
	a.Tracker.TrackHashesComputed(numHashesComputed)
}

func (a HashApi) sendPasswords(w http.ResponseWriter, r *http.Request) {
	var numPasswords uint64
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&numPasswords)
	if err != nil {
		fmt.Println(err)
	}

	if uint64(a.Passwords.Size()) < numPasswords {
		numPasswords = uint64(a.Passwords.Size())
	}

	passwords := a.getPasswords(numPasswords)

	json.NewEncoder(w).Encode(passwords)

	a.Tracker.TrackPasswordsSent(numPasswords)
}

func (a HashApi) getPasswords(n uint64) []string {
	var passwords []string
	i := uint64(0)
	for ; i < n; i++ {
		password, _ := a.Passwords.Get()
		passwords = append(passwords, password)
	}

	return passwords
}
