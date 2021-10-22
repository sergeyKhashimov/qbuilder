package parts

import (
	"fmt"
	expression2 "github.com/slmder/qbuilder/parts/expression"
	"strings"
)

type OrWhere struct {
	Expressions []expression2.RawExpression
}

func (w *OrWhere) Reset() {
	w.Expressions = []expression2.RawExpression{}
}

func (w *OrWhere) Add(expr string) {
	w.Expressions = append(w.Expressions, expression2.RawExpression{Expression: expr})
}

func (w *OrWhere) String() string {
	expr := ""
	for _, e := range w.Expressions {
		expr = fmt.Sprintf("%s OR (%s)", expr, strings.Trim(e.String(), " "))
	}
	return expr
}
