package main

import (
	"collector"
	"fmt"
)

const (
	gruplacURL                 = "https://scienti.minciencias.gov.co/gruplac/jsp/visualiza/visualizagr.jsp?nro=00000000001394"
	managedWorksDefinitionPath = "examples/managed_works/managed_works_definition.yml"
)

func main() {

	c := collector.NewCollector()
	content, err := c.GetContent(gruplacURL)

	if err != nil {
		return
	}

	definition, err := collector.NewDefinitionFromFile(managedWorksDefinitionPath)

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
