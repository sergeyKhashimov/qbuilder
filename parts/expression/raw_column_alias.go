package expression

type RawColumnAlias struct {
	Alias string
}

func (p RawColumnAlias) String() string {
	return p.Alias
}
