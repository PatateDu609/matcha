package payloads

import (
	"context"

	"github.com/PatateDu609/matcha/utils/database"
	"github.com/PatateDu609/matcha/utils/log"
)

func SetUserByIdentifier(ctx context.Context, identifier string, u User) error {

	log.Logger.Infof("in payload user = ", u)

	patch := database.Patch {
		"first_name":	u.FirstName,
		"last_name":	u.LastName,
		"full_name":	u.FullName,
		"username":		u.Username,
		"email":		u.Email,
		"biography":	u.Biography,
	}

	cond := database.NewCondition("id", database.EqualTo, identifier)

	_, err := database.Update[User](ctx, patch, cond)
	if err != nil {
		return err
	}
	return nil
}
