import payloads

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/PatateDu609/matcha/config"
	"github.com/PatateDu609/matcha/utils/database"
)

type User struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
	Images    string `json:"images"`
	Score     string `json:"score"`
	Online    string `json:"online"`
}
