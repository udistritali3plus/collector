package collector

type eqFilter struct {
	next FilterFunction
}

func NewEqFilter() FilterFunction {
	return &eqFilter{}
}

func (e *eqFilter) Next(fn FilterFunction) {
	e.next = fn
}

func (e *eqFilter) Apply(r Result, def FilterDefinition) Result {

	if equalsFunctionName == def.GetOperation() {
		return e.filter(r, def)
	}

	if e.next == nil {
		return r
	}

	return e.next.Apply(r, def)
}

func (e *eqFilter) filter(r Result, def FilterDefinition) Result {

	predicate := def.GetValue()

	results := make(Result, 0)

	for _, v := range r {
		value, ok := v[def.Field]

		if !ok {
			continue
		}

		if predicate == value {
			results = append(results, v)
		}
	}
	return results
}
