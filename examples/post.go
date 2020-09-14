package examples

import (
	"errors"
	"net/http"
)

type GithubError struct {
	StatusCode       int    `json:"-"`
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}

type Repository struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Private     bool   `json:"private"`
}

func CreateRepo(request Repository) (*Repository, error) {
	response, err := httpClient.Post("https://api.github.com/user/repos", request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusCreated {
		var githubError GithubError
		if err := response.UnmarshalJson(&githubError); err != nil {
			return nil, errors.New("error processing github error response when creating a new repo")
		}
		return nil, errors.New(githubError.Message)
	}

	var result Repository
	if err := response.UnmarshalJson(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
