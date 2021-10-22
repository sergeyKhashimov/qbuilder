package qbuilder

import (
	"fmt"
	"github.com/slmder/qbuilder/parts"
)

type AndXBuilder struct {
	parts parts.AndX
}

func AndX(x ...string) *AndXBuilder {
	b := &AndXBuilder{}
	for _, expr := range x {
		b.Add(expr)
	}
	return b
}

func (s *AndXBuilder) Add(x string) *AndXBuilder {
	s.parts.Add(x)
	return s
}

func (s *AndXBuilder) Addf(format string, a ...interface{}) *AndXBuilder {
	s.parts.Add(fmt.Sprintf(format, a...))
	return s
}

func (s *AndXBuilder) ToSQL() string {
	return s.parts.String()
}
