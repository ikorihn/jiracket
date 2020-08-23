package jira

import (
	"context"
	"net/http"

	"github.com/r57ty7/jiracket/domain"
	"github.com/r57ty7/jiracket/domain/interfaces"
)

type searchRepository struct {
	client Client
}

type SearchResults struct {
	Expand          string            `json:"expand,omitempty" yaml:"expand,omitempty"`
	Issues          []domain.Issue    `json:"issues,omitempty" yaml:"issues,omitempty"`
	MaxResults      int               `json:"maxResults,omitempty" yaml:"maxResults,omitempty"`
	Names           map[string]string `json:"names,omitempty" yaml:"names,omitempty"`
	Schema          JSONTypeMap       `json:"schema,omitempty" yaml:"schema,omitempty"`
	StartAt         int               `json:"startAt,omitempty" yaml:"startAt,omitempty"`
	Total           int               `json:"total,omitempty" yaml:"total,omitempty"`
	WarningMessages WarningMessages   `json:"warningMessages,omitempty" yaml:"warningMessages,omitempty"`
}

// NewSearchRepository returns domain SearchRepository
func NewSearchRepository(httpClient *http.Client, baseURL string) interfaces.SearchRepository {
	client := NewClient(httpClient, baseURL)
	return &searchRepository{
		client,
	}
}

func (r *searchRepository) Search(ctx context.Context, jql string) ([]Issue, error) {
	req, err := r.client.NewRequest(http.MethodGet, EndpointSearch, nil)

	searchResults := new(SearchResults)
	_, err := r.client.Do(req, searchResults)
	if err != nil {
		return err
	}

	return searchResults.Issues, nil
}
