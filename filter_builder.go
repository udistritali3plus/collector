package collector

type filterBuilder struct {
}

func NewFilterBuilder() *filterBuilder {
	return &filterBuilder{}
}

func (b *filterBuilder) Build() FilterFunction {
	eq := NewEqFilter()
	return eq
}
