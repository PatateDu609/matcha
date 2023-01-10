package database

import (
	"context"
	"fmt"
	"strings"
)

func Insert[T Writable](ctx context.Context, data T) error {
	dataKeys, dataValues := data.PrepareInsertion()

	conn, err := GetConnFromCtx(ctx)
	if err != nil {
		return err
	}

	if dataValues == nil || dataKeys == nil {
		return fmt.Errorf("couldn't setup correctly the insertion into the database")
	}

	placeholders := make([]string, len(dataValues))
	for i := range placeholders {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	//goland:noinspection SyntaxError
	req := fmt.Sprintf(
		"INSERT INTO %q (%s) VALUES (%s)",
		data.GetName(),
		strings.Join(dataKeys, ", "),
		strings.Join(placeholders, ", "),
	)
	res, err := conn.Exec(ctx, req, dataValues...)
	if err != nil {
		return fmt.Errorf("insertion failed: %s", err)
	}
	if !res.Insert() {
		return fmt.Errorf("insertion failed: unknown error")
	}
	return nil
}
