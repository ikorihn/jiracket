package interfaces

import (
	"context"

	"github.com/r57ty7/jiracket/domain"
)

type SearchRepository interface {
	Search(ctx context.Context, jql string) ([]domain.Issue, error)
}
