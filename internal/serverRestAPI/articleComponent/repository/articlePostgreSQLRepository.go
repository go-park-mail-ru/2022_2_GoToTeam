package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/articleComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
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
