package collector

type baseFilter struct {
	next          FilterFunction
	operationName string

	predicateFunction func(predicate, value string) bool
}

func (f *baseFilter) Next(fn FilterFunction) {
	f.next = fn
}

func (f *baseFilter) Apply(r Result, def FilterDefinition) Result {

	if f.operationName == def.GetOperation() {
		return f.filter(r, def)
	}

	if f.next == nil {
		return r
	}

	return f.next.Apply(r, def)
}

func (f *baseFilter) filter(r Result, def FilterDefinition) Result {
	predicate := def.GetValue()
	results := make(Result, 0)

	for _, v := range r {
		value, ok := v[def.Field]
		if !ok {
			continue
		}

		if f.predicateFunction(predicate, value) {
			results = append(results, v)
		}
	}
	return results
}
