package parts

import (
	"fmt"
	"strings"
)

type Having struct {
	Having string
}

func (p Having) String() string {
	if strings.Trim(p.Having, " ") == "" {
		return ""
	}

	return fmt.Sprintf("HAVING %s", p.Having)
}
