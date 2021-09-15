package parts

import (
	"fmt"
	expression2 "github.com/slmder/qbuilder/parts/expression"
	"strings"
)

type Where struct {
	Expressions []expression2.RawExpression
}

func (w *Where) Reset() {
	w.Expressions = []expression2.RawExpression{}
}

func (w *Where) Add(expr string) {
	w.Expressions = append(w.Expressions, expression2.RawExpression{Expression: expr})
}

func (w *Where) String() string {
	if len(w.Expressions) > 0 {
		return fmt.Sprintf("WHERE %s", joinWhereExpressions(w.Expressions))
	}
	return ""
}

func joinWhereExpressions(expressions []expression2.RawExpression) string {
	var res []string
	for _, expr := range expressions {
		res = append(res, fmt.Sprintf("(%s)", strings.Trim(expr.String(), " ")))
	}
	return strings.Join(res, " AND ")
}
