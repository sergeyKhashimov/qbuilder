package qbuilder

import (
	"fmt"
	"github.com/slmder/qbuilder/parts"
)

type OrXBuilder struct {
	parts parts.OrX
}

func OrX(x ...string) *OrXBuilder {
	b := &OrXBuilder{}
	for _, expr := range x {
		b.Add(expr)
	}
	return b
}

func (s *OrXBuilder) Add(x string) *OrXBuilder {
	s.parts.Add(x)
	return s
}

func (s *OrXBuilder) Addf(format string, a ...interface{}) *OrXBuilder {
	s.parts.Add(fmt.Sprintf(format, a...))
	return s
}

func (s *OrXBuilder) ToSQL() string {
	return s.parts.String()
}
