package api

import (
	"encoding/json"
	"fmt"
	"github.com/ZacharyGroff/CrowdCrack/interfaces"
	"github.com/ZacharyGroff/CrowdCrack/models"
	"net/http"
)

type HashApi struct {
	Config    *models.Config
	Hashes    interfaces.FlushingQueue
	Passwords interfaces.Queue
	Logger    interfaces.Logger
	Tracker   interfaces.Tracker
}

func NewHashApi(p interfaces.ConfigProvider, h interfaces.FlushingQueue, q interfaces.Queue, l interfaces.Logger, t interfaces.Tracker) interfaces.Api {
	return &HashApi{
		Config:    p.GetConfig(),
		Hashes:    h,
		Logger:    l,
		Passwords: q,
		Tracker:   t,
	}
}

func (a HashApi) HandleRequests() {
	logMessage := fmt.Sprintf("Api listening to requests on port %d", a.Config.ApiPort)
	a.Logger.LogMessage(logMessage)
	http.HandleFunc("/current-hash", a.getHashName)
	http.HandleFunc("/hashes", a.retrieveHashes)
	http.HandleFunc("/passwords", a.sendPasswords)
	port := fmt.Sprintf(":%d", a.Config.ApiPort)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		a.Logger.LogMessage(err.Error())
		panic(err)
	}
}

func (a HashApi) getHashName(w http.ResponseWriter, r *http.Request) {
	if a.Config.Verbose {
		a.Logger.LogMessage("API received a request for current hash name.")
	}
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
		for a.Hashes.Put(hash) != nil {
		}
	}

	json.NewEncoder(w).Encode("Submission Successful")

	numHashesComputed := uint64(len(hashSubmission.Results))
	a.Tracker.TrackHashesComputed(numHashesComputed)
	if a.Config.Verbose {
		logMessage := fmt.Sprintf("API received a hash submission containing %d computed hashes", numHashesComputed)
		a.Logger.LogMessage(logMessage)
	}
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
	if a.Config.Verbose {
		logMessage := fmt.Sprintf("API received a request for %d passwords", numPasswords)
		a.Logger.LogMessage(logMessage)
	}
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
