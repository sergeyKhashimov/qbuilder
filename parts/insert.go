package parts

import (
	"fmt"
)

type Insert struct {
	Relation string
}

func (i *Insert) String() string {
	if i.Relation != "" {
		return fmt.Sprintf("INSERT INTO %s", i.Relation)
	}
	return ""
}
