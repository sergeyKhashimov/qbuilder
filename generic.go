package qbuilder

import "fmt"

type Conditions map[string]interface{}

type SortDirection int

const (
	SortDirectionASC SortDirection = iota
	SortDirectionDESC
)

func (d SortDirection) String() string {
	return [...]string{"ASC", "DESC"}[d]
}

type Sort map[string]SortDirection

func StrToDirection(str string) (SortDirection, error) {
	m := map[string]SortDirection{
		"ASC":  SortDirectionASC,
		"DESC": SortDirectionDESC,
	}

	if direction, ok := m[str]; ok {
		return direction, nil
	}
	return SortDirectionDESC, fmt.Errorf("invalid direction %s", str)
}
