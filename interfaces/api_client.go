package interfaces

import (
	"github.com/ZacharyGroff/CrowdCrack/models"
)

type ApiClient interface {
	GetHashName() (int, string)
	GetPasswords(int) (int, []string)
	SubmitHashes(models.HashSubmission) int
}
