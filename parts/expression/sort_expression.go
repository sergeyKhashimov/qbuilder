package expression

import "fmt"

type SortExpression struct {
	Expression string
	Direction  string
}

func (s SortExpression) String() string {
	return  fmt.Sprintf("%s %s", s.Expression, s.Direction)
}