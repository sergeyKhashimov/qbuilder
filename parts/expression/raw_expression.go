package expression

type RawExpression struct {
	Expression string
}

func (p RawExpression) String() string {
	return p.Expression
}