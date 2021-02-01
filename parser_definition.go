package collector

type ParserDefinition struct {
	LookAheadManual   bool `yaml:"lookahead-manual"`
	SkipResults       bool `yaml:"skip-results"`
	SkipNumberResults int  `yaml:"skip-number-results"`
}
