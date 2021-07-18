package parts

type Value struct {
	Value       string
}

func (v Value) String() string {
	return v.Value
}
