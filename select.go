package qbuilder

import (
	"fmt"
	"github.com/fortuwealth/backend/pkg/db/qbuilder/parts"
	"github.com/fortuwealth/backend/pkg/db/qbuilder/parts/expression"
	"strings"
)

type selectBuilder struct {
	builder
	sel      parts.Select
	from     parts.From
	alias    parts.Alias
	join     parts.Join
	where    parts.Where
	limit    parts.Limit
	offset   parts.Offset
	having   parts.Having
	groupBy  parts.GroupBy
	orderBy  parts.OrderBy
	withLock string
	lWrap    string
	rWrap    string
}

func (s *selectBuilder) Select(sel ...string) *selectBuilder {
	for _, expr := range sel {
		s.sel.Add(expr)
	}
	return s
}

func (s *selectBuilder) SubSelect(sel string) *selectBuilder {
	s.Select(sel)
	s.rWrap, s.lWrap = "(", ")"
	return s
}

func (s *selectBuilder) From(rel string) *selectBuilder {
	s.from = parts.From{Relation: rel}
	return s
}

func (s *selectBuilder) Alias(alias string) *selectBuilder {
	s.alias = parts.Alias{Alias: alias}
	return s
}

func (s *selectBuilder) InnerJoin(rel string, alias string, cond string) *selectBuilder {
	s.join.Add(rel, parts.Alias{Alias: alias}, cond, expression.DirectionInner)
	return s
}

func (s *selectBuilder) LeftJoin(rel string, alias string, cond string) *selectBuilder {
	s.join.Add(rel, parts.Alias{Alias: alias}, cond, expression.DirectionLeft)
	return s
}

func (s *selectBuilder) RightJoin(rel string, alias string, cond string) *selectBuilder {
	s.join.Add(rel, parts.Alias{Alias: alias}, cond, expression.DirectionRight)
	return s
}

func (s *selectBuilder) Where(expr ...string) *selectBuilder {
	s.where.Reset()
	for _, e := range expr {
		s.where.Add(e)
	}
	return s
}

func (s *selectBuilder) AndWhere(expr string) *selectBuilder {
	s.where.Add(expr)
	return s
}

func (s *selectBuilder) WithLock() *selectBuilder {
	s.withLock = "FOR UPDATE"
	return s
}

func (s *selectBuilder) Offset(offset uint32) *selectBuilder {
	s.offset = parts.Offset{Offset: offset}
	return s
}

func (s *selectBuilder) Limit(limit uint32) *selectBuilder {
	s.limit = parts.Limit{Limit: limit}
	return s
}

func (s *selectBuilder) Having(expr string) *selectBuilder {
	s.having = parts.Having{Having: expr}
	return s
}

func (s *selectBuilder) GroupBy(expr string) *selectBuilder {
	s.groupBy.Set(expr)
	return s
}

func (s *selectBuilder) AddGroupBy(expr string) *selectBuilder {
	s.groupBy.Set(expr)
	return s
}

func (s *selectBuilder) OrderBy(sort Sort) *selectBuilder {
	if sort != nil {
		for expr, direction := range sort {
			s.orderBy.Add(expr, direction.String())
		}
	}
	return s
}

func (s *selectBuilder) AddOrderBy(expr string, direction SortDirection) *selectBuilder {
	s.orderBy.Add(expr, direction.String())
	return s
}

func (s *selectBuilder) SetParameter(name string, value interface{}) *selectBuilder {
	s.parameters.Set(name, value)
	return s
}

func (s *selectBuilder) SetParameters(params map[string]interface{}) *selectBuilder {
	for name, value := range params {
		s.parameters.Set(name, value)
	}
	return s
}

func (s *selectBuilder) RemoveParameter(name string) *selectBuilder {
	s.parameters.Remove(name)
	return s
}

func (s *selectBuilder) ToSQL() string {
	expr := []string{
		s.sel.String(),
		s.from.String(),
		s.alias.String(),
		s.join.String(),
		s.where.String(),
		s.groupBy.String(),
		s.having.String(),
		s.orderBy.String(),
		s.limit.String(),
		s.offset.String(),
		s.withLock,
	}
	return fmt.Sprintf("%s%s%s", s.lWrap, strings.Trim(strings.Join(expr, " "), " "), s.rWrap)
}
