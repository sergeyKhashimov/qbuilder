package parts

import (
	"fmt"
)

type Delete struct {
	Relation string
}

func (d *Delete) String() string {
	return  fmt.Sprintf("DELETE FROM %s", d.Relation)
}