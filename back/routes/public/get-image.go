package public

import (
	"net/http"

	"github.com/PatateDu609/matcha/routes/payloads"
	"github.com/go-chi/chi"

	"github.com/PatateDu609/matcha/utils/log"
)

func getImage(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "uuid")

	if userid == "" {
		log.Logger.Errorf("bad request: empty string")
		w.WriteHeader(400)
		return
	}

	log.Logger.Infof("looking images for uuid: %s", userid)

	user := payloads.GetImageByOwner(w, r, userid)
	if user == nil {
		return
	}

	if err := payloads.Marshal(user, w); err != nil {
		log.Logger.Errorf("%+v", err)
		http.Error(w, "internal error: couldn't get image user", http.StatusInternalServerError)
	}
}