package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrNoConnection = errors.New("no connection object found in context")
)

func VerifyUser(ctx context.Context, userID string) error {
	conn, ok := ctx.Value(ctxKey).(*pgxpool.Conn)
	if !ok {
		return ErrNoConnection
	}

	_, err := conn.Exec(ctx, "call verify_user($1::uuid);", userID)

	return err
}
