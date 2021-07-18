package qbuilder

func Select(sel... string) *selectBuilder {
	builder := selectBuilder{}
	builder.parameters = NewParameterBag()
	return builder.Select(sel...)
}

func SubSelect(sel string) *selectBuilder {
	builder := selectBuilder{}
	builder.parameters = NewParameterBag()
	return builder.SubSelect(sel)
}

func Insert(into string) *insertBuilder {
	builder := insertBuilder{}
	builder.parameters = NewParameterBag()
	return builder.Insert(into)
}

func Update(rel string) *updateBuilder {
	builder := updateBuilder{}
	builder.parameters = NewParameterBag()
	return builder.Update(rel)
}

func Delete(rel string) *deleteBuilder {
	builder := deleteBuilder{}
	builder.parameters = NewParameterBag()
	return builder.Delete(rel)
}
