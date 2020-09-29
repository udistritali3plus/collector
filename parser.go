package collector

import (
	"errors"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

type Parser interface {
	Parse(def *definition, content string) (Result, error)
}

type parser struct {
	fieldParser FieldParser
}

func NewParser() Parser {
	item := NewItemParser()
	return &parser{item}
}

func (p *parser) Parse(def *definition, content string) (Result, error) {
	results := make([]map[string]string, 0)
	items, err := p.getItems(def.Path, content)
	if err != nil {
		return results, err
	}
	results = p.parseItems(items, def)
	return results, nil
}

func (p *parser) getItems(path, content string) ([]*html.Node, error) {

	doc, err := htmlquery.Parse(strings.NewReader(content))
	if err != nil {
		return nil, errors.New(contentParsingFailed)
	}

	items := htmlquery.Find(doc, path)
	return items, nil
}

func (p *parser) parseItems(items []*html.Node, def *definition) Result {

	fields := make([]map[string]string, 0)

	for _, item := range items {
		content := htmlquery.OutputHTML(item, true)
		field := p.fieldParser.parse(def, content)
		fields = append(fields, field)
	}
	return fields
}
