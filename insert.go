package qbuilder

import (
	"fmt"
	"github.com/slmder/qbuilder/parts"
	"strings"
)

type InsertBuilder struct {
	builder
	insert    parts.Insert
	columns   parts.Columns
	values    parts.Values
	returning parts.Returning
}

func (i *InsertBuilder) Insert(into string) *InsertBuilder {
	i.insert = parts.Insert{Relation: into}
	return i
}

func (i *InsertBuilder) Row(set map[string]string) *InsertBuilder {
	var value string
	for column, val := range set {
		i.columns.Add(column)
		value = strings.Trim(fmt.Sprintf("%s,%s", value, strings.Trim(val, ", ")), ", ")
	}
	i.values.Add(value)
	return i
}

func (i *InsertBuilder) RowE(obj interface{}, exclude ...string) *InsertBuilder {
	return i.Row(StringMap(obj, exclude...))
}

func (i *InsertBuilder) Value(val string) *InsertBuilder {
	i.values.Add(val)
	return i
}

func (i *InsertBuilder) Columns(alias ...string) *InsertBuilder {
	for _, a := range alias {
		i.columns.Add(a)
	}
	return i
}

func (i *InsertBuilder) Returning(expr string) *InsertBuilder {
	i.returning.Expr(expr)
	return i
}

func (i *InsertBuilder) SetParameter(name string, value interface{}) *InsertBuilder {
	i.parameters.Set(name, value)
	return i
}

func (i *InsertBuilder) SetParameters(params map[string]interface{}) *InsertBuilder {
	for name, value := range params {
		i.parameters.Set(name, value)
	}
	return i
}

func (i *InsertBuilder) RemoveParameter(name string) *InsertBuilder {
	i.parameters.Remove(name)
	return i
}

func (i *InsertBuilder) ToSQL() string {
	expressions := []string{
		i.insert.String(),
		i.columns.String(),
		i.values.String(),
		i.returning.String(),
	}
	return strings.Trim(strings.Join(expressions, " "), " ")
}
