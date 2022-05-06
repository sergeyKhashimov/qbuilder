package parts

import (
	"fmt"

	"github.com/slmder/qbuilder/parts/expression"
)

type CTEDefinition struct {
	Name    string
	Columns Columns
	SQL     expression.RawExpression
}

func (d CTEDefinition) String() string {
	return fmt.Sprintf("%s %s AS (%s)", d.Name, d.Columns.String(), d.SQL.String())
}
