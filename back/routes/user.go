package routes

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/PatateDu609/matcha/routes/payloads"
	"github.com/PatateDu609/matcha/utils/log"
	"github.com/PatateDu609/matcha/utils/database"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	userid := chi.URLParam(r, "uuid")

	if userid == "" {
		log.Logger.Errorf("bad request: empty string")
		w.WriteHeader(400)
		return
	}

	cond := database.NewCondition("id", database.EqualTo, userid)
	tab, err := database.Select[payloads.User](r.Context(), cond)

	if err == nil {
		 if err = payloads.Marshal(tab[0], w); err != nil {
			log.Logger.Errorf("%+v", err)
			http.Error(w, "internal error: couldn't get user", http.StatusInternalServerError)
		 }
	} else {
		log.Logger.Errorf("%+v", err)
		http.Error(w, "internal error: couldn't get user", http.StatusInternalServerError)
	}
}