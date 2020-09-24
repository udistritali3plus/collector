// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Collector is an autogenerated mock type for the Collector type
type Collector struct {
	mock.Mock
}

// GetContent provides a mock function with given fields: url
func (_m *Collector) GetContent(url string) (string, error) {
	ret := _m.Called(url)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
