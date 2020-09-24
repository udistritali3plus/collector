package collector

import (
	"cvlac/core/fixtures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDefinition(t *testing.T) {
	d, e := NewDefinition(fixtures.DoctoralThesisContent)
	assert.Nil(t, e)
	assert.NotNil(t, d)
	assert.NotNil(t, d.Name)
	assert.NotNil(t, d.Fields)

	assert.Equal(t, fixtures.DoctoralThesisDefinitionName, d.Name)
	expectedFields := 9
	assert.Equal(t, expectedFields, len(d.Fields))

	for _, f := range d.Fields {
		assert.NotEmpty(t, f.Name)
		assert.NotEmpty(t, f.Ex)
	}
}
