package feedComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type FeedUsecaseInterface interface {
	GetFeed(ctx context.Context) ([]*models.Article, error)
	GetFeedForUserByLogin(ctx context.Context, login string) ([]*models.Article, error)
	GetFeedForCategory(ctx context.Context, category string) ([]*models.Article, error)
}

type FeedRepositoryInterface interface {
	GetAllArticles(ctx context.Context) ([]*models.Article, error)
	GetArticles(ctx context.Context) ([]*models.Article, error)
	GetFeed(ctx context.Context) ([]*models.Article, error)
	GetFeedForUserByLogin(ctx context.Context, login string) ([]*models.Article, error)
	GetFeedForCategory(ctx context.Context, category string) ([]*models.Article, error)
	UserExistsByLogin(ctx context.Context, login string) (bool, error)
	CategoryExists(ctx context.Context, category string) (bool, error)
}
