package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/tagComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/tagComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"database/sql"
)

type tagPostgreSQLRepository struct {
	database *sql.DB
	logger   *logger.Logger
}

func NewTagPostgreSQLRepository(database *sql.DB, logger *logger.Logger) tagComponentInterfaces.TagRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewTagPostgreSQLRepository function.")

	tagRepository := &tagPostgreSQLRepository{
		database: database,
		logger:   logger,
	}

	logger.LogrusLogger.Info("tagPostgreSQLRepository has created.")

	return tagRepository
}

func (tpsr *tagPostgreSQLRepository) GetAllTags(ctx context.Context) ([]*models.Tag, error) {
	tpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetAllTags function.")

	tags := make([]*models.Tag, 0, 10)

	rows, err := tpsr.database.Query(`
SELECT tag_name
FROM tags;
`)
	if err != nil {
		tpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.TagRepositoryError
	}
	defer rows.Close()

	for rows.Next() {
		tag := &models.Tag{}
		if err := rows.Scan(&tag.TagName); err != nil {
			tpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.TagRepositoryError
		}

		tags = append(tags, tag)
	}

	tpsr.logger.LogrusLoggerWithContext(ctx).Debugf("Got tags: %#v\n", tags)

	return tags, nil
}
