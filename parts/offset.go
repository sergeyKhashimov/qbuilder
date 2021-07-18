package parts

import (
	"fmt"
)

type Offset struct {
	Offset uint32
}

func (o Offset) String() string {
	if o.Offset < 1 {
		return ""
	}

	return fmt.Sprintf("OFFSET %d", o.Offset)
}
