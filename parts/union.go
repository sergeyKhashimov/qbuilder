package parts

import (
	"fmt"

	"github.com/slmder/qbuilder/parts/expression"
)

type Union struct {
	Expression expression.RawExpression
}

func (u *Union) Set(expr string) {
	u.Expression = expression.RawExpression{
		Expression: expr,
	}
}

func (u Union) String() string {
	if u.Expression.String() == "" {
		return ""
	}
	return fmt.Sprintf("UNION (%s)", u.Expression.String())
}
