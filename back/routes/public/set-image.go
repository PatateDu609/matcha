package public

import (
	"net/http"

	"github.com/PatateDu609/matcha/routes/payloads"
	"github.com/PatateDu609/matcha/utils/log"
)

func setImage(w http.ResponseWriter, r *http.Request) {
	payload, err := payloads.Unmarshal[payloads.Image](r.Body)
	if err != nil {
		log.Logger.Errorf("bad request: %+v", err)
		http.Error(w, "Bad formed request", http.StatusBadRequest)
		return
	}

	log.Logger.Infof("received: %+v", payload)
	if err = payload.Push(w, r); err != nil {
		log.Logger.Errorf("%+v", err)
		http.Error(w, "internal error: couldn't insert image", http.StatusInternalServerError)
		return
	}

	log.Logger.Infof("The image has been correctly inserted with number=%s", payload.Number)

	w.Header().Set("Content-Type", "application/json")
	if err := payloads.Marshal(payload, w); err != nil {
		log.Logger.Errorf("couldn't send response to user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
