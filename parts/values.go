package parts

import (
	"fmt"
	"strings"
)

type Values struct {
	values []Value
}

func (v *Values) String() string {
	return fmt.Sprintf("VALUES %s", joinValues(v.values))
}

func (v *Values) Add(value string) {
	v.values = append(v.values, Value{Value: value})

}
func joinValues(values []Value) string {
	var res []string
	for _, value := range values {
		res = append(res, fmt.Sprintf("(%s)", value.String()))
	}
	return strings.Join(res, ", ")
}
