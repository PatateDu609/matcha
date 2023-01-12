package payloads

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/PatateDu609/matcha/config"
	"github.com/PatateDu609/matcha/utils/database"
)

type SignUp struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (s *SignUp) GetName() string {
	return "users"
}

func (s *SignUp) GetAlias() string {
	return "u"
}

func (s *SignUp) GetColumns() []string {
	return database.GetColumns[SignUp](true)
}

func (s *SignUp) GetMandatoryColumns() []string {
	return database.GetColumns[SignUp](false)
}

func (s *SignUp) PrepareInsertion() ([]string, []any) {
	keys := s.GetColumns()
	return keys, database.PrepareValues(s)
}

func (s *SignUp) Push(r *http.Request) error {
	ctx := r.Context()

	hash, err := bcrypt.GenerateFromPassword([]byte(s.Password), config.BcryptCost)
	if err != nil {
		return fmt.Errorf("couldn't hash password: %v", err)
	}
	s.Password = string(hash)

	if err := database.Insert(ctx, s); err != nil {
		return fmt.Errorf("couldn't insert user: %+v", err)
	}
	return nil
}
