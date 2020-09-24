package collector

import (
	"regexp"
	"strings"
)

const (
	// It's a regular expression to retrieve the content between () in string.
	parenthesesRegularExp = `\((.*?)\)`
)

type LookAheadParser interface {
	apply(expression string) []string
}

type lookAheadParser struct {
}

func NewLookAheadParser() LookAheadParser {
	return &lookAheadParser{}
}

func (p *lookAheadParser) apply(expression string) []string {

	re := regexp.MustCompile(parenthesesRegularExp)
	expressions := re.FindAllString(expression, -1)
	matches := make([]string, 0)

	for _, element := range expressions {
		element = strings.Trim(element, "(")
		element = strings.Trim(element, ")")
		matches = append(matches, element)
	}

	return matches
}
