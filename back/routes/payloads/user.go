package payloads

import (
	"github.com/PatateDu609/matcha/utils/database"
)

type User struct {
	FirstName  string  `json:"first-name"`
	LastName   string  `json:"last-name"`
	FullName   string  `json:"full-name"`
	Username   string  `json:"username"`
	Email      string  `json:"email"`
	Biography  *string `json:"bio"`
	FameRating int     `json:"score"`
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
