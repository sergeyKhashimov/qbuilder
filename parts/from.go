package parts

import (
	"fmt"
)

type From struct {
	Relation string
}

func (f From) String() string {
	if f.Relation != "" {
		return fmt.Sprintf("FROM %s", f.Relation)
	}
	return ""
}
