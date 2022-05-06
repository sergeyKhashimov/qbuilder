package qbuilder

import (
	"fmt"
	"strings"

	"github.com/slmder/qbuilder/parts"
)

type DeleteBuilder struct {
	builder
	delete    parts.Delete
	using     parts.Using
	where     parts.Where
	returning parts.Returning
}

func (d *DeleteBuilder) Delete(rel string) *DeleteBuilder {
	d.delete = parts.Delete{Relation: rel}
	return d
}

func (d *DeleteBuilder) Using(rel string) *DeleteBuilder {
	d.using = parts.Using{Relation: rel}
	return d
}

func (d *DeleteBuilder) Where(expr ...string) *DeleteBuilder {
	d.where.Reset()
	for _, e := range expr {
		d.where.Add(e)
	}
	return d
}

func (d *DeleteBuilder) Wheref(format string, a ...interface{}) *DeleteBuilder {
	return d.Where(fmt.Sprintf(format, a...))
}

func (d *DeleteBuilder) AndWhere(expr string) *DeleteBuilder {
	d.where.Add(expr)
	return d
}

func (d *DeleteBuilder) AndWheref(format string, a ...interface{}) *DeleteBuilder {
	d.AndWhere(fmt.Sprintf(format, a...))
	return d
}

func (d *DeleteBuilder) SetParameter(name string, value interface{}) *DeleteBuilder {
	d.parameters.Set(name, value)
	return d
}

func (d *DeleteBuilder) SetParameters(params map[string]interface{}) *DeleteBuilder {
	for name, value := range params {
		d.parameters.Set(name, value)
	}
	return d
}

func (d *DeleteBuilder) RemoveParameter(name string) *DeleteBuilder {
	d.parameters.Remove(name)
	return d
}

func (d *DeleteBuilder) Returning(expr string) *DeleteBuilder {
	d.returning.Expr(expr)
	return d
}

func (d *DeleteBuilder) With(name string, sql string, cols ...string) *DeleteBuilder {
	d.with.AddDefinition(name, sql, cols...)
	return d
}

func (d *DeleteBuilder) WithRecursive(name string, sql string, cols ...string) *DeleteBuilder {
	d.with.Recursive = true
	return d.With(name, sql, cols...)
}

func (d *DeleteBuilder) ToSQL() string {
	expressions := []string{
		d.with.String(),
		d.delete.String(),
		d.using.String(),
		d.where.String(),
		d.returning.String(),
	}
	return strings.Trim(strings.Join(expressions, " "), " ")
}
