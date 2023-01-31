package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const dbTag = "db"

type Relation interface {
	GetName() string
	GetAlias() string

	// GetColumns returns all columns (mandatory or not)
	GetColumns() []string

	// GetMandatoryColumns returns all mandatory columns
	GetMandatoryColumns() []string
}

type Patch map[string]any

type Writable interface {
	Relation

	// PrepareInsertion split struct into two arrays of keys and values
	PrepareInsertion() ([]string, []any)
}

type Readable interface{}

type Base struct {
	ID        uuid.UUID  // entry's identifier
	CreatedAt time.Time  // entry's creation date
	UpdatedAt *time.Time // entry's last update date
	DeletedAt *time.Time // entry's deletion date
}

func (patch Patch) prepareUpdate() (string, []any) {
	l := len(patch)
	columns := make([]string, 0, l)
	values := make([]any, 0, l)

	i := 1

	for column, value := range patch {
		columns = append(columns, fmt.Sprintf("%s = $%d", column, i))
		values = append(values, value)

		i++
	}

	return strings.Join(columns, ", "), values
}
