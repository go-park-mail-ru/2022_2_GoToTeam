package articleComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type ArticleUsecaseInterface interface {
	GetArticleById(ctx context.Context, id int) (*models.Article, error)
}

type ArticleRepositoryInterface interface {
	GetArticleById(ctx context.Context, id int) (*models.Article, error)
	GetTagsForArticle(ctx context.Context, articleId int) ([]string, error)
}
