package qbuilder

import (
	"fmt"
	"github.com/slmder/qbuilder/parts"
	"github.com/slmder/qbuilder/parts/expression"
	"strings"
)

type SelectBuilder struct {
	builder
	sel      parts.Select
	from     parts.From
	alias    parts.Alias
	join     parts.Join
	where    parts.Where
	orWhere  parts.OrWhere
	limit    parts.Limit
	offset   parts.Offset
	having   parts.Having
	groupBy  parts.GroupBy
	orderBy  parts.OrderBy
	withLock string
	lWrap    string
	rWrap    string
}

func (s *SelectBuilder) Select(sel ...string) *SelectBuilder {
	s.sel.Reset()
	for _, expr := range sel {
		s.sel.Add(expr)
	}
	return s
}

func (s *SelectBuilder) Selectf(format string, args ...interface{}) *SelectBuilder {
	s.Select(fmt.Sprintf(format, args...))
	return s
}

func (s *SelectBuilder) AddSelect(sel string) *SelectBuilder {
	s.sel.Add(sel)
	return s
}

func (s *SelectBuilder) AddSelectf(format string, args ...interface{}) *SelectBuilder {
	s.AddSelect(fmt.Sprintf(format, args...))
	return s
}

func (s *SelectBuilder) SubSelect(sel string) *SelectBuilder {
	s.Select(sel)
	s.rWrap, s.lWrap = "(", ")"
	return s
}

func (s *SelectBuilder) From(rel string) *SelectBuilder {
	s.from = parts.From{Relation: rel}
	return s
}

func (s *SelectBuilder) Alias(alias string) *SelectBuilder {
	s.alias = parts.Alias{Alias: alias}
	return s
}

func (s *SelectBuilder) InnerJoin(rel string, alias string, cond string) *SelectBuilder {
	s.join.Add(rel, parts.Alias{Alias: alias}, cond, expression.DirectionInner)
	return s
}

func (s *SelectBuilder) LeftJoin(rel string, alias string, cond string) *SelectBuilder {
	s.join.Add(rel, parts.Alias{Alias: alias}, cond, expression.DirectionLeft)
	return s
}

func (s *SelectBuilder) RightJoin(rel string, alias string, cond string) *SelectBuilder {
	s.join.Add(rel, parts.Alias{Alias: alias}, cond, expression.DirectionRight)
	return s
}

func (s *SelectBuilder) Join(rel string, alias string, cond string) *SelectBuilder {
	return s.InnerJoin(rel, alias, cond)
}

func (s *SelectBuilder) Where(expr ...string) *SelectBuilder {
	s.where.Reset()
	for _, e := range expr {
		s.where.Add(e)
	}
	return s
}

func (s *SelectBuilder) Wheref(format string, a ...interface{}) *SelectBuilder {
	return s.Where(fmt.Sprintf(format, a...))
}

func (s *SelectBuilder) GetWhere() *parts.Where {
	return &s.where
}

func (s *SelectBuilder) AndWhere(expr string) *SelectBuilder {
	s.where.Add(expr)
	return s
}

func (s *SelectBuilder) AndWheref(format string, a ...interface{}) *SelectBuilder {
	s.AndWhere(fmt.Sprintf(format, a...))
	return s
}

func (s *SelectBuilder) OrWhere(expr string) *SelectBuilder {
	s.orWhere.Add(expr)
	return s
}

func (s *SelectBuilder) WithLock() *SelectBuilder {
	s.withLock = "FOR UPDATE"
	return s
}

func (s *SelectBuilder) Offset(offset uint32) *SelectBuilder {
	s.offset = parts.Offset{Offset: offset}
	return s
}

func (s *SelectBuilder) Limit(limit uint32) *SelectBuilder {
	s.limit = parts.Limit{Limit: limit}
	return s
}

func (s *SelectBuilder) Having(expr string) *SelectBuilder {
	s.having = parts.Having{Having: expr}
	return s
}

func (s *SelectBuilder) GroupBy(expr string) *SelectBuilder {
	s.groupBy.Set(expr)
	return s
}

func (s *SelectBuilder) AddGroupBy(expr string) *SelectBuilder {
	s.groupBy.Set(expr)
	return s
}

func (s *SelectBuilder) OrderBy(sort Sort) *SelectBuilder {
	if sort != nil {
		for expr, direction := range sort {
			s.orderBy.Add(expr, direction.String())
		}
	}
	return s
}

func (s *SelectBuilder) AddOrderBy(expr string, direction SortDirection) *SelectBuilder {
	s.orderBy.Add(expr, direction.String())
	return s
}

func (s *SelectBuilder) SetParameter(name string, value interface{}) *SelectBuilder {
	s.parameters.Set(name, value)
	return s
}

func (s *SelectBuilder) SetParameters(params map[string]interface{}) *SelectBuilder {
	for name, value := range params {
		s.parameters.Set(name, value)
	}
	return s
}

func (s *SelectBuilder) RemoveParameter(name string) *SelectBuilder {
	s.parameters.Remove(name)
	return s
}

func (s *SelectBuilder) ToSQL() string {
	expr := []string{
		s.sel.String(),
		s.from.String(),
		s.alias.String(),
		s.join.String(),
		s.where.String(),
		s.orWhere.String(),
		s.groupBy.String(),
		s.having.String(),
		s.orderBy.String(),
		s.limit.String(),
		s.offset.String(),
		s.withLock,
	}
	return fmt.Sprintf("%s%s%s", s.lWrap, strings.Trim(strings.Join(expr, " "), " "), s.rWrap)
}
