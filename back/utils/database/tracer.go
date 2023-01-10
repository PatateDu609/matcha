package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/PatateDu609/matcha/utils/log"
	"github.com/alecthomas/chroma/quick"
	"github.com/jackc/pgx/v5"
)

type Tracer struct{}

func (t Tracer) TraceQueryStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	builder := &strings.Builder{}
	err := quick.Highlight(builder, data.SQL, "sql", "terminal", "monokai")
	if err != nil {
		log.Logger.Warnf("couldn't highlight `%s`: %s", data.SQL, err)
		builder.Reset()
		builder.WriteString(data.SQL)
	}

	log.Logger.Trace(fmt.Sprintf("Executing query `%s` with args %v", builder.String(), data.Args))

	ctx = context.WithValue(ctx, "sql", builder.String())

	return ctx
}

func (t Tracer) TraceQueryEnd(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryEndData) {
	sql, ok := ctx.Value("sql").(string)
	if !ok {
		log.Logger.Tracef("database answered: %s", data.CommandTag)
	} else {
		log.Logger.Tracef("request: %s, database answered: %s", sql, data.CommandTag)
	}
}
