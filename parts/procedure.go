package parts

import (
	"fmt"
	"strings"

	"github.com/slmder/qbuilder/parts/expression"
)

type Procedure struct {
	Name string
	Args []expression.RawExpression
}

func (p *Procedure) String() string {
	return fmt.Sprintf("CALL %s(%s);", p.Name, joinProcArgsExpressions(p.Args))
}

func (p *Procedure) SetArgs(args ...string) {
	var res []expression.RawExpression
	for _, a := range args {
		p.AddArg(a)
	}
	p.Args = res
}

func (p *Procedure) AddArg(arg string) {
	p.Args = append(p.Args, expression.RawExpression{Expression: arg})
}

func joinProcArgsExpressions(expressions []expression.RawExpression) string {
	var res []string
	for _, expr := range expressions {
		res = append(res, expr.String())
	}
	return fmt.Sprintf("%s", strings.Join(res, ", "))
}
