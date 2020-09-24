package collector

import (
	"fmt"
	"regexp"
	"strings"
)

type FieldParser interface {
	parse(def *definition, document string) map[string]string
}

type fieldParser struct {
	lookAheadParser LookAheadParser
}

func NewItemParser() FieldParser {
	lookAheadParser := NewLookAheadParser()
	return &fieldParser{lookAheadParser}
}

func (p *fieldParser) parse(def *definition, document string) map[string]string {

	results := make(map[string]string, 0)

	fields := def.Fields
	for _, f := range fields {

		content := p.getFieldValue(f.Ex, document)
		if def.Parser.LookAheadManual {
			content = p.applyLookAhead(f.Ex, content)
		}

		name := p.getFieldName(f.Name)
		results[strings.ToLower(name)] = content
	}
	return results
}

func (p *fieldParser) getFieldValue(ex, document string) string {
	exp, err := regexp.Compile(ex)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	content := exp.FindString(document)
	if len(content) <= 0 {
		return ""
	}
	return content
}

func (p *fieldParser) getFieldName(name string) string {
	return strings.ReplaceAll(name, " ", "_")
}

func (p *fieldParser) applyLookAhead(ex, text string) string {

	expressions := p.lookAheadParser.apply(ex)
	for _, e := range expressions {
		text = strings.ReplaceAll(text, e, "")
	}
	return text
}
