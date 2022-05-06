package expression

import (
	"strings"
)

type RawExpression struct {
	Expression string
}

func (p RawExpression) String() string {
	return strings.Trim(p.Expression, " ")
}
