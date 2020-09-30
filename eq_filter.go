package collector

type eqFilter struct {
	*baseFilter
}

func NewEqFilter() FilterFunction {

	f := &eqFilter{&baseFilter{}}
	f.operationName = equalsFunctionName
	f.predicateFunction =  f.evaluatePredicate
	return f
}

func (f *eqFilter) evaluatePredicate(predicate, value string) bool {
	return value == predicate
}
