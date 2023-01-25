package room

import (
	"net/http"

	"github.com/PatateDu609/matcha/routes/payloads"
	"github.com/PatateDu609/matcha/utils/log"
	"github.com/google/uuid"
)

type roomCreatePayload struct {
	User1 string `json:"user1"`
	User2 string `json:"user2"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	payload, err := payloads.Unmarshal[roomCreatePayload](r.Body)
	if err != nil {
		log.Logger.Errorf("couldn't parse user input: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user1, err := uuid.Parse(payload.User1)
	if err != nil {
		log.Logger.Errorf("couldn't parse user input: %s", err)
		http.Error(w, "user1 is bad formatted", http.StatusBadRequest)
		return
	}
	user2, err := uuid.Parse(payload.User2)
	if err != nil {
		log.Logger.Errorf("couldn't parse user input: %s", err)
		http.Error(w, "user2 is bad formatted", http.StatusBadRequest)
		return
	}

	room, err := payloads.NewRoom(r, user1, user2)
	if err != nil {
		log.Logger.Errorf("couldn't insert new room: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := payloads.Marshal(room, w); err != nil {
		log.Logger.Errorf("couldn't marshal room: %+v", err)
		http.Error(w, "internal error: couldn't send room data", http.StatusInternalServerError)
	}
}
