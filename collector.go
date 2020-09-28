package collector

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type Collector interface {
	GetContent(url string) (string, error)
}

type collector struct {
	client HTTPClient
}

func NewCollector() Collector {
	c := &http.Client{}
	return &collector{client: c}
}

func (c *collector) GetContent(url string) (string, error) {

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", errors.New(requestCreationFailed)
	}

	response, err := c.client.Do(request)

	if c.isError(response, err) {
		return "", errors.New(requestExecutionFailed)
	}

	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New(contentReadingFailed)
	}
	bodyString := string(bodyBytes)
	return bodyString, nil
}

func (c *collector) isError(response *http.Response, err error) bool {
	return err != nil || response.StatusCode == http.StatusGatewayTimeout
}
