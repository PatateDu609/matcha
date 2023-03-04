package oauth

type GooglePayload struct {
	AuthUser string `json:"authuser"`
	Code     string `json:"code"`
	Prompt   string `json:"prompt"`
	Scope    string `json:"scope"`
	State    string `json:"state"`
}
