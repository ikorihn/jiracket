package interfaces

import "context"

type SearchRepository interface {
	Search(ctx context.Context, jql string) ([]Issue, error)
}
