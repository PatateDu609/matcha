package session

import (
	"net/http"

	"github.com/PatateDu609/matcha/auth/session"
	"github.com/PatateDu609/matcha/routes/payloads"
	"github.com/PatateDu609/matcha/utils/log"
	"github.com/google/uuid"
)

func GetUUID(w http.ResponseWriter, r *http.Request) {
	sess, err := session.GlobalSessions.SessionStart(nil, r)
	if err != nil {
		log.Logger.Errorf("couldn't start session: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	id, ok := sess.Get("uuid").(string)
	if !ok {
		log.Logger.Errorf("couldn't get uuid from session")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if _, err = uuid.Parse(id); err != nil {
		log.Logger.Errorf("couldn't parse uuid from session: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	resp := struct {
		UUID string `json:"uuid"`
	}{
		UUID: id,
	}
	if err = payloads.Marshal(resp, w); err != nil {
		log.Logger.Errorf("couldn't send uuid to client: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
