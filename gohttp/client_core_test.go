package gohttp

import (
	"testing"
)

func TestGetRequestBody(t *testing.T) {
	// Initialization:
	client := httpClient{}

	t.Run("NoBodyNilResponse", func(t *testing.T) {
		// Execution
		body, err := client.getRequestBody("", nil)

		// Validation
		if err != nil {
			t.Error("no error expected when passing a nil body")
		}

		if body != nil {
			t.Error("no body expected when passing a nil body")
		}
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		// Execution
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("application/json", requestBody)

		// Validation
		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}

		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})

	t.Run("BodyWithXml", func(t *testing.T) {

	})

	t.Run("BodyWithJsonAsDefault", func(t *testing.T) {

	})
}
