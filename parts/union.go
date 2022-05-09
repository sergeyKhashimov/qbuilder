package parts

import (
	"fmt"

	"github.com/slmder/qbuilder/parts/expression"
)

type Union struct {
	Blocks []expression.RawExpression
}

func (u *Union) Add(expr string, all bool) {
	cmd := "UNION"
	if all {
		cmd = fmt.Sprintf("%s ALL", cmd)
	}
	u.Blocks = append(u.Blocks, expression.RawExpression{
		Expression: fmt.Sprintf("%s (%s)", cmd, expr),
	})
}

func (u Union) String() string {
	if len(u.Blocks) == 0 {
		return ""
	}
	return joinUnionExpressions(u.Blocks)
}

func joinUnionExpressions(blocks []expression.RawExpression) string {
	var res string
	for _, block := range blocks {
		res = fmt.Sprintf("%s %s", res, block.Expression)
	}
	return res
}
