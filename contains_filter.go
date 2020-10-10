package collector

import "strings"

type containsFilter struct {
	*baseFilter
}

func NewContainsFilter() FilterFunction {
	f := &containsFilter{&baseFilter{}}
	f.operationName = containsFunctionName
	f.predicateFunction = f.evaluatePredicate
	return f
}

func (f *containsFilter) evaluatePredicate(predicate, value string) bool {
	return strings.Contains(value, predicate)
}
