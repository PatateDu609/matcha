package routes

import (
	"io"
	"net/http"

	"github.com/PatateDu609/matcha/auth"
	"github.com/PatateDu609/matcha/auth/session"
	"github.com/PatateDu609/matcha/utils/log"
)

func verify(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Logger.Errorf("couldn't read user input: %s", err)
		return
	}

	if !session.GlobalSessions.CheckSessionCookie(r) {
		w.WriteHeader(http.StatusUnauthorized)
		log.Logger.Errorf("user is not authenticated")
		return
	}

	sess, err := session.GlobalSessions.SessionStart(w, r)
	if err != nil {
		log.Logger.Errorf("couldn't start session: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := auth.ConfirmEmail(r.Context(), bytes, sess); err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Logger.Errorf("couldn't confirm email: %s", err)
	}

	w.WriteHeader(http.StatusOK)
}
