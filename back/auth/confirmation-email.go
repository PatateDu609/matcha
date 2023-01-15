package auth

import (
	"fmt"
	"strings"

	"github.com/PatateDu609/matcha/auth/jwt"
	"github.com/PatateDu609/matcha/config"
	"github.com/PatateDu609/matcha/utils/mail"
)

func ConfirmEmail(id string, username, email string) error {
	data := map[string]any{
		"uuid":         id,
		"email":        email,
		"username":     username,
		"confirmation": false,
	}

	tok, err := jwt.SignToken(data)
	if err != nil {
		return err
	}

	if email == "" || username == "" || id == "" {
		return fmt.Errorf("please provide an uuid, an email address and a username")
	}

	link := fmt.Sprintf("%s/validate/email/%s", config.Conf.API.FrontURL, string(tok))

	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("Hi %s!\n\n", username))
	builder.WriteString("You subscribed successfully to Bubbler!\n")
	builder.WriteString("Please confirm your email address using the link below!:\n")
	builder.WriteString(fmt.Sprintf("%s\n\n", link))
	builder.WriteString("We hope this little app will help you bubble up your relations!\n")
	builder.WriteString("Gboucett & Rbourgea")

	return mail.Send(email, "Email confirmation", builder.String())
}
