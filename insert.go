package qbuilder

import (
	parts2 "github.com/fortuwealth/backend/pkg/db/qbuilder/parts"
	"strings"
)

type insertBuilder struct {
	builder
	insert    parts2.Insert
	columns   parts2.Columns
	values    parts2.Values
	returning parts2.Returning
}

func (i *insertBuilder) Insert(into string) *insertBuilder {
	i.insert = parts2.Insert{Relation: into}
	return i
}

func (i *insertBuilder) Value(val string) *insertBuilder {
	i.values.Add(val)
	return i
}

func (i *insertBuilder) Columns(alias ...string) *insertBuilder {
	for _, a := range alias {
		i.columns.Add(a)
	}
	return i
}

func (i *insertBuilder) Returning(alias ...string) *insertBuilder {
	for _, a := range alias {
		i.returning.Add(a)
	}
	return i
}

func (i *insertBuilder) SetParameter(name string, value interface{}) *insertBuilder {
	i.parameters.Set(name, value)
	return i
}

func (i *insertBuilder) SetParameters(params map[string]interface{}) *insertBuilder {
	for name, value := range params {
		i.parameters.Set(name, value)
	}
	return i
}

func (i *insertBuilder) RemoveParameter(name string) *insertBuilder {
	i.parameters.Remove(name)
	return i
}

func (i *insertBuilder) ToSQL() string {
	expressions := []string{
		i.insert.String(),
		i.columns.String(),
		i.values.String(),
		i.returning.String(),
	}
	return strings.Trim(strings.Join(expressions, " "), " ")
}
