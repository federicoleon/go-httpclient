package gohttp

import (
	"fmt"
	"net/http"
)

// Mock structure provides a clean way to configure HTTP mocks based on
// the combination between request method, URL and request body.
type Mock struct {
	Method      string
	Url         string
	RequestBody string

	Error              error
	ResponseBody       string
	ResponseStatusCode int
}

// GetResponse returns a Response object based on the mock configuration.
func (m *Mock) GetResponse() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := Response{
		status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		statusCode: m.ResponseStatusCode,
		body:       []byte(m.ResponseBody),
	}
	return &response, nil
}
