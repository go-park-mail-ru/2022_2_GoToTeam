package articleComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type ArticleUsecaseInterface interface {
	GetArticleById(ctx context.Context, id int) (*models.Article, error)
	AddArticleBySession(ctx context.Context, article *models.Article, session *models.Session) error
	RemoveArticleById(ctx context.Context, articleId int) error
}

type ArticleRepositoryInterface interface {
	GetArticleById(ctx context.Context, id int) (*models.Article, error)
	GetTagsForArticle(ctx context.Context, articleId int) ([]string, error)
	AddArticle(ctx context.Context, article *models.Article) (int, error)
	DeleteArticleById(ctx context.Context, articleId int) (int64, error)
}
