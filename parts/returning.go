package parts

import (
	"fmt"
	expression2 "github.com/slmder/qbuilder/parts/expression"
	"strings"
)

type Returning struct {
	aliases []expression2.RawColumnAlias
}

func (r *Returning) String() string {
	if len(r.aliases) > 0 {
		return fmt.Sprintf("RETURNING %s", joinReturningAliases(r.aliases))
	}
	return ""
}

func (r *Returning) Add(alias string) {
	expr := strings.Split(alias, ",")
	for _, e := range expr {
		r.aliases = append(r.aliases, expression2.RawColumnAlias{Alias: strings.Trim(e, " ")})
	}
}

func (r *Returning) Reset() {
	r.aliases = []expression2.RawColumnAlias{}
}

func joinReturningAliases(aliases []expression2.RawColumnAlias) string {
	var res []string
	for _, alias := range aliases {
		res = append(res, alias.String())
	}
	return strings.Join(res, ", ")
}
