package payloads

import (
	"net/http"

	"github.com/PatateDu609/matcha/utils/database"
	"github.com/PatateDu609/matcha/utils/log"
	"github.com/google/uuid"
)

type User struct {
	ID         string  `json:"id"`
	FirstName  string  `json:"first-name"`
	LastName   string  `json:"last-name"`
	FullName   string  `json:"full-name"`
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	Biography  *string `json:"bio"`
	FameRating int     `json:"score"`
	Password   string  `json:"-"` // this field is ignored when exchanging data with user, it is intended for internal usage only
}

func (s User) GetName() string {
	return "users"
}

func (s User) GetAlias() string {
	return "u"
}

func (s User) GetColumns() []string {
	return database.GetColumns[User](true)
}

func (s User) GetMandatoryColumns() []string {
	return database.GetColumns[User](false)
}

func GetUserByIdentifier(w http.ResponseWriter, r *http.Request, identifier string) *User {

	cond := database.NewCondition("username", database.EqualTo, identifier).
		Or(database.NewCondition("email", database.EqualTo, identifier))

	if _, err := uuid.Parse(identifier); err == nil {
		cond = database.NewCondition("id", database.EqualTo, identifier)
	}

	arr, err := database.Select[User](r.Context(), cond)
	if err != nil {
		log.Logger.Errorf("couldn't get user: %+v", err)
		http.Error(w, "internal error: couldn't get user", http.StatusInternalServerError)

		return nil
	}

	if len(arr) == 0 {
		log.Logger.Errorf("couldn't find user for identifier `%s`", identifier)
		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	return &arr[0]
}
