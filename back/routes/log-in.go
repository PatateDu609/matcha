package routes

import (
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/PatateDu609/matcha/auth"
	"github.com/PatateDu609/matcha/routes/payloads"
	"github.com/PatateDu609/matcha/utils/log"
)

func logIn(w http.ResponseWriter, r *http.Request) {
	payload, err := payloads.Unmarshal[payloads.LogIn](r.Body)
	if err != nil {
		log.Logger.Errorf("couldn't unmarshal payload: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := payloads.GetUserByIdentifier(w, r, payload.Identifier)
	if user == nil {
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			log.Logger.Infof("user provided bad credentials")
			http.Error(w, "password or username/email invalid", http.StatusUnauthorized)
			return
		}
		log.Logger.Errorf("couldn't compare password and hash: %+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"uuid":      user.ID,
		"username":  user.Username,
		"email":     user.Email,
		"full name": user.FullName,
	}

	auth.Authenticate(w, r, data)

	if err := payloads.Marshal(user, w); err != nil {
		log.Logger.Errorf("couldn't send data to user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
