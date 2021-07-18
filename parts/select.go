package parts

import (
	"fmt"
	expression2 "github.com/sergeyKhashimov/qbuilder/parts/expression"
	"strings"
)

type Select struct {
	aliases []expression2.RawColumnAlias
}

func (p *Select) String() string {
	return fmt.Sprintf("SELECT %s", joinColumns(p.aliases))
}

func (p *Select) Add(alias string) {
	expr := strings.Split(alias, ",")
	for _, e := range expr {
		p.aliases = append(p.aliases, expression2.RawColumnAlias{Alias: strings.Trim(e, " ")})
	}
}

func (p *Select) Reset() {
	p.aliases = []expression2.RawColumnAlias{}
}

func joinColumns(aliases []expression2.RawColumnAlias) string {
	var res []string
	for _, alias := range aliases {
		res = append(res, strings.Trim(alias.String(), " "))
	}
	columns := strings.Join(res, ", ")
	if columns == "" {
		return "*"
	}
	return columns
}
