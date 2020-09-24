# Go HTTP Client

A production-ready HTTP client in Go with lots of useful features and using nothing more than the standard library of the language.

## Installation

```bash
# Go Modules
require github.com/federicoleon/go-httpclient
```

## Usage
In order to use the library for making HTTP calls you need to import the corresponding HTTP package:

```go
import "github.com/federicoleon/go-httpclient/gohttp"
```

## Configuring the client
Once you have imported the package, you can now start using the client. First you need to configure and build the client as you need:

```go
headers := make(http.Header)
headers.Set("Some-Common-Header", "value-for-all-requests")

// Create a new builder:
client := gohttp.NewBuilder().

	// You can set global headers to be used in every request made by this client:
	SetHeaders(headers).

	// Configure the timeout for getting a new connection:
	SetConnectionTimeout(2 * time.Second).

	// Configure the timeout for performing the actual HTTP call:
	SetResponseTimeout(3 * time.Second).

	// Configure the User-Agent header that will be used for all of the requests:
	SetUserAgent("Your-User-Agent").

	// Finally, build the client and start using it!
	Build()
```

## Performing HTTP calls
The ``Client`` interface provides convenient methods that you can use to perform different HTTP calls. If you get an error then you can safely ignore the response object since it won't be there.

**Important:** There is no need to read & close anything from the response since the client is doing all of this for you. You just need to get the response and start using it!

Take a look at all of the [EXAMPLES](examples) for more information.

### Get

```go
type Endpoints struct {
	CurrentUserUrl    string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

func GetEndpoints() (*Endpoints, error) {
	// Make the request and wait for the response:
	response, err := httpClient.Get("https://api.github.com")
	if err != nil {
		// Deal with the error as you need:
		return nil, err
	}

	// Interacting with the response:
	fmt.Println(fmt.Sprintf("Status Code: %d", response.StatusCode))
	fmt.Println(fmt.Sprintf("Status: %s", response.Status))
	fmt.Println(fmt.Sprintf("Body: %s\n", response.String()))

	// Processing JSON responses:
	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		// Deal with the unmarshal error as you need:
		return nil, err
	}

	fmt.Println(fmt.Sprintf("Repository URL: %s", endpoints.RepositoryUrl))
	return &endpoints, nil
}
```

### Post

```go

// The struct representing the actual JSON response from the API we're calling:
type GithubError struct {
	StatusCode       int    `json:"-"`
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}

// The struct representing the JSON body we're going to send:
type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Private     bool   `json:"private"`
}

func CreateRepo(request Repository) (*Repository, error) {
	// Make the request and wait for the response:
	response, err := httpClient.Post("https://api.github.com/user/repos", request)
	if err != nil {
		return nil, err
	}

	// Deal with failed status codes:
	if response.StatusCode != http.StatusCreated {
		var githubError GithubError
		if err := response.UnmarshalJson(&githubError); err != nil {
			return nil, errors.New("error processing github error response when creating a new repo")
		}
		return nil, errors.New(githubError.Message)
	}

	// Deal with successful response:
	var result Repository
	if err := response.UnmarshalJson(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

```

## Testing

The library provides a convenient package for mocking requests and getting a particular response. The mock key is generated using the ``HTTP method``, the ``request URL`` and the ``request body``. Every request with these same elements will return the same mock.

In order to use the mocking features you need to import the corresponding package:

```go
import "github.com/federicoleon/go-httpclient/gohttp_mock"
```

### Starting the mock server:
```go
func TestMain(m *testing.M) {
	// Tell the HTTP library to mock any further requests from here.
	gohttp_mock.MockupServer.Start()

	// Start the test cases for this pacakge:
	os.Exit(m.Run())
}
```

Once you start the mock server, every request will be handled by this server and will not be sent against the real API. If there is no mock matching the current request you'll get an error saying ``no mock matching {METHOD} from '{URL}' with given body``.

### Configuring a given HTTP mock:

```go
// Delete all mocks in every new test case to ensure a clean environment:
gohttp_mock.MockupServer.DeleteMocks()

// Configure a new mock:
gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
	Method:      http.MethodPost,
	Url:         "https://api.github.com/user/repos",
	RequestBody: `{"name":"test-repo","private":true}`,

	Error: errors.New("timeout from github"),
})
```

In this case, we're telling the client that when we send a POST request against that URL and with that body, we want that particular error. In this case, no response was returned. Let's see how you can configure a particular response:


```go
// Delete all mocks in every new test case to ensure a clean environment:
gohttp_mock.MockupServer.DeleteMocks()

// Configure a new mock:
gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
	Method:      http.MethodPost,
	Url:         "https://api.github.com/user/repos",
	RequestBody: `{"name":"test-repo","private":true}`,

	ResponseStatusCode: http.StatusCreated,
	ResponseBody:       `{"id":123,"name":"test-repo"}`,
})
```

In this case, we get a response with status code ``201 Created`` and that particular response body.

## Collaboration

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as needed. Changes with no test cases will NOT be reviewed!
