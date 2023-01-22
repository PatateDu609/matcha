package auth

import (
	"net/http"

	"github.com/PatateDu609/matcha/auth/session"
	"github.com/PatateDu609/matcha/utils/log"
)

func Authenticate(w http.ResponseWriter, r *http.Request, data map[string]string) {
	sess, err := session.GlobalSessions.SessionStart(w, r)
	if err != nil {
		log.Logger.Errorf("couldn't start session: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for key, val := range data {
		if err := sess.Set(key, val); err != nil {
			log.Logger.Warnf("couldn't set `%s` in session: %+v", key, err)
		}
	}
}
