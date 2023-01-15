package payloads

type LogIn struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
