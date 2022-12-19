package articleComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type ArticleUsecaseInterface interface {
	GetArticleById(ctx context.Context, id int) (*models.Article, error)
	AddArticleBySession(ctx context.Context, article *models.Article) error
	UpdateArticle(ctx context.Context, article *models.Article) error
	RemoveArticleById(ctx context.Context, articleId int) error
	ProcessLike(ctx context.Context, likeData *models.LikeData) (int, error)
}

type ArticleRepositoryInterface interface {
	GetArticleById(ctx context.Context, id int, email string) (*models.Article, error)
	GetTagsForArticle(ctx context.Context, articleId int) ([]string, error)
	AddArticle(ctx context.Context, article *models.Article) (int, error)
	UpdateArticle(ctx context.Context, article *models.Article) error
	DeleteArticleById(ctx context.Context, articleId int) (int64, error)
	GetAuthorEmailForArticle(ctx context.Context, articleId int) (string, error)
	AddLike(ctx context.Context, isLike bool, articleId int, email string) (int, error)
	RemoveLike(ctx context.Context, articleId int, email string) (int64, error)
	GetArticleRating(ctx context.Context, articleId int) (int, error)
}
