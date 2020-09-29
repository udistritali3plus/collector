package collector

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type definition struct {
	Name    string             `yaml:"name"`
	Path    string             `yaml:"path"`
	Fields  []FieldDefinition  `yaml:"fields"`
	Parser  ParserDefinition   `yaml:"parser"`
	Filters []FilterDefinition `yaml:"filters"`
}

type FieldDefinition struct {
	Name string `yaml:"name"`
	Ex   string `yaml:"ex"`
}

func NewDefinition(content string) (*definition, error) {
	definition := &definition{}

	err := yaml.Unmarshal([]byte(content), definition)
	if err != nil {
		message := fmt.Sprintf("%s, err:%v", UnmarshallDefinitionError, err)
		return nil, errors.New(message)
	}

	return definition, nil
}

func NewDefinitionFromFile(path string) (*definition, error) {
	yamlFile, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, errors.New(FileNotFoundError)
	}

	definition := &definition{}
	err = yaml.Unmarshal(yamlFile, definition)
	if err != nil {
		return nil, errors.New(UnmarshallDefinitionError)
	}

	return definition, nil
}
