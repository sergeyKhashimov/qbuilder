package qbuilder

func Select(sel... string) *SelectBuilder {
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
	return Select(SelectList(obj, alias...))
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
