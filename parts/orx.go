package parts

import (
	"fmt"
	"github.com/slmder/qbuilder/parts/expression"
	"strings"

)

type OrX struct {
	Expressions []expression.RawExpression
}

func (w *OrX) Reset() {
	w.Expressions = []expression.RawExpression{}
}

func (w *OrX) Add(expr string) {
	w.Expressions = append(w.Expressions, expression.RawExpression{Expression: expr})
}

func (w *OrX) String() string {
	if len(w.Expressions) > 0 {
		return joinOrXExpressions(w.Expressions)
	}
	return ""
}

func joinOrXExpressions(expressions []expression.RawExpression) string {
	res := make([]string, len(expressions))
	for i, expr := range expressions {
		res[i] = fmt.Sprintf("(%s)", strings.Trim(expr.String(), " "))
	}
	return strings.Join(res, " OR ")
}
