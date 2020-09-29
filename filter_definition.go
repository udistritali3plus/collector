package collector

import (
	"regexp"
	"strings"
)

type FilterDefinition struct {
	Fn    string `yaml:"fn"`
	Field string `yaml:field`
}

func (d FilterDefinition) GetValue() string {
	re := regexp.MustCompile(parenthesesRegularExp)
	value := re.FindString(d.Fn)
	value = strings.Trim(value, "(")
	value = strings.Trim(value, ")")
	return value
}

func (d FilterDefinition) GetOperation() string {
	re := regexp.MustCompile(`.*\(`)
	operation := re.FindString(d.Fn)
	operation = strings.Trim(operation, "(")
	return operation
}
