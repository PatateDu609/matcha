package database

import (
	"fmt"
	"os"
	"strings"

	"github.com/PatateDu609/matcha/utils/log"
)

type Comparison uint8

const (
	EqualTo Comparison = iota
	NotEqualTo
	LessThan
	GreaterThan
	LessThanOrEqualTo
	GreaterThanOrEqualTo
)

func (c Comparison) String() string {
	switch c {
	case EqualTo:
		return "="
	case NotEqualTo:
		return "<>"
	case LessThan:
		return "<"
	case GreaterThan:
		return ">"
	case LessThanOrEqualTo:
		return "<="
	case GreaterThanOrEqualTo:
		return ">="
	}
	return ""
}

type logicalOperator uint8

const (
	and logicalOperator = iota
	or
)

func (op logicalOperator) String() string {
	switch op {
	case and:
		return "AND"
	case or:
		return "OR"
	}
	return ""
}

type logicalOperatorStruct struct {
	value logicalOperator
	next  *Condition
	prev  *Condition
}

// Condition represent the condition filter in a SQL query
type Condition struct {
	not        bool                   // Is the value preceded by a not operator
	leftValue  string                 // the value to be compared
	comparison Comparison             // The comparison between the expression and the value
	prev       *logicalOperatorStruct // the previous logical operator that links to the previous expression
	next       *logicalOperatorStruct // the next logical operator that links to the next expression
	value      any                    // This is the actual value to be compared (it will be replaced by a placeholder for the query string)
}

func NewCondition(expression string, comp Comparison, value any) *Condition {
	return &Condition{
		not:        false,
		leftValue:  expression,
		comparison: comp,
		value:      value,
	}
}

func (cond *Condition) Negate() *Condition {
	cond.not = !cond.not
	return cond
}

func (cond *Condition) last() *Condition {
	res := cond
	for {
		if res.next == nil {
			break
		}

		if res.next.next == nil {
			log.Logger.Fatal("a LogicalOperator should have a successor")
			os.Exit(1)
		}

		res = res.next.next
	}

	return res
}

func (cond *Condition) first() *Condition {
	res := cond
	for {
		if res.prev == nil {
			break
		}

		if res.prev.prev == nil {
			log.Logger.Fatal("a LogicalOperator should have a successor")
			os.Exit(1)
		}

		res = res.prev.prev
	}

	return res
}

func (cond *Condition) linkLogical(op logicalOperator, expression *Condition) *Condition {
	last := cond.last()
	last.next = &logicalOperatorStruct{
		value: op,
		next:  expression,
		prev:  last,
	}
	return last
}

func (cond *Condition) And(expression *Condition) *Condition {
	return cond.linkLogical(and, expression)
}

func (cond *Condition) Or(expression *Condition) *Condition {
	return cond.linkLogical(or, expression)
}

func (cond *Condition) String() string {
	res := strings.Builder{}

	current := cond.first()
	for i := 0; ; i++ {
		res.WriteString(current.leftValue)
		res.WriteString(fmt.Sprintf(" = $%d", i+1))

		if current.next == nil {
			break
		}
		res.WriteString(" ")
		res.WriteString(current.next.value.String())
		res.WriteString(" ")
		current = current.next.next
	}

	return res.String()
}

func (cond *Condition) Values() []any {
	var res []any

	current := cond.first()
	for {
		res = append(res, current.value)

		if current.next == nil {
			break
		}
		current = current.next.next
	}

	return res
}
