package jira

import (
	"context"
	"net/http"
	"net/url"

	"github.com/r57ty7/jiracket/domain"
	"github.com/r57ty7/jiracket/domain/interfaces"
)

type searchRepository struct {
	client *Client
}

type SearchResults struct {
	Expand     string            `json:"expand,omitempty" yaml:"expand,omitempty"`
	Issues     []domain.Issue    `json:"issues,omitempty" yaml:"issues,omitempty"`
	MaxResults int               `json:"maxResults,omitempty" yaml:"maxResults,omitempty"`
	Names      map[string]string `json:"names,omitempty" yaml:"names,omitempty"`
	StartAt    int               `json:"startAt,omitempty" yaml:"startAt,omitempty"`
	Total      int               `json:"total,omitempty" yaml:"total,omitempty"`
}

// NewSearchRepository returns domain SearchRepository
func NewSearchRepository(client *Client) interfaces.SearchRepository {
	return &searchRepository{
		client: client,
	}
}

func (r *searchRepository) Search(ctx context.Context, jql string) ([]domain.Issue, error) {
	u := url.URL{
		Path: EndpointSearch,
	}
	uv := url.Values{}
	if jql != "" {
		uv.Add("jql", jql)
	}

	u.RawQuery = uv.Encode()

	req, err := r.client.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	searchResults := new(SearchResults)
	_, err = r.client.Do(req, searchResults)
	if err != nil {
		return searchResults.Issues, err
	}

	return searchResults.Issues, nil
}
