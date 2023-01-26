package payloads

import (
	"net/http"

	"github.com/PatateDu609/matcha/utils/database"
	"github.com/PatateDu609/matcha/utils/log"
	"github.com/google/uuid"
)

func SetUserByIdentifier(w http.ResponseWriter, r *http.Request, identifier string) *User {
	cond := database.NewCondition("username", database.EqualTo, identifier).
		Or(database.NewCondition("email", database.EqualTo, identifier))

	if _, err := uuid.Parse(identifier); err == nil {
		cond = database.NewCondition("id", database.EqualTo, identifier)
	}

	arr, err := database.Select[User](r.Context(), cond)
	if err != nil {
		log.Logger.Errorf("couldn't Set user: %+v", err)

		if w != nil {
			http.Error(w, "internal error: couldn't Set user", http.StatusInternalServerError)
		}

		return nil
	}

	if len(arr) == 0 {
		log.Logger.Errorf("couldn't find user for identifier `%s`", identifier)
		if w != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		return nil
	}

	return &arr[0]
}
