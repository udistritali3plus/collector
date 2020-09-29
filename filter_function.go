package collector

const (
	equalsFunctionName   = "eq"
	containsFunctionName = "contains"
)

type FilterFunction interface {
	Apply(r Result, def FilterDefinition) Result

	Next(fn FilterFunction)
}
