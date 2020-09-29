package collector

type API interface {
	GetContent(url, definitionContent string) (Result, error)
	GetContentFromFileDefinition(url, definitionPath string) (Result, error)
}
type api struct {
}

func NewAPI() API {
	return &api{}
}

func (a *api) GetContentFromFileDefinition(url, definitionPath string) (Result, error) {

	definition, err := NewDefinitionFromFile(definitionPath)
	if err != nil {
		return nil, err
	}
	return a.getResults(definition, url)
}

func (a *api) GetContent(url, definitionContent string) (Result, error) {
	definition, err := NewDefinition(definitionContent)
	if err != nil {
		return nil, err
	}
	return a.getResults(definition, url)
}

func (a *api) getResults(definition *definition, url string) (Result, error) {

	c := NewCollector()
	content, err := c.GetContent(url)

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	p := NewParser()
	r, err := p.Parse(definition, content)

	if err != nil {
		return r, err
	}

	filter := NewFilter()
	r = filter.Apply(r, definition)
	return r, nil
}
