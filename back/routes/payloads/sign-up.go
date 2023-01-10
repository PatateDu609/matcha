package payloads

type SignUp struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
