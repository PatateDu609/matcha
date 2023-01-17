package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/PatateDu609/matcha/auth/jwt"
	"github.com/PatateDu609/matcha/auth/session"
	"github.com/PatateDu609/matcha/config"
	"github.com/PatateDu609/matcha/utils/database"
	"github.com/PatateDu609/matcha/utils/mail"
)

func SendEmailConfirmation(id string, username, email string) error {
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

	//goland:noinspection HttpUrlsUsage
	link := fmt.Sprintf("http://%s/validate/email?token=%s", config.Conf.API.FrontURL, string(tok))

	builder := strings.Builder{}

	builder.WriteString(fmt.Sprintf("Hi %s!\n\n", username))
	builder.WriteString("You subscribed successfully to Bubbler!\n")
	builder.WriteString("Please confirm your email address using the link below!:\n")
	builder.WriteString(fmt.Sprintf("%s\n\n", link))
	builder.WriteString("We hope this little app will help you bubble up your relations!\n")
	builder.WriteString("Gboucett & Rbourgea")

	return mail.Send(email, "Email confirmation", builder.String())
}

func ConfirmEmail(ctx context.Context, token []byte, sess session.Session) error {
	id, ok := sess.Get("uuid").(string)
	if !ok {
		return fmt.Errorf("couldn't get user id")
	}

	username, ok := sess.Get("username").(string)
	if !ok {
		return fmt.Errorf("couldn't get username")
	}

	email, ok := sess.Get("email").(string)
	if !ok {
		return fmt.Errorf("couldn't get email")
	}

	tok, err := jwt.ValidateToken(token)
	if err != nil {
		return fmt.Errorf("couldn't validate token: %s", err)
	}

	data := tok.PrivateClaims()
	if claim, ok := data["uuid"].(string); !ok || claim != id {
		return fmt.Errorf("invalid uuid provided")
	}
	if claim, ok := data["username"].(string); !ok || claim != username {
		return fmt.Errorf("invalid username provided")
	}
	if claim, ok := data["email"].(string); !ok || claim != email {
		return fmt.Errorf("invalid email provided")
	}

	return database.VerifyUser(ctx, id)
}
