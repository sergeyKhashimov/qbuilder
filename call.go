package qbuilder

import (
	"fmt"

	"github.com/slmder/qbuilder/parts"
)

type CallBuilder struct {
	builder
	proc parts.Procedure
}

func (s *CallBuilder) Call(name string) *CallBuilder {
	s.proc = parts.Procedure{Name: name}
	return s
}

func (s *CallBuilder) Callf(format string, args ...interface{}) *CallBuilder {
	s.Call(fmt.Sprintf(format, args...))
	return s
}

func (s *CallBuilder) Arg(arg string) *CallBuilder {
	s.proc.AddArg(arg)
	return s
}

func (s *CallBuilder) WithArgs(args ...string) *CallBuilder {
	s.proc.SetArgs(args...)
	return s
}

func (s *CallBuilder) ToSQL() string {
	return s.proc.String()
}
