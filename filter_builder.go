package collector

type filterBuilder struct {
}

func NewFilterBuilder() *filterBuilder {
	return &filterBuilder{}
}

func (b *filterBuilder) Build() FilterFunction {
	eq := NewEqFilter()

	contains := NewContainsFilter()
	eq.Next(contains)

	return eq
}
