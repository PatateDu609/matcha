package database

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/PatateDu609/matcha/utils/log"
)

func Select[T Relation](ctx context.Context, cond *Condition, options ...Option) ([]T, error) {
	var instanceOf T
	var res []T

	// The context should already have a Conn object
	conn, err := GetConnFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	valueOf := reflect.ValueOf(instanceOf)
	if valueOf.Kind() != reflect.Struct {
		return nil, fmt.Errorf("can't use non struct types (type is %s)", valueOf.Kind().String())
	}

	queryBuilder := strings.Builder{}

	columnsArr := instanceOf.GetColumns()
	columns := strings.Join(columnsArr, ", ")
	//goland:noinspection SyntaxError
	queryBuilder.WriteString(fmt.Sprintf("SELECT %s FROM %q", columns, instanceOf.GetName()))
	var values []any
	if cond != nil {
		values = cond.Values()
		queryBuilder.WriteString(fmt.Sprintf(" WHERE %s", cond.String()))
	}

	for _, opt := range options {
		queryBuilder.WriteString(" ")
		queryBuilder.WriteString(opt.String())
	}

	query := queryBuilder.String()
	req, err := conn.Query(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer req.Close()

	reflectedType := valueOf.Type()
	for req.Next() {
		row := reflect.New(reflectedType)
		arr := GetInterfaceArray(row.Interface().(*T))
		// logger.Trace(fmt.Sprintf("arr: %q", arr))

		if err = req.Scan(arr...); err != nil {
			log.Logger.Error(fmt.Sprintf("couldn't scan into interface: %s", err))
			return nil, fmt.Errorf("couldn't scan values: %s", err)
		}

		log.Logger.Trace(fmt.Sprintf("current row: %q", row))
		res = append(res, row.Elem().Interface().(T))
	}

	log.Logger.Debug(fmt.Sprintf("got %d rows", len(res)))
	return res, nil
}
