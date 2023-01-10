package database

import (
	"fmt"
)

type OptionKind uint8

const (
	order OptionKind = iota
	paginate
)

type Option interface {
	kind() OptionKind
	fmt.Stringer
}

type Order struct {
	Expression string
	Asc        bool
}

func (ord Order) kind() OptionKind {
	return order
}

func (ord Order) String() string {
	ordStr := "ASC"
	if !ord.Asc {
		ordStr = "DESC"
	}

	return fmt.Sprintf("ORDER BY %s %s", ord.Expression, ordStr)
}

type Paginate struct {
	Count *uint64
	Start uint64
}

func (pg Paginate) kind() OptionKind {
	return paginate
}

func (pg Paginate) String() string {
	return fmt.Sprintf("LIMIT %d OFFSET %d", pg.Count, pg.Start)
}
