package public

import (
	"net/http"

	"github.com/PatateDu609/matcha/auth/session"
	"github.com/PatateDu609/matcha/routes/payloads"
	"github.com/PatateDu609/matcha/utils/log"
)

func setCurrentUser(w http.ResponseWriter, r *http.Request) {
	payload, err := payloads.Unmarshal[payloads.User](r.Body)
	if !session.GlobalSessions.CheckSessionCookie(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sess, err := session.GlobalSessions.SessionStart(w, r)
	if err != nil {
		log.Logger.Errorf("couldn't initialize session: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	raw, ok := sess.Get("uuid").(string)

	if raw == "" || !ok {
		log.Logger.Errorf("couldn't extract user id from session")

		if err := session.GlobalSessions.SessionDestroy(w, r); err != nil {
			log.Logger.Errorf("couldn't destroy session: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user := payloads.SetUserByIdentifier(r.Context(), raw, payload)

	if user == nil {
		return
	}

	if err := payloads.Marshal(user, w); err != nil {
		log.Logger.Errorf("%+v", err)
		http.Error(w, "internal error: couldn't get user", http.StatusInternalServerError)
	}
}
