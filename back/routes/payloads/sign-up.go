package payloads

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/PatateDu609/matcha/auth/session"
	"github.com/PatateDu609/matcha/config"
	"github.com/PatateDu609/matcha/utils"
	"github.com/PatateDu609/matcha/utils/database"
	"github.com/PatateDu609/matcha/utils/log"
	"github.com/google/uuid"
)

type SignUp struct {
	ID        uuid.UUID `json:"-"`
	FirstName string    `json:"first-name"`
	LastName  string    `json:"last-name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

type SignUpResponse struct {
	ID string `json:"id"`
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

func (s *SignUp) Push(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	id, err := uuid.NewRandom()
	if err != nil {
		return fmt.Errorf("couldn't get uuid: %v", err)
	}
	s.ID = id

	hash, err := bcrypt.GenerateFromPassword([]byte(s.Password), config.BcryptCost)
	if err != nil {
		return fmt.Errorf("couldn't hash password: %v", err)
	}
	s.Password = string(hash)

	s.FirstName = utils.TitleCase(s.LastName)
	s.LastName = utils.TitleCase(s.LastName)

	if err := database.Insert(ctx, s); err != nil {
		return fmt.Errorf("couldn't insert user: %+v", err)
	}

	sess := session.GlobalSessions.SessionStart(w, r)
	data := map[string]string{
		"uuid":      s.ID.String(),
		"full name": fmt.Sprintf("%s %s", s.FirstName, s.LastName),
		"username":  s.Username,
		"email":     s.Email,
	}
	for key, val := range data {
		if err := sess.Set(key, val); err != nil {
			log.Logger.Warnf("couldn't set `%s` in session: %+v", key, err)
		}
	}
	return nil
}
