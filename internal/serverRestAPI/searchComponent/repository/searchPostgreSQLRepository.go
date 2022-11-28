package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/searchComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/searchComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/logger"
	"context"
	"database/sql"
	"fmt"
)

type searchPostgreSQLRepository struct {
	database *sql.DB
	logger   *logger.Logger
}

func NewSearchPostgreSQLRepository(database *sql.DB, logger *logger.Logger) searchComponentInterfaces.SearchRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewSearchPostgreSQLRepository function.")

	searchRepository := &searchPostgreSQLRepository{
		database: database,
		logger:   logger,
	}

	logger.LogrusLogger.Info("searchPostgreSQLRepository has created.")

	return searchRepository
}

func (spsr *searchPostgreSQLRepository) getArticlesString(articles []*models.Article) string {
	articlesString := ""
	for _, v := range articles {
		articlesString += fmt.Sprintf("%#v\n", v)
	}

	return articlesString
}

// GetArticlesByTag TODO OFFSET LIMIT
func (spsr *searchPostgreSQLRepository) GetArticlesByTag(ctx context.Context, tag string) ([]*models.Article, error) {
	spsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetArticlesByTag function.")

	articles := make([]*models.Article, 0, 10)

	rows, err := spsr.database.Query(`
SELECT A.article_id,
       A.title,
       COALESCE(A.description, ''),
       A.rating,
       A.comments_count,
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
		 JOIN tags_articles TA ON A.article_id = TA.article_id
		 JOIN tags T ON T.tag_id = TA.tag_id
WHERE LOWER(T.tag_name) = LOWER($1);
`, tag)
	if err != nil {
		spsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.SearchRepositoryError
	}
	defer rows.Close()

	for rows.Next() {
		article := &models.Article{}
		if err := rows.Scan(&article.ArticleId, &article.Title, &article.Description, &article.Rating, &article.CommentsCount, &article.CoverImgPath, &article.CoAuthor.Username, &article.CoAuthor.Login, &article.Publisher.Username, &article.Publisher.Login, &article.CategoryName); err != nil {
			spsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.SearchRepositoryError
		}
		article.Tags, err = spsr.GetTagsForArticle(ctx, article.ArticleId)
		if err != nil {
			spsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.SearchRepositoryError
		}

		articles = append(articles, article)
	}

	spsr.logger.LogrusLoggerWithContext(ctx).Debug("Got articles: \n" + spsr.getArticlesString(articles))

	return articles, nil
}

func (spsr *searchPostgreSQLRepository) GetArticlesBySearchParameters(ctx context.Context, substringToSearch string, login string, categoryName string, tagName string) ([]*models.Article, error) {
	spsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetArticlesBySearchParameters function.")

	articles := make([]*models.Article, 0, 10)

	spsr.logger.LogrusLoggerWithContext(ctx).Debug("Input parameters to search: substringToSearch = ", substringToSearch, " login = ", login, " categoryName = ", categoryName, " tagName = ", tagName)
	if substringToSearch == "" {
		substringToSearch = "%"
	} else {
		substringToSearch = "%" + substringToSearch + "%"
	}
	if login == "" {
		login = "%"
	}
	if categoryName == "" {
		categoryName = "%"
	}
	if tagName == "" {
		tagName = "%"
	}
	spsr.logger.LogrusLoggerWithContext(ctx).Debug("Formed regExp parameters to search: substringToSearch = ", substringToSearch, " login = ", login, " categoryName = ", categoryName, " tagName = ", tagName)

	rows, err := spsr.database.Query(`
SELECT A.article_id,
       A.title,
       COALESCE(A.description, ''),
       A.rating,
       A.comments_count,
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
         JOIN tags_articles TA ON A.article_id = TA.article_id
         JOIN tags T ON T.tag_id = TA.tag_id
WHERE ((LOWER(A.title) LIKE LOWER($1)) OR (LOWER(A.content) LIKE LOWER($1)))
  AND (UP.login LIKE $2)
  AND (C.category_name LIKE $3)
  AND (T.tag_name LIKE $4);
`, substringToSearch, login, categoryName, tagName)

	if err != nil {
		spsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.SearchRepositoryError
	}
	defer rows.Close()

	for rows.Next() {
		article := &models.Article{}
		if err := rows.Scan(&article.ArticleId, &article.Title, &article.Description, &article.Rating, &article.CommentsCount, &article.CoverImgPath, &article.CoAuthor.Username, &article.CoAuthor.Login, &article.Publisher.Username, &article.Publisher.Login, &article.CategoryName); err != nil {
			spsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.SearchRepositoryError
		}
		article.Tags, err = spsr.GetTagsForArticle(ctx, article.ArticleId)
		if err != nil {
			spsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.SearchRepositoryError
		}

		articles = append(articles, article)
	}

	spsr.logger.LogrusLoggerWithContext(ctx).Debug("Got articles: \n" + spsr.getArticlesString(articles))

	return articles, nil
}

func (spsr *searchPostgreSQLRepository) GetTagsForArticle(ctx context.Context, articleId int) ([]string, error) {
	spsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetTagsForArticle function.")

	tags := make([]string, 0, 10)
	rows, err := spsr.database.Query(`
SELECT T.tag_name
FROM tags T
         JOIN tags_articles TA ON T.tag_id = TA.tag_id
         JOIN articles A ON TA.article_id = A.article_id
WHERE A.article_id = $1;
`, articleId)
	if err != nil {
		spsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.SearchRepositoryError
	}
	defer rows.Close()

	for rows.Next() {
		tag := ""
		if err := rows.Scan(&tag); err != nil {
			spsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.SearchRepositoryError
		}

		tags = append(tags, tag)
	}

	return tags, nil
}

func (spsr *searchPostgreSQLRepository) TagExists(ctx context.Context, tag string) (bool, error) {
	spsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the TagExists function.")

	row := spsr.database.QueryRow(`
SELECT T.tag_name
FROM tags T WHERE LOWER(T.tag_name) = LOWER($1);
`, tag)

	tagTmp := ""
	if err := row.Scan(&tagTmp); err != nil {
		if err == sql.ErrNoRows {
			spsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return false, nil
		}
		spsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return true, repositoryToUsecaseErrors.SearchRepositoryError
	}

	spsr.logger.LogrusLoggerWithContext(ctx).Debug("Got tag: ", tagTmp)

	return true, nil
}
