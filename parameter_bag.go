package qbuilder

type ParameterBag struct {
	parameters map[string]interface{}
}

func NewParameterBag() ParameterBag {
	return ParameterBag{parameters: make(map[string]interface{})}
}

func (p *ParameterBag) Get(name string) interface{} {
	return p.parameters[name]
}

func (p *ParameterBag) Has(name string) bool {
	if _, ok := p.parameters[name]; !ok{
		return false
	}
	return true
}

func (p *ParameterBag) Set(name string, value interface{}) {
	if p.parameters == nil {
		p.parameters = make(map[string]interface{})
	}
	p.parameters[name] = value
}

func (p *ParameterBag) Remove(name string) {
	if p.parameters == nil {
		p.parameters = make(map[string]interface{})
	}
	delete(p.parameters, name)
}

func (p *ParameterBag) All() map[string]interface{} {
	if p.parameters == nil {
		p.parameters = make(map[string]interface{})
	}
	return p.parameters
}
