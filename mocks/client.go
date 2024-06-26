// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	"github.com/federicoleon/go-httpclient/core"

	"net/http"

	"github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Delete provides a mock function with given fields: url, headers
func (_m *Client) Delete(url string, headers ...http.Header) (*core.Response, error) {
	_va := make([]interface{}, len(headers))
	for _i := range headers {
		_va[_i] = headers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, url)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 *core.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...http.Header) (*core.Response, error)); ok {
		return rf(url, headers...)
	}
	if rf, ok := ret.Get(0).(func(string, ...http.Header) *core.Response); ok {
		r0 = rf(url, headers...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...http.Header) error); ok {
		r1 = rf(url, headers...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Do provides a mock function with given fields: req
func (_m *Client) Do(req *http.Request) (*core.Response, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Do")
	}

	var r0 *core.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*http.Request) (*core.Response, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*http.Request) *core.Response); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: url, headers
func (_m *Client) Get(url string, headers ...http.Header) (*core.Response, error) {
	_va := make([]interface{}, len(headers))
	for _i := range headers {
		_va[_i] = headers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, url)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *core.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...http.Header) (*core.Response, error)); ok {
		return rf(url, headers...)
	}
	if rf, ok := ret.Get(0).(func(string, ...http.Header) *core.Response); ok {
		r0 = rf(url, headers...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...http.Header) error); ok {
		r1 = rf(url, headers...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Options provides a mock function with given fields: url, headers
func (_m *Client) Options(url string, headers ...http.Header) (*core.Response, error) {
	_va := make([]interface{}, len(headers))
	for _i := range headers {
		_va[_i] = headers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, url)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Options")
	}

	var r0 *core.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...http.Header) (*core.Response, error)); ok {
		return rf(url, headers...)
	}
	if rf, ok := ret.Get(0).(func(string, ...http.Header) *core.Response); ok {
		r0 = rf(url, headers...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...http.Header) error); ok {
		r1 = rf(url, headers...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: url, body, headers
func (_m *Client) Patch(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	_va := make([]interface{}, len(headers))
	for _i := range headers {
		_va[_i] = headers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, url, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Patch")
	}

	var r0 *core.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string, interface{}, ...http.Header) (*core.Response, error)); ok {
		return rf(url, body, headers...)
	}
	if rf, ok := ret.Get(0).(func(string, interface{}, ...http.Header) *core.Response); ok {
		r0 = rf(url, body, headers...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string, interface{}, ...http.Header) error); ok {
		r1 = rf(url, body, headers...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Post provides a mock function with given fields: url, body, headers
func (_m *Client) Post(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	_va := make([]interface{}, len(headers))
	for _i := range headers {
		_va[_i] = headers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, url, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Post")
	}

	var r0 *core.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string, interface{}, ...http.Header) (*core.Response, error)); ok {
		return rf(url, body, headers...)
	}
	if rf, ok := ret.Get(0).(func(string, interface{}, ...http.Header) *core.Response); ok {
		r0 = rf(url, body, headers...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string, interface{}, ...http.Header) error); ok {
		r1 = rf(url, body, headers...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: url, body, headers
func (_m *Client) Put(url string, body interface{}, headers ...http.Header) (*core.Response, error) {
	_va := make([]interface{}, len(headers))
	for _i := range headers {
		_va[_i] = headers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, url, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Put")
	}

	var r0 *core.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string, interface{}, ...http.Header) (*core.Response, error)); ok {
		return rf(url, body, headers...)
	}
	if rf, ok := ret.Get(0).(func(string, interface{}, ...http.Header) *core.Response); ok {
		r0 = rf(url, body, headers...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string, interface{}, ...http.Header) error); ok {
		r1 = rf(url, body, headers...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
