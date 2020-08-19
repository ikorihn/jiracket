package client

import (
	"github.com/andygrunwald/go-jira"
)

type JiraClient struct {
	url string
}

func NewJiraClient(url string) *JiraClient {
	return &JiraClient{
		url: url,
	}
}

func (jc *JiraClient) Search(jql string) error {
	tp := jira.BasicAuthTransport{
		Username: "user",
		Password: "pas",
	}

	client, err := jira.NewClient(tp.Client(), jc.url)
	if err != nil {
		return err
	}
	_, _, err = client.Issue.Search(jql, &jira.SearchOptions{})
	if err != nil {
		return err
	}
	return nil
}
