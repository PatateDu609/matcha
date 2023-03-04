package oauth

import (
	"github.com/google/uuid"
)

func GetRandomState() (uuid.UUID, error) {
	return uuid.NewRandom()
}
