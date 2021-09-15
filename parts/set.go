package parts

import (
	"fmt"
	"github.com/slmder/qbuilder/parts/expression"
	"strings"
)

type Set struct {
	Expressions []expression.SetExpression
}

func (s *Set) String() string {
	if len(s.Expressions) > 0 {
		return fmt.Sprintf("SET %s", joinSetExpressions(s.Expressions))
	}
	return ""
}

func (s *Set) Add(column string, value string) {
	s.Expressions = append(s.Expressions, expression.SetExpression{Column: column, Value: value})
}

func joinSetExpressions(expressions []expression.SetExpression) string {
	var res []string
	for _, expr := range expressions {
		res = append(res, strings.Trim(expr.String(), " "))
	}
	return strings.Join(res, ", ")
}
