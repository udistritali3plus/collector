package collector

type API interface {
	GetContent(url, definitionContent string) ([]map[string]string, error)
	GetContentFromFileDefinition(url, definitionPath string) ([]map[string]string, error)
}
type api struct {
}

func NewAPI() API {
	return &api{}
}

func (a *api) GetContentFromFileDefinition(url, definitionPath string) ([]map[string]string, error) {

	c := NewCollector()
	content, err := c.GetContent(url)

	if err != nil {
		return nil, err
	}

	definition, err := NewDefinitionFromFile(definitionPath)

	if err != nil {
		return nil, err
	}

	p := NewParser()
	return p.Parse(definition, content)
}

func (a *api) GetContent(url, definitionContent string) ([]map[string]string, error) {

	c := NewCollector()
	content, err := c.GetContent(url)

	if err != nil {
		return nil, err
	}

	definition, err := NewDefinition(definitionContent)

	if err != nil {
		return nil, err
	}

	p := NewParser()
	return p.Parse(definition, content)
}
