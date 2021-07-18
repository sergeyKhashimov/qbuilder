package parts

import (
	"fmt"
	"strings"
)

type Using struct {
	Relation string
}

func (u *Using) String() string {
	if strings.Trim(u.Relation, " ") == "" {
		return ""
	}
	return fmt.Sprintf("USING %s", u.Relation)
}
