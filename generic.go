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

func FieldList(obj interface{}, formatter func(string) string) []string {
	objType := reflect.TypeOf(obj)
	switch objType.Kind() {
	case reflect.Ptr, reflect.Slice:
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
	TaggedNames(objType, &names, formatter)
	return names
}

func SelectList(obj interface{}, alias ...string) string {
	names := FieldList(obj, func(raw string) string {
		if len(alias) > 0 && alias[0] != "" {
			return fmt.Sprintf("%s.%s", alias[0], raw)
		}
		return raw
	})
	return strings.Join(names, ", ")
}

func TaggedNames(t reflect.Type, names *[]string, formatter func(string) string) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag, ok := field.Tag.Lookup("db"); ok {
			if tag != "" {
				name := tag
				if formatter != nil {
					name = formatter(name)
				}
				*names = append(*names, name)
			}
		} else if field.Type.Kind() == reflect.Struct {
			TaggedNames(field.Type, names, formatter)
		}
	}
}
