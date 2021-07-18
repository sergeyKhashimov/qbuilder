package parts

import (
	"fmt"
	expression2 "github.com/fortuwealth/backend/pkg/db/qbuilder/parts/expression"
	"strings"
)

type GroupBy struct {
	GroupBy []expression2.RawExpression
}

func (p GroupBy) String() string {
	if len(p.GroupBy) > 0 {
		return joinGroupByExpressions(p.GroupBy)
	}
	return ""
}

func (p *GroupBy) Set(expr string) {
	p.GroupBy = []expression2.RawExpression{{Expression: expr}}
}

func (p *GroupBy) Add(expr string) {
	p.GroupBy = append(p.GroupBy, expression2.RawExpression{Expression: expr})
}

func joinGroupByExpressions(expressions []expression2.RawExpression) string {
	var res []string
	for _, expr := range expressions {
		res = append(res, expr.String())
	}
	return fmt.Sprintf("GROUP BY %s", strings.Join(res, ", "))
}
