package routes

import (
	"net/http"

	"github.com/PatateDu609/matcha/routes/payloads"
	"github.com/PatateDu609/matcha/utils/log"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	payload, err := payloads.Marshal[payloads.SignUp](r.Body)
	if err != nil {
		log.Logger.Errorf("bad request: %+v", err)
		http.Error(w, "Bad formed request", http.StatusBadRequest)
		return
	}

	log.Logger.Infof("received: %+v", payload)
	if err = payload.Push(r); err != nil {
		log.Logger.Errorf("%+v", err)
		http.Error(w, "internal error: couldn't insert user", http.StatusInternalServerError)
	} else {
		log.Logger.Infof("The user has been correctly inserted")
		w.WriteHeader(http.StatusCreated)
	}
}
