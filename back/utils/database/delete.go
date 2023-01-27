package database

import (
	"context"
	"fmt"
)

func Delete(ctx context.Context, name string, cond *Condition) error {
	conn, err := GetConnFromCtx(ctx)
	if err != nil {
		return err
	}

	req := fmt.Sprintf("DELETE FROM %s WHERE %q", name, cond.String())

	res, err := conn.Exec(ctx, req, cond.Values()...)
	if err != nil {
		return fmt.Errorf("deletion failed: %s", err)
	}
	if !res.Delete() {
		return fmt.Errorf("deletion failed: unknown error")
	}
	return nil
}
