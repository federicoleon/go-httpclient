package gomime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaders(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		header string
		value  string
	}{
		{
			name:   "test content type header",
			header: HeaderContentType,
			value:  "Content-Type",
		},
		{
			name:   "test user agent header",
			header: HeaderUserAgent,
			value:  "User-Agent",
		},
		{
			name:   "test content type json header",
			header: ContentTypeJson,
			value:  "application/json",
		},
		{
			name:   "test content type xml header",
			header: ContentTypeXml,
			value:  "application/xml",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase

		t.Run(testCase.name, func(tt *testing.T) {
			tt.Parallel()

			assert.EqualValues(tt, testCase.value, testCase.header)
		})
	}
}
