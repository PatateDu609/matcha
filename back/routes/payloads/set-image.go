package payloads

import (
	"fmt"
	"net/http"

	"github.com/PatateDu609/matcha/utils/database"
)

func (i *Image) PrepareInsertion() ([]string, []any) {
	keys := i.GetColumns()
	return keys, database.PrepareValues(i)
}

func (i *Image) Push(r *http.Request) error {
	ctx := r.Context()

	if err := database.Insert(ctx, i); err != nil {
		return fmt.Errorf("couldn't insert image: %+v", err)
	}
	return nil
}
