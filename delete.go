package qbuilder

import (
	parts2 "github.com/fortuwealth/backend/pkg/db/qbuilder/parts"
	"strings"
)

type deleteBuilder struct {
	builder
	delete parts2.Delete
	using  parts2.Using
	where  parts2.Where
}

func (d *deleteBuilder) Delete(rel string) *deleteBuilder {
	d.delete = parts2.Delete{Relation: rel}
	return d
}

func (d *deleteBuilder) Using(rel string) *deleteBuilder {
	d.using = parts2.Using{Relation: rel}
	return d
}

func (d *deleteBuilder) Where(expr... string) *deleteBuilder {
	d.where.Reset()
	for _, e := range expr {
		d.where.Add(e)
	}
	return d
}

func (d *deleteBuilder) AndWhere(expr string) *deleteBuilder {
	d.where.Add(expr)
	return d

}
func (d *deleteBuilder) SetParameter(name string, value interface{}) *deleteBuilder {
	d.parameters.Set(name, value)
	return d
}

func (d *deleteBuilder) SetParameters(params map[string]interface{}) *deleteBuilder {
	for name, value := range params {
		d.parameters.Set(name, value)
	}
	return d
}

func (d *deleteBuilder) RemoveParameter(name string) *deleteBuilder {
	d.parameters.Remove(name)
	return d
}

func (d *deleteBuilder) ToSQL() string {
	expressions := []string{
		d.delete.String(),
		d.using.String(),
		d.where.String(),
	}
	return strings.Trim(strings.Join(expressions, " "), " ")
}