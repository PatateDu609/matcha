package routes

import (
	"net/http"

	"github.com/PatateDu609/matcha/auth/session"
	"github.com/PatateDu609/matcha/routes/payloads"
	"github.com/PatateDu609/matcha/utils/log"
)

func getCurrentUser(w http.ResponseWriter, r *http.Request) {
	if !session.GlobalSessions.CheckSessionCookie(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sess := session.GlobalSessions.SessionStart(w, r)

	raw, ok := sess.Get("uuid").(string)

	if raw == "" || !ok {
		log.Logger.Errorf("couldn't extract user id from session")

		session.GlobalSessions.SessionDestroy(w, r)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user := payloads.GetUserByIdentifier(w, r, raw)

	if user == nil {
		return
	}

	if err := payloads.Marshal(user, w); err != nil {
		log.Logger.Errorf("%+v", err)
		http.Error(w, "internal error: couldn't get user", http.StatusInternalServerError)
	}
}
