package qbuilder

import (
	"fmt"
	"strings"

	"github.com/slmder/qbuilder/parts"
	"github.com/slmder/qbuilder/parts/expression"
)

type SelectBuilder struct {
	builder
	sel       parts.Select
	from      parts.From
	alias     parts.Alias
	join      parts.Join
	where     parts.Where
	orWhere   parts.OrWhere
	limit     parts.Limit
	offset    parts.Offset
	having    parts.Having
	groupBy   parts.GroupBy
	orderBy   parts.OrderBy
	union     parts.Union
	forClause string
	lWrap     string
	rWrap     string
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

func (s *SelectBuilder) From(rel string, alias ...string) *SelectBuilder {
	s.from = parts.From{Relation: rel}
	if len(alias) > 0 {
		s.Alias(alias[0])
	}
	return s
}

func (s *SelectBuilder) Alias(alias string) *SelectBuilder {
	s.alias = parts.Alias{Alias: alias}
	return s
}

func (s *SelectBuilder) InnerJoin(rel string, alias string, on string) *SelectBuilder {
	s.join.Add(rel, parts.Alias{Alias: alias}, on, expression.DirectionInner)
	return s
}

func (s *SelectBuilder) LeftJoin(rel string, alias string, on string) *SelectBuilder {
	s.join.Add(rel, parts.Alias{Alias: alias}, on, expression.DirectionLeft)
	return s
}

func (s *SelectBuilder) RightJoin(rel string, alias string, on string) *SelectBuilder {
	s.join.Add(rel, parts.Alias{Alias: alias}, on, expression.DirectionRight)
	return s
}

func (s *SelectBuilder) CrossJoin(rel string, alias string) *SelectBuilder {
	s.join.Add(rel, parts.Alias{Alias: alias}, "", expression.DirectionCross)
	return s
}

func (s *SelectBuilder) Join(rel string, alias string, on string) *SelectBuilder {
	return s.InnerJoin(rel, alias, on)
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
	s.forClause = "FOR UPDATE"
	return s
}

func (s *SelectBuilder) For(mode RowLevelLockMode) *SelectBuilder {
	s.forClause = fmt.Sprintf("FOR %s", mode)
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
	s.ResetGroupBy()
	s.groupBy.Set(expr)
	return s
}

func (s *SelectBuilder) AddGroupBy(expr string) *SelectBuilder {
	s.groupBy.Set(expr)
	return s
}

func (s *SelectBuilder) ResetGroupBy() *SelectBuilder {
	s.groupBy.Reset()
	return s
}

func (s *SelectBuilder) OrderBy(sort Sort) *SelectBuilder {
	s.orderBy.Reset()
	for expr, direction := range sort {
		s.orderBy.Add(expr, direction.String())
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

func (s *SelectBuilder) Union(expr string) *SelectBuilder {
	s.union.Add(expr, false)
	return s
}

func (s *SelectBuilder) UnionAll(expr string) *SelectBuilder {
	s.union.Add(expr, true)
	return s
}

func (s *SelectBuilder) With(name string, sql string, cols ...string) *SelectBuilder {
	s.with.AddDefinition(name, sql, cols...)
	return s
}

func (s *SelectBuilder) WithRecursive(name string, sql string, cols ...string) *SelectBuilder {
	s.with.Recursive = true
	return s.With(name, sql, cols...)
}

func (s *SelectBuilder) ToSQL() string {
	expr := []string{
		s.with.String(),
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
		s.forClause,
		s.union.String(),
	}
	return fmt.Sprintf("%s%s%s", s.lWrap, strings.Trim(strings.Join(expr, " "), " "), s.rWrap)
}
