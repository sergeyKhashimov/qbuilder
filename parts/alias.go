package parts

import (
	"fmt"
)

type Alias struct {
	Alias string
}

func (p *Alias) String() string {
	if p.Alias != "" {
		return fmt.Sprintf("AS %s", p.Alias)
	}
	return ""
}