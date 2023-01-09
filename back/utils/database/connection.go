package database

import (
	"context"

	"github.com/PatateDu609/matcha/utils/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

var pool *pgxpool.Pool

const ctxKey = "database_conn"

func SetupPool(config *pgxpool.Config) {
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	var err error
	pool, err = pgxpool.NewWithConfig(context.Background(), config)

	if err != nil {
		log.Logger.Fatalf("couldn't connect to database: %+v", err)
	}
}

func Acquire(ctx context.Context) (context.Context, error) {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return context.WithValue(ctx, ctxKey, nil), err
	}

	return context.WithValue(ctx, ctxKey, conn), nil
}
