package database

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/PatateDu609/matcha/utils/log"
)

func Update[T Relation](ctx context.Context, patch Patch, cond *Condition) ([]T, error) {
	var instanceOf T
	var res []T

	conn, err := GetConnFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	valueOf := reflect.ValueOf(instanceOf)
	if valueOf.Kind() != reflect.Struct {
		return nil, fmt.Errorf("can't use non struct types (type is %s)", valueOf.Kind().String())
	}

	patchString, patchValues := patch.prepareUpdate()
	tableName := instanceOf.GetName()
	condString, condValues := cond.stringify(len(patch)), cond.Values()
	columns := strings.Join(instanceOf.GetColumns(), ", ")

	req := fmt.Sprintf("UPDATE %s SET %s WHERE %s RETURNING %s", tableName, patchString, condString, columns)

	rows, err := conn.Query(ctx, req, []any{patchValues, condValues}...)
	if err != nil {
		return nil, err
	}

	if !rows.CommandTag().Update() {
		return nil, fmt.Errorf("no row has been updated")
	}
	log.Logger.Tracef("updated %d", rows.CommandTag().RowsAffected())

	reflectedType := valueOf.Type()
	for rows.Next() {
		row := reflect.New(reflectedType)
		arr := GetInterfaceArray(row.Interface().(*T))
		log.Logger.Trace(fmt.Sprintf("arr: %q", arr))

		if err = rows.Scan(arr...); err != nil {
			log.Logger.Error(fmt.Sprintf("couldn't scan into interface: %s", err))
			return nil, fmt.Errorf("couldn't scan values: %s", err)
		}

		log.Logger.Trace(fmt.Sprintf("current row: %q", row))
		res = append(res, row.Elem().Interface().(T))
	}

	log.Logger.Debugf("updated %d rows", len(res))
	return res, nil
}
