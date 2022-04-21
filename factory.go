package qbuilder

func Select(sel ...string) *SelectBuilder {
	builder := SelectBuilder{}
	builder.parameters = NewParameterBag()
	return builder.Select(sel...)
}

func SubSelect(sel string) *SelectBuilder {
	builder := SelectBuilder{}
	builder.parameters = NewParameterBag()
	return builder.SubSelect(sel)
}

func SelectE(obj interface{}, alias ...string) *SelectBuilder {
	if len(alias) > 0 && alias[0] != "" {
		return Select(SelectList(obj, alias[0])).Alias(alias[0])
	}
	return Select(SelectList(obj))
}

func Insert(into string) *InsertBuilder {
	builder := InsertBuilder{}
	builder.parameters = NewParameterBag()
	return builder.Insert(into)
}

func Update(rel string) *UpdateBuilder {
	builder := UpdateBuilder{}
	builder.parameters = NewParameterBag()
	return builder.Update(rel)
}

func Delete(rel string) *DeleteBuilder {
	builder := DeleteBuilder{}
	builder.parameters = NewParameterBag()
	return builder.Delete(rel)
}

func Call(proc string) *CallBuilder {
	builder := CallBuilder{}
	builder.parameters = NewParameterBag()
	return builder.Call(proc)
}

func Callf(proc string, a ...interface{}) *CallBuilder {
	builder := CallBuilder{}
	builder.parameters = NewParameterBag()
	return builder.Callf(proc, a...)
}
