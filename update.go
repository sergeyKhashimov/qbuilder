package qbuilder

import (
	parts2 "github.com/sergeyKhashimov/qbuilder/parts"
	"strings"
)

type updateBuilder struct {
	builder
	update parts2.Update
	set    parts2.Set
	from   parts2.From
	where  parts2.Where
}

func (u *updateBuilder) Update(rel string) *updateBuilder {
	u.update = parts2.Update{Relation: rel}
	return u
}

func (u *updateBuilder) From(rel string) *updateBuilder {
	u.from = parts2.From{Relation: rel}
	return u
}

func (u *updateBuilder) Set(column string, value string) *updateBuilder {
	u.set.Add(column, value)
	return u
}

func (u *updateBuilder) Where(expr... string) *updateBuilder {
	u.where.Reset()
	for _, e := range expr {
		u.where.Add(e)
	}
	return u
}

func (u *updateBuilder) AndWhere(expr string) *updateBuilder {
	u.where.Add(expr)
	return u
}

func (u *updateBuilder) SetParameter(name string, value interface{}) *updateBuilder {
	u.parameters.Set(name, value)
	return u
}

func (u *updateBuilder) SetParameters(params map[string]interface{}) *updateBuilder {
	for name, value := range params {
		u.parameters.Set(name, value)
	}
	return u
}

func (u *updateBuilder) RemoveParameter(name string) *updateBuilder {
	u.parameters.Remove(name)
	return u
}

func (u *updateBuilder) ToSQL() string {
	expressions := []string{
		u.update.String(),
		u.set.String(),
		u.from.String(),
		u.where.String(),
	}
	return strings.Trim(strings.Join(expressions, " "), " ")
}