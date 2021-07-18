package expression

import (
	"fmt"
)

type JoinExpression struct {
	Relation  string
	Alias     string
	Condition string
	Direction Direction
}

func (j JoinExpression) String() string {
	return fmt.Sprintf("%s JOIN %s %s ON %s", j.Direction, j.Relation, j.Alias, j.Condition)
}

type Direction int

const (
	DirectionLeft Direction = iota
	DirectionRight
	DirectionInner
)

func (d Direction) String() string {
	return [...]string{"LEFT", "RIGHT", "INNER"}[d]
}
