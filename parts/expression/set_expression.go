package expression

import "fmt"

type SetExpression struct {
	Column string
	Value  string
}

func (s SetExpression) String() string {
	return  fmt.Sprintf("%s = %s", s.Column, s.Value)
}