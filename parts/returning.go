package parts

import (
	"fmt"
	"github.com/slmder/qbuilder/parts/expression"
)

type Returning struct {
	expr expression.RawExpression
}

func (r *Returning) String() string {
	if r.expr.Expression != "" {
		return fmt.Sprintf("RETURNING %s", r.expr)
	}
	return ""
}

func (r *Returning) Expr(expr string) {
	r.expr.Expression = expr
}

func (r *Returning) Reset() {
	r.expr.Expression = ""
}
