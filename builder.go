package qbuilder

import (
	"fmt"

	"github.com/slmder/qbuilder/parts"
)

type BuilderInterface interface {
	ToSQL() string
}

type builder struct {
	parameters ParameterBag
	with       parts.With
}

func (b *builder) Parameters() ParameterBag {
	return b.parameters
}

func ToArgsAndExpressions(conditions map[string]interface{}) ([]interface{}, []string) {
	var args []interface{}
	var expressions []string

	for field, value := range conditions {
		if value == nil {
			expressions = append(expressions, fmt.Sprintf("%s IS NULL", field))
		} else {
			args = append(args, value)
			expressions = append(expressions, fmt.Sprintf("%s = $%d", field, len(args)))
		}
	}
	return args, expressions
}
