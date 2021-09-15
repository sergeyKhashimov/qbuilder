package parts

import (
	expression2 "github.com/slmder/qbuilder/parts/expression"
	"strings"
)

type Join struct {
	Expressions []expression2.JoinExpression
}

func (j Join) String() string {
	if len(j.Expressions) > 0 {
		return joinJoinExpressions(j.Expressions)
	}
	return ""
}

func (j *Join) Add(rel string, alias Alias, cond string, direct expression2.Direction) {
	j.Expressions = append(j.Expressions, expression2.JoinExpression{Relation: rel, Alias: alias.String(), Condition: cond, Direction: direct})
}

func joinJoinExpressions(expressions []expression2.JoinExpression) string {
	var res []string
	for _, expr := range expressions {
		res = append(res, strings.Trim(expr.String(), " "))
	}
	return strings.Join(res, "\n")
}
