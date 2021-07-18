package parts

import (
	"fmt"
	expression2 "github.com/sergeyKhashimov/qbuilder/parts/expression"
	"strings"
)

type Columns struct {
	aliases []expression2.RawColumnAlias
}

func (c *Columns) String() string {
	if len(c.aliases) > 0 {
		return fmt.Sprintf("(%s)", joinColumnAliases(c.aliases))
	}
	return ""
}

func (c *Columns) Add(alias string) {
	expr := strings.Split(alias, ",")
	for _, e := range expr {
		c.aliases = append(c.aliases, expression2.RawColumnAlias{Alias: strings.Trim(e, " ")})
	}
}

func (c *Columns) Reset() {
	c.aliases = []expression2.RawColumnAlias{}
}

func joinColumnAliases(aliases []expression2.RawColumnAlias) string {
	var res []string
	for _, alias := range aliases {
		res = append(res, alias.String())
	}
	return strings.Join(res, ", ")
}
