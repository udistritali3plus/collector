package collector

type Filter interface {
	Apply(r Result, def *definition) Result
}

type filter struct{}

func NewFilter() Filter {
	return &filter{}
}

func (f *filter) Apply(r Result, def *definition) Result {

	filter := NewFilterBuilder().Build()

	for _, f := range def.Filters {
		r = filter.Apply(r, f)
	}
	return r
}
