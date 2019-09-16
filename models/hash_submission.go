package models

type HashSubmission struct {
	HashType string `json:"hashType"`
	Results []string `json:"results"`
}
