package qbuilder

import (
	"fmt"
	"reflect"
	"strings"
)

type RowLevelLockMode int

const (
	LockModeUpdate RowLevelLockMode = iota
	LockModeUpdateNowait
	LockModeShare
	LockModeShareNowait
	LockModeNoKeyUpdate
	LockModeKeyShare
)

func (m RowLevelLockMode) String() string {
	return [...]string{"UPDATE", "UPDATE NOWAIT", "SHARE", "SHARE NOWAIT", "NO KEY UPDATE", "KEY SHARE"}[m]
}

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

func FieldList(obj interface{}, alias ...string) string {
	objType := reflect.TypeOf(obj)
	switch objType.Kind() {
	case reflect.Ptr:
		val := objType.Elem()
		if val.Kind() != reflect.Struct {
			argErr := fmt.Errorf("source must be a struct or struct pointer %s given", val.Kind().String())
			panic(argErr)
		}
		objType = val
	case reflect.Struct:
	default:
		argErr := fmt.Errorf("source must be a struct or struct pointer %s given", objType.Kind().String())
		panic(argErr)
	}
	names := make([]string, 0)
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		if db, ok := field.Tag.Lookup("db"); ok {
			if db != "" {
				name := db
				if len(alias) > 0 && alias[0] != "" {
					name = fmt.Sprintf("%s.%s", alias[0], db)
				}
				names = append(names, name)
			}
		}
	}
	return strings.Join(names, ", ")
}
