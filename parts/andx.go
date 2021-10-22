package parts

import (
	"fmt"
	"github.com/slmder/qbuilder/parts/expression"
	"strings"

)

type AndX struct {
	Expressions []expression.RawExpression
}

func (w *AndX) Reset() {
	w.Expressions = []expression.RawExpression{}
}

func (w *AndX) Add(expr string) {
	w.Expressions = append(w.Expressions, expression.RawExpression{Expression: expr})
}

func (w *AndX) String() string {
	if len(w.Expressions) > 0 {
		return joinAndXExpressions(w.Expressions)
	}
	return ""
}

func joinAndXExpressions(expressions []expression.RawExpression) string {
	res := make([]string, len(expressions))
	for i, expr := range expressions {
		res[i] = fmt.Sprintf("(%s)", strings.Trim(expr.String(), " "))
	}
	return strings.Join(res, " AND ")
}
