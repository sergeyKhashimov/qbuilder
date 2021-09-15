package qbuilder

import (
	parts2 "github.com/slmder/qbuilder/parts"
	"strings"
)

type DeleteBuilder struct {
	builder
	delete parts2.Delete
	using  parts2.Using
	where  parts2.Where
}

func (d *DeleteBuilder) Delete(rel string) *DeleteBuilder {
	d.delete = parts2.Delete{Relation: rel}
	return d
}

func (d *DeleteBuilder) Using(rel string) *DeleteBuilder {
	d.using = parts2.Using{Relation: rel}
	return d
}

func (d *DeleteBuilder) Where(expr... string) *DeleteBuilder {
	d.where.Reset()
	for _, e := range expr {
		d.where.Add(e)
	}
	return d
}

func (d *DeleteBuilder) AndWhere(expr string) *DeleteBuilder {
	d.where.Add(expr)
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

func (d *DeleteBuilder) ToSQL() string {
	expressions := []string{
		d.delete.String(),
		d.using.String(),
		d.where.String(),
	}
	return strings.Trim(strings.Join(expressions, " "), " ")
}
