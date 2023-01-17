package database

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"time"

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

	config.MaxConns = int32(runtime.NumCPU() * 5)
	config.ConnConfig.ConnectTimeout = time.Second * 2

	config.ConnConfig.Tracer = Tracer{}

	config.BeforeAcquire = func(ctx context.Context, conn *pgx.Conn) bool {
		log.Logger.Traceln("Acquiring connection...")
		return true
	}

	var err error
	pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Logger.Fatalf("couldn't connect to database: %+v", err)
	}
}

func GetConnFromCtx(ctx context.Context) (conn *pgxpool.Conn, err error) {
	err = nil
	conn, ok := ctx.Value(ctxKey).(*pgxpool.Conn)
	if ok {
		return
	}

	conn = nil
	err = fmt.Errorf("no database connection found")
	return
}

func AcquireMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		conn, err := pool.Acquire(context.Background())
		if err != nil {
			log.Logger.Errorf("couldn't acquire database connection: %s", err)
			return
		}
		ctx = context.WithValue(ctx, ctxKey, conn)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

		conn.Release()
	}

	return http.HandlerFunc(fn)
}
