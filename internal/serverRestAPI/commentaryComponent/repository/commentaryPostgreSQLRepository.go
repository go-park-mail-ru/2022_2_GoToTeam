package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/commentaryComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/commentaryComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"database/sql"
	"strconv"
)

type commentaryPostgreSQLRepository struct {
	database *sql.DB
	logger   *logger.Logger
}

func NewCommentaryPostgreSQLRepository(database *sql.DB, logger *logger.Logger) commentaryComponentInterfaces.CommentaryRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewCommentaryPostgreSQLRepository function.")

	commentaryRepository := &commentaryPostgreSQLRepository{
		database: database,
		logger:   logger,
	}

	logger.LogrusLogger.Info("commentaryPostgreSQLRepository has created.")

	return commentaryRepository
}

func (cpsr *commentaryPostgreSQLRepository) AddCommentaryByEmail(ctx context.Context, commentary *models.Commentary) (int, error) {
	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the AddCommentaryByEmail function.")

	cpsr.logger.LogrusLoggerWithContext(ctx).Debugf("Commentary to add: %#v", commentary)

	if commentary.CommentForCommentId == "" {
		commentary.CommentForCommentId = "-1"
	}
	commentForCommentId, err := strconv.Atoi(commentary.CommentForCommentId)
	if err != nil {
		cpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.CommentaryRepositoryError
	}

	cpsr.logger.LogrusLoggerWithContext(ctx).Debugf("Converted commentForCommentId: %#v", commentForCommentId)

	row := cpsr.database.QueryRow(`
INSERT INTO comments (content, publisher_id, article_id, comment_for_comment_id)
VALUES ($1, (SELECT user_id FROM users WHERE email = $2), $3, (CASE WHEN $4 = -1 THEN NULL ELSE $4 END))
RETURNING article_id;
`, commentary.Content, commentary.Publisher.Email, commentary.ArticleId, commentForCommentId)

	var commentaryLastInsertId int
	if err := row.Scan(&commentaryLastInsertId); err != nil {
		cpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.CommentaryRepositoryError
	}

	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Got commentaryLastInsertId: ", commentaryLastInsertId)

	return commentaryLastInsertId, nil
}

func (cpsr *commentaryPostgreSQLRepository) GetAllCommentsForArticle(ctx context.Context, articleId int, email string) ([]*models.Commentary, error) {
	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetAllCommentsForArticle function.")

	cpsr.logger.LogrusLoggerWithContext(ctx).Debugf("Input articleId: %#v", articleId)

	commentaries := make([]*models.Commentary, 0, 10)

	rows, err := cpsr.database.Query(`
SELECT C.comment_id,
       C.content,
       C.rating,
       C.article_id,
       COALESCE(C.comment_for_comment_id::text, ''),
       COALESCE(UP.username, ''),
       UP.login,
       (CASE WHEN CL.is_like = true THEN 1 ELSE (CASE WHEN CL.is_like = false THEN -1 ELSE 0 END) END) liked
FROM comments C
         JOIN users UP ON C.publisher_id = UP.user_id
         LEFT JOIN comments_likes CL ON CL.user_id = (SELECT user_id FROM users WHERE email = $2) AND CL.comment_id = C.comment_id
WHERE C.article_id = $1;
`, articleId, email)
	if err != nil {
		cpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.CommentaryRepositoryError
	}
	defer rows.Close()

	for rows.Next() {
		commentary := &models.Commentary{}
		if err := rows.Scan(&commentary.CommentId, &commentary.Content, &commentary.Rating, &commentary.ArticleId, &commentary.CommentForCommentId, &commentary.Publisher.Username, &commentary.Publisher.Login, &commentary.Liked); err != nil {
			cpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.CommentaryRepositoryError
		}

		commentaries = append(commentaries, commentary)
	}

	cpsr.logger.LogrusLoggerWithContext(ctx).Debugf("Got commentaries for article:%#v \n", commentaries)

	return commentaries, nil
}

func (cpsr *commentaryPostgreSQLRepository) AddLike(ctx context.Context, isLike bool, commentId int, email string) (int, error) {
	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the AddLike function.")

	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Input isLike = ", isLike, " commentId = ", commentId, " email = ", email)

	row := cpsr.database.QueryRow(`
INSERT INTO comments_likes (is_like, comment_id, user_id)  VALUES ($1, $2,
        (SELECT user_id FROM users WHERE email = $3)) RETURNING comment_id;
`, isLike, commentId, email)

	var commentLastInsertId int
	if err := row.Scan(&commentLastInsertId); err != nil {
		cpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.CommentaryRepositoryError
	}

	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Got commentLastInsertId: ", commentLastInsertId)

	return commentLastInsertId, nil
}

func (cpsr *commentaryPostgreSQLRepository) RemoveLike(ctx context.Context, commentId int, email string) (int64, error) {
	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the RemoveLike function.")

	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Input commentId = ", commentId, " email = ", email)

	result, err := cpsr.database.Exec(
		"DELETE FROM comments_likes WHERE comment_id = $1 AND user_id = (SELECT user_id FROM users WHERE email = $2) RETURNING *;",
		commentId, email,
	)
	if err != nil {
		cpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.CommentaryRepositoryError
	}

	removedRowsCount, err := result.RowsAffected()
	if err != nil {
		cpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.CommentaryRepositoryError
	}
	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Removed likes count: ", removedRowsCount)

	return removedRowsCount, nil
}

func (cpsr *commentaryPostgreSQLRepository) GetCommentaryRating(ctx context.Context, commentaryId int) (int, error) {
	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetCommentaryRating function.")

	rows, err := cpsr.database.Query(`
SELECT rating
FROM comments
WHERE comment_id = $1;
`, commentaryId)
	if err != nil {
		cpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.CommentaryRepositoryError
	}
	defer rows.Close()

	var rating int
	isMultipleReturnedRows := false
	for rows.Next() {
		if isMultipleReturnedRows {
			cpsr.logger.LogrusLoggerWithContext(ctx).Error("Multiple returning rows.")
			return 0, repositoryToUsecaseErrors.CommentaryRepositoryError
		}
		isMultipleReturnedRows = true

		if err := rows.Scan(&rating); err != nil {
			cpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return 0, repositoryToUsecaseErrors.CommentaryRepositoryError
		}
	}

	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Got rating: ", rating)

	return rating, nil
}
