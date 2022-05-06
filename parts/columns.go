package parts

import (
	"fmt"
	"strings"

	"github.com/slmder/qbuilder/parts/expression"
)

type Columns struct {
	aliases []expression.RawColumnAlias
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
		c.aliases = append(c.aliases, expression.RawColumnAlias{Alias: strings.Trim(e, " ")})
	}
}

func (c *Columns) Reset() {
	c.aliases = []expression.RawColumnAlias{}
}

func joinColumnAliases(aliases []expression.RawColumnAlias) string {
	var res = make([]string, len(aliases))
	for _, alias := range aliases {
		res = append(res, alias.String())
	}
	return strings.Join(res, ", ")
}
