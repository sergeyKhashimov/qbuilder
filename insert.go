package qbuilder

import (
	"fmt"
	"strings"

	"github.com/slmder/qbuilder/parts"
)

type InsertBuilder struct {
	builder
	insert     parts.Insert
	columns    parts.Columns
	values     parts.Values
	onConflict *OnConflict
	returning  parts.Returning
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

func (i *InsertBuilder) OnConflict() *OnConflict {
	if i.onConflict == nil {
		i.onConflict = &OnConflict{
			Insert: i,
		}
	}
	return i.onConflict
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

func (i *InsertBuilder) With(name string, sql string, cols ...string) *InsertBuilder {
	i.with.AddDefinition(name, sql, cols...)
	return i
}

func (i *InsertBuilder) WithRecursive(name string, sql string, cols ...string) *InsertBuilder {
	i.with.Recursive = true
	return i.With(name, sql, cols...)
}

func (i *InsertBuilder) ToSQL() string {
	expressions := []string{
		i.with.String(),
		i.insert.String(),
		i.columns.String(),
		i.values.String(),
		i.onConflict.String(),
		i.returning.String(),
	}
	return strings.Trim(strings.Join(expressions, " "), " ")
}
