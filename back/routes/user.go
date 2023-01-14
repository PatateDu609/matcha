package routes

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/PatateDu609/matcha/routes/payloads"
	"github.com/PatateDu609/matcha/utils/database"
	"github.com/PatateDu609/matcha/utils/log"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "uuid")

	if userid == "" {
		log.Logger.Errorf("bad request: empty string")
		w.WriteHeader(400)
		return
	}

	log.Logger.Infof("looking for uuid: %s", userid)

	cond := database.NewCondition("id", database.EqualTo, userid)
	tab, err := database.Select[payloads.User](r.Context(), cond)

	if err != nil {
		log.Logger.Errorf("%+v", err)
		http.Error(w, "internal error: couldn't get user", http.StatusInternalServerError)

		return
	}

	if len(tab) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err = payloads.Marshal(tab[0], w); err != nil {
		log.Logger.Errorf("%+v", err)
		http.Error(w, "internal error: couldn't get user", http.StatusInternalServerError)
	}
}
