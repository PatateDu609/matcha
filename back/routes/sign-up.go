package routes

import (
	"net/http"

	"github.com/PatateDu609/matcha/auth"
	"github.com/PatateDu609/matcha/routes/payloads"
	"github.com/PatateDu609/matcha/utils/log"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	payload, err := payloads.Unmarshal[payloads.SignUp](r.Body)
	if err != nil {
		log.Logger.Errorf("bad request: %+v", err)
		http.Error(w, "Bad formed request", http.StatusBadRequest)
		return
	}

	log.Logger.Infof("received: %+v", payload)
	if err = payload.Push(w, r); err != nil {
		log.Logger.Errorf("%+v", err)
		http.Error(w, "internal error: couldn't insert user", http.StatusInternalServerError)
		return
	}

	log.Logger.Infof("The user has been correctly inserted with id=%s", payload.ID)

	response := payloads.SignUpResponse{
		ID: payload.ID.String(),
	}

	log.Logger.Infof("Sending response: %+v", response)

	if err := auth.ConfirmEmail(payload.ID.String(), payload.Username, payload.Email); err != nil {
		log.Logger.Errorf("couldn't send confirmation email: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := payloads.Marshal(response, w); err != nil {
		log.Logger.Errorf("couldn't send response to user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
