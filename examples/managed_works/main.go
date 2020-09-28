package main

import (
	"collector"
	"fmt"
)

const (
	gruplac_url                   = "https://scienti.minciencias.gov.co/managed_works/jsp/visualiza/visualizagr.jsp?nro=00000000001394"
	managed_works_definition_paht = "examples/managed_works/managed_works_definition.yml"
)

func main() {

	c := collector.NewCollector()
	content, err := c.GetContent(gruplac_url)

	if err != nil {
		return
	}

	definition, err := collector.NewDefinitionFromFile(managed_works_definition_paht)

	if err != nil {
		panic(err)
	}

	p := collector.NewParser()
	results, err := p.Parse(definition, content)

	if err != nil {
		panic(err)
	}

	for _, fields := range results {
		for field, value := range fields {
			fmt.Printf("%s: %s\n", field, value)
		}
		fmt.Println()
	}
}
