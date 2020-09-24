package examples

import (
	"errors"
	"fmt"
	"github.com/federicoleon/go-httpclient/gohttp_mock"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start test cases for package 'examples'")

	// Tell the HTTP library to mock any further requests from here.
	gohttp_mock.MockupServer.Start()

	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		// Initialization:
		gohttp_mock.MockupServer.DeleteMocks()
		gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})

		// Execution:
		endpoints, err := GetEndpoints()

		// Validation:
		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		// Initialization:
		gohttp_mock.MockupServer.DeleteMocks()
		gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		})

		// Execution:
		endpoints, err := GetEndpoints()

		// Validation:
		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error was expected")
		}

		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		// Initialization:
		gohttp_mock.MockupServer.DeleteMocks()
		gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		})

		// Execution:
		endpoints, err := GetEndpoints()

		// Validation:
		if err != nil {
			t.Error(fmt.Sprintf("no error was expected and we got '%s'", err.Error()))
		}

		if endpoints == nil {
			t.Error("endpoints were expected and we got nil")
		}

		if endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("invalid current user url")
		}
	})
}
