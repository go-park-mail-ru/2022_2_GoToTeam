package commentaryComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type CommentaryUsecaseInterface interface {
	AddCommentaryBySession(ctx context.Context, commentary *models.Commentary) error
	GetAllCommentariesForArticle(ctx context.Context, articleId int) ([]*models.Commentary, error)
	ProcessLike(ctx context.Context, likeData *models.LikeData) (int, error)
}

type CommentaryRepositoryInterface interface {
	AddCommentaryByEmail(ctx context.Context, commentary *models.Commentary) (int, error)
	GetAllCommentsForArticle(ctx context.Context, articleId int) ([]*models.Commentary, error)
	AddLike(ctx context.Context, isLike bool, commentId int, email string) (int, error)
	RemoveLike(ctx context.Context, commentId int, email string) (int64, error)
	GetCommentaryRating(ctx context.Context, commentaryId int) (int, error)
}
