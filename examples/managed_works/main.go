package main

import (
	"collector"
	"collector/fixtures"
	"fmt"
)

const (
	gruplac_url = "https://scienti.minciencias.gov.co/gruplac/jsp/visualiza/visualizagr.jsp?nro=00000000001394"
)

func main() {

	c := collector.NewCollector()
	content, err := c.GetContent(gruplac_url)

	if err != nil {
		return
	}

	definition, err := collector.NewDefinition(fixtures.DoctoralThesisContent)

	p := collector.NewParser()
	results, err := p.Parse(definition, content)

	for _, fields := range results {
		for field, value := range fields {
			fmt.Printf("%s: %s\n", field, value)
		}
		fmt.Println()
	}
}
