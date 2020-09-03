package gomime

import "testing"

func TestHeaders(t *testing.T) {
	if HeaderContentType != "Content-Type" {
		t.Error("invalid content type header")
	}

	if HeaderUserAgent != "User-Agent" {
		t.Error("invalid user agent header")
	}

	if ContentTypeJson != "application/json" {
		t.Error("invalid content type json header")
	}
}
