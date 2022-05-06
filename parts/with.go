package parts

import (
	"fmt"
	"strings"

	"github.com/slmder/qbuilder/parts/expression"
)

type With struct {
	Recursive   bool
	Definitions []CTEDefinition
}

func (f With) String() string {
	var recursive string
	if f.Recursive {
		recursive = "RECURSIVE "
	}
	if len(f.Definitions) == 0 {
		return ""
	}
	return fmt.Sprintf("WITH %s%s", recursive, JoinCTEDefinitions(f.Definitions))
}

func (f *With) AddDefinition(name string, sql string, cols ...string) {
	var aliases = make([]expression.RawColumnAlias, len(cols))
	for i, c := range cols {
		aliases[i] = expression.RawColumnAlias{
			Alias: c,
		}
	}
	f.Definitions = append(f.Definitions, CTEDefinition{
		Name:    name,
		Columns: Columns{aliases: aliases},
		SQL:     expression.RawExpression{Expression: sql},
	})
}

func JoinCTEDefinitions(defs []CTEDefinition) string {
	var res = make([]string, len(defs))
	for i, def := range defs {
		res[i] = def.String()
	}
	return strings.Join(res, ", ")
}
