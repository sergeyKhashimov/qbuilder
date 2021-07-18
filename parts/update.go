package parts

import (
	"fmt"
)

type Update struct {
	Relation string
}

func (u *Update) String() string {
	return  fmt.Sprintf("UPDATE %s", u.Relation)
}