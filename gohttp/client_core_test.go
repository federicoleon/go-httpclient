package gohttp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequestBody(t *testing.T) {
	t.Parallel()

	// Initialization:
	client := httpClient{}

	t.Run("no body nil response", func(tt *testing.T) {
		tt.Parallel()

		// Execution
		body, err := client.getRequestBody("", nil)

		assert.Nil(tt, err)
		assert.Nil(tt, body)
	})

	t.Run("body with json", func(tt *testing.T) {
		tt.Parallel()

		// Execution
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("application/json", requestBody)

		// Validation
		assert.Nil(tt, err)
		assert.NotNil(tt, body)
		assert.EqualValues(tt, `["one","two"]`, string(body))
	})

	t.Run("body with xml", func(tt *testing.T) {
		tt.Parallel()

		// Execution
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("application/xml", requestBody)

		// Validation
		assert.Nil(tt, err)
		assert.NotNil(tt, body)
		assert.EqualValues(tt, `<string>one</string><string>two</string>`, string(body))
	})

	t.Run("body with json as default", func(tt *testing.T) {
		tt.Parallel()

		// Execution
		requestBody := []string{"one", "two"}

		body, err := client.getRequestBody("", requestBody)

		// Validation
		assert.Nil(tt, err)
		assert.NotNil(tt, body)
		assert.EqualValues(tt, `["one","two"]`, string(body))
	})
}
