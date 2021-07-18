package parts

import (
	"fmt"
)

type Limit struct {
	Limit uint32
}

func (l Limit) String() string {
	if l.Limit < 1 {
		return ""
	}

	return fmt.Sprintf("LIMIT %d", l.Limit)
}