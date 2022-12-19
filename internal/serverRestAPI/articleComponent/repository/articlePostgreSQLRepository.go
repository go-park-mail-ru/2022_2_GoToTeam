package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/articleComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/articleComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"database/sql"
)

type articlePostgreSQLRepository struct {
	database *sql.DB
	logger   *logger.Logger
}

func NewArticlePostgreSQLRepository(database *sql.DB, logger *logger.Logger) articleComponentInterfaces.ArticleRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewArticlePostgreSQLRepository function.")

	articleRepository := &articlePostgreSQLRepository{
		database: database,
		logger:   logger,
	}

	logger.LogrusLogger.Info("articlePostgreSQLRepository has created.")

	return articleRepository
}

func (apsr *articlePostgreSQLRepository) GetArticleById(ctx context.Context, id int) (*models.Article, error) {
	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetArticleById function.")

	row := apsr.database.QueryRow(`
SELECT A.article_id,
       A.title,
       COALESCE(A.description, ''),
       A.rating,
       A.comments_count,
       A.content,
       COALESCE(A.cover_img_path, ''),
       COALESCE(COALESCE(UC.username, ''), ''),
       COALESCE(UC.login, ''),
       COALESCE(UP.username, ''),
       UP.login,
       COALESCE(C.category_name, '')
FROM articles A
         LEFT JOIN users UC ON A.co_author_id = UC.user_id
         JOIN users UP ON A.publisher_id = UP.user_id
         LEFT JOIN categories C ON A.category_id = C.category_id
WHERE A.article_id = $1;
`, id)

	article := &models.Article{}
	if err := row.Scan(
		&article.ArticleId,
		&article.Title,
		&article.Description,
		&article.Rating,
		&article.CommentsCount,
		&article.Content,
		&article.CoverImgPath,
		&article.CoAuthor.Username,
		&article.CoAuthor.Login,
		&article.Publisher.Username,
		&article.Publisher.Login,
		&article.CategoryName); err != nil {
		if err == sql.ErrNoRows {
			apsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return nil, repositoryToUsecaseErrors.ArticleRepositoryArticleDoesntExistError
		}
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.ArticleRepositoryError
	}

	tags, err := apsr.GetTagsForArticle(ctx, article.ArticleId)
	if err != nil {
		return nil, err
	}
	article.Tags = tags

	apsr.logger.LogrusLoggerWithContext(ctx).Debugf("Got article: %#v", article)

	return article, nil
}

func (apsr *articlePostgreSQLRepository) GetTagsForArticle(ctx context.Context, articleId int) ([]string, error) {
	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetTagsForArticle function.")

	tags := make([]string, 0, 10)
	rows, err := apsr.database.Query(`
SELECT T.tag_name
FROM tags T
         JOIN tags_articles TA ON T.tag_id = TA.tag_id
         JOIN articles A ON TA.article_id = A.article_id
WHERE A.article_id = $1;
`, articleId)
	if err != nil {
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.ArticleRepositoryError
	}
	defer rows.Close()

	for rows.Next() {
		tag := ""
		if err := rows.Scan(&tag); err != nil {
			apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.ArticleRepositoryError
		}

		tags = append(tags, tag)
	}

	return tags, nil
}

func (apsr *articlePostgreSQLRepository) AddArticle(ctx context.Context, article *models.Article) (int, error) {
	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the AddArticle function.")

	row := apsr.database.QueryRow(`
INSERT INTO articles (title, description, content, cover_img_path, co_author_id, publisher_id, category_id)  VALUES ($1, $2, $3, $4, 
        (SELECT user_id FROM users WHERE login = $5), 
        (SELECT user_id FROM users WHERE email = $6), 
        (SELECT categories.category_id FROM categories WHERE category_name = $7)) RETURNING article_id;
`, article.Title, article.Description, article.Content, article.CoverImgPath, article.CoAuthor.Login, article.Publisher.Email, article.CategoryName)

	var articleLastInsertId int
	if err := row.Scan(&articleLastInsertId); err != nil {
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.ArticleRepositoryError
	}

	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Got articleLastInsertId: ", articleLastInsertId)

	apsr.logger.LogrusLoggerWithContext(ctx).Debugf("Trying to add tags: %#v", article.Tags)

	for _, tagName := range article.Tags {
		row2 := apsr.database.QueryRow(`
INSERT INTO tags_articles (article_id, tag_id)  VALUES ($1, 
        (SELECT tag_id FROM tags WHERE tag_name = $2)) RETURNING article_id;
`, articleLastInsertId, tagName)

		var tagLastInsertArticleId int
		if err := row2.Scan(&tagLastInsertArticleId); err != nil {
			apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return 0, repositoryToUsecaseErrors.ArticleRepositoryError
		}

		apsr.logger.LogrusLoggerWithContext(ctx).Debug("Got tagLastInsertArticleId: ", tagLastInsertArticleId)
	}

	return articleLastInsertId, nil
}

func (apsr *articlePostgreSQLRepository) UpdateArticle(ctx context.Context, article *models.Article) error {
	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UpdateArticle function.")

	_, err := apsr.database.Exec(`
UPDATE articles SET title = $1, description = $2, content = $3, category_id = 
        (SELECT categories.category_id FROM categories WHERE category_name = $4)
WHERE article_id = $5;
`, article.Title, article.Description, article.Content, article.CategoryName, article.ArticleId)

	if err != nil {
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return repositoryToUsecaseErrors.ArticleRepositoryError
	}

	apsr.logger.LogrusLoggerWithContext(ctx).Debugf("Trying to delete old tags for the article: %#v", article.Tags)

	result, err := apsr.database.Exec(
		"DELETE FROM tags_articles WHERE article_id = $1 RETURNING *",
		article.ArticleId,
	)
	if err != nil {
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return repositoryToUsecaseErrors.ArticleRepositoryError
	}

	removedRowsCount, err := result.RowsAffected()
	if err != nil {
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return repositoryToUsecaseErrors.ArticleRepositoryError
	}
	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Removed tags for the article count: ", removedRowsCount)

	apsr.logger.LogrusLoggerWithContext(ctx).Debugf("Trying to add new tags for the article: %#v", article.Tags)

	for _, tagName := range article.Tags {
		row2 := apsr.database.QueryRow(`
INSERT INTO tags_articles (article_id, tag_id)  VALUES ($1, 
        (SELECT tag_id FROM tags WHERE tag_name = $2)) RETURNING article_id;
`, article.ArticleId, tagName)

		var tagLastInsertArticleId int
		if err := row2.Scan(&tagLastInsertArticleId); err != nil {
			apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return repositoryToUsecaseErrors.ArticleRepositoryError
		}

		apsr.logger.LogrusLoggerWithContext(ctx).Debug("Got tagLastInsertArticleId: ", tagLastInsertArticleId)
	}

	return nil
}

func (apsr *articlePostgreSQLRepository) DeleteArticleById(ctx context.Context, articleId int) (int64, error) {
	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the DeleteArticleById function.")

	result, err := apsr.database.Exec(
		"DELETE FROM articles WHERE article_id = $1 RETURNING *;",
		articleId,
	)
	if err != nil {
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.ArticleRepositoryError
	}

	removedRowsCount, err := result.RowsAffected()
	if err != nil {
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.ArticleRepositoryError
	}
	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Removed articles count: ", removedRowsCount)

	return removedRowsCount, nil
}

func (apsr *articlePostgreSQLRepository) GetAuthorEmailForArticle(ctx context.Context, articleId int) (string, error) {
	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetAuthorEmailForArticle function.")

	rows, err := apsr.database.Query(`
SELECT U.email
FROM users U
JOIN articles A on U.user_id = A.publisher_id
WHERE A.article_id = $1;
`, articleId)
	if err != nil {
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return "", repositoryToUsecaseErrors.ArticleRepositoryError
	}
	defer rows.Close()

	var email string
	isMultipleReturnedRows := false
	for rows.Next() {
		if isMultipleReturnedRows {
			apsr.logger.LogrusLoggerWithContext(ctx).Error("Multiple returning rows.")
			return "", repositoryToUsecaseErrors.ArticleRepositoryError
		}
		isMultipleReturnedRows = true

		if err := rows.Scan(&email); err != nil {
			apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return "", repositoryToUsecaseErrors.ArticleRepositoryError
		}
	}

	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Got email: ", email)

	return email, nil
}

func (apsr *articlePostgreSQLRepository) AddLike(ctx context.Context, isLike bool, articleId int, email string) (int, error) {
	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the AddLike function.")

	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Input isLike = ", isLike, " articleId = ", articleId, " email = ", email)

	row := apsr.database.QueryRow(`
INSERT INTO articles_likes (is_like, article_id, user_id)  VALUES ($1, $2,
        (SELECT user_id FROM users WHERE email = $3)) RETURNING article_id;
`, isLike, articleId, email)

	var likeLastInsertId int
	if err := row.Scan(&likeLastInsertId); err != nil {
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.ArticleRepositoryError
	}

	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Got likeLastInsertId: ", likeLastInsertId)

	return likeLastInsertId, nil
}

func (apsr *articlePostgreSQLRepository) RemoveLike(ctx context.Context, articleId int, email string) (int64, error) {
	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the RemoveLike function.")

	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Input articleId = ", articleId, " email = ", email)

	result, err := apsr.database.Exec(
		"DELETE FROM articles_likes WHERE article_id = $1 AND user_id = (SELECT user_id FROM users WHERE email = $2) RETURNING *;",
		articleId, email,
	)
	if err != nil {
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.ArticleRepositoryError
	}

	removedRowsCount, err := result.RowsAffected()
	if err != nil {
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.ArticleRepositoryError
	}
	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Removed likes count: ", removedRowsCount)

	return removedRowsCount, nil
}

func (apsr *articlePostgreSQLRepository) GetArticleRating(ctx context.Context, articleId int) (int, error) {
	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetArticleRating function.")

	rows, err := apsr.database.Query(`
SELECT rating
FROM articles
WHERE article_id = $1;
`, articleId)
	if err != nil {
		apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.ArticleRepositoryError
	}
	defer rows.Close()

	var rating int
	isMultipleReturnedRows := false
	for rows.Next() {
		if isMultipleReturnedRows {
			apsr.logger.LogrusLoggerWithContext(ctx).Error("Multiple returning rows.")
			return 0, repositoryToUsecaseErrors.ArticleRepositoryError
		}
		isMultipleReturnedRows = true

		if err := rows.Scan(&rating); err != nil {
			apsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return 0, repositoryToUsecaseErrors.ArticleRepositoryError
		}
	}

	apsr.logger.LogrusLoggerWithContext(ctx).Debug("Got rating: ", rating)

	return rating, nil
}
