package collector

import (
	"collector/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type collectorSuite struct {
	suite.Suite
	client    *mocks.HTTPClient
	collector Collector
}

func TestInitCollector(t *testing.T) {
	suite.Run(t, new(collectorSuite))
}

func (s *collectorSuite) SetupTest() {
	s.client = &mocks.HTTPClient{}
	s.collector = &collector{client: s.client}
}

func (s *collectorSuite) TestCollector_GetContent() {

	s.client.Mock.On("Do", mock.Anything).
		Return().
		Once()

	s.collector.GetContent("https://coins.com")
}
