package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "mocked-http-client")
	client.builder = &clientBuilder{
		headers:   commonHeaders,
		userAgent: "cool-user-agent",
	}

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(finalHeaders) != 3 {
		t.Error("we expect 3 headers")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("invalid request id received")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type received")
	}

	if finalHeaders.Get("User-Agent") != "mocked-http-client" {
		t.Error("invalid user agent received")
	}
}
