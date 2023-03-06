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
		"username":		u.Username,
		"email":		u.Email,
		"biography":	u.Biography,
		"gender":		u.Gender,
		"orientation":	u.Orientation,
	}

	cond := database.NewCondition("id", database.EqualTo, identifier)

	_, err := database.Update[User](ctx, patch, cond)
	if err != nil {
		return err
	}
	return nil
}
