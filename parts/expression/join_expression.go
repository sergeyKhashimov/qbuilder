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
	join := fmt.Sprintf("%s JOIN %s %s", j.Direction, j.Relation, j.Alias)
	if j.Condition != "" {
		join = fmt.Sprintf("%s ON %s", join, j.Condition)
	}
	return join
}

type Direction int

const (
	DirectionLeft Direction = iota
	DirectionRight
	DirectionInner
	DirectionCross
)

func (d Direction) String() string {
	return [...]string{"LEFT", "RIGHT", "INNER", "CROSS"}[d]
}
