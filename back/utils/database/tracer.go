package database

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/PatateDu609/matcha/utils/log"
	"github.com/alecthomas/chroma/quick"
	"github.com/jackc/pgx/v5"
)

type Tracer struct{}

var insertRegex = regexp.MustCompile(`(?i)^insert into "([a-z\-_0-9]+)" \((.*)\) values \(.*\)$`)

func stringifyValue(value any) string {
	ref := reflect.TypeOf(value)
	if ref.Kind() == reflect.Pointer || ref.Kind() == reflect.Ptr {
		val := reflect.ValueOf(value)
		if val.IsNil() {
			return fmt.Sprintf("'nil'")
		} else {
			return fmt.Sprintf("'%+v'", val.Elem().Interface())
		}
	}
	return fmt.Sprintf("'%+v'", value)
}

func explainSQLFallback(sql string, values []any) string {
	for i, value := range values {
		src := fmt.Sprintf("$%d", i+1)
		dest := stringifyValue(value)

		sql = strings.Replace(sql, src, dest, 1)
	}
	return sql
}

func explainInsertSQL(sql string, values []any) string {
	matches := insertRegex.FindStringSubmatch(sql)
	if matches == nil {
		return explainSQLFallback(sql, values)
	}

	columnsGroup := matches[2]
	tableName := matches[1]
	columns := strings.Split(columnsGroup, ",")
	for i := range columns {
		columns[i] = fmt.Sprintf("%s = %s", columns[i], stringifyValue(values[i]))
	}
	return fmt.Sprintf("INSERT INTO \"%s\" (%s)", tableName, strings.Join(columns, ","))
}

func explainSQL(sql string, values []any) string {
	if strings.Contains(strings.ToLower(sql), "insert") {
		return explainInsertSQL(sql, values)
	}
	return explainSQLFallback(sql, values)
}

func (t Tracer) TraceQueryStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	builder := &strings.Builder{}
	err := quick.Highlight(builder, explainSQL(data.SQL, data.Args), "sql", "terminal", "monokai")
	if err != nil {
		log.Logger.Warnf("couldn't highlight `%s`: %s", data.SQL, err)
		builder.Reset()
		builder.WriteString(data.SQL)
	}
	request := builder.String()
	log.Logger.Trace(fmt.Sprintf("Executing query `%s`", request))

	ctx = context.WithValue(ctx, "sql", request)

	return ctx
}

func (t Tracer) TraceQueryEnd(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryEndData) {
	sql, ok := ctx.Value("sql").(string)
	if data.Err != nil {
		log.Logger.Errorf("an error occured: %+v", data.Err)
	}
	if !ok {
		log.Logger.Tracef("database answered: %s", data.CommandTag)
	} else {
		log.Logger.Tracef("request: %s, database answered: %s", sql, data.CommandTag)
	}
}
