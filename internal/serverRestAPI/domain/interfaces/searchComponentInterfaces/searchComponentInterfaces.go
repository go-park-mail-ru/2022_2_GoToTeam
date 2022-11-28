package searchComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type SearchUsecaseInterface interface {
	GetArticlesByTag(ctx context.Context, tag string) ([]*models.Article, error)
	GetArticlesBySearchParameters(ctx context.Context, substringToSearch string, login string, categoryName string, tagName string) ([]*models.Article, error)
}

type SearchRepositoryInterface interface {
	GetArticlesByTag(ctx context.Context, tag string) ([]*models.Article, error)
	GetArticlesBySearchParameters(ctx context.Context, substringToSearch string, login string, categoryName string, tagName string) ([]*models.Article, error)
	TagExists(ctx context.Context, tag string) (bool, error)
}
