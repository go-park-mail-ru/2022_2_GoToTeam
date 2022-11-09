package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/profileComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/profileComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"context"
	"database/sql"
)

type profilePostgreSQLRepository struct {
	database *sql.DB
	logger   *logger.Logger
}

func NewProfilePostgreSQLRepository(database *sql.DB, logger *logger.Logger) profileComponentInterfaces.ProfileRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewProfilePostgreSQLRepository function.")

	profileRepository := &profilePostgreSQLRepository{
		database: database,
		logger:   logger,
	}

	logger.LogrusLogger.Info("profilePostgreSQLRepository has created.")

	return profileRepository
}

func (ppsr profilePostgreSQLRepository) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	ppsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetProfileByEmail function.")

	row := ppsr.database.QueryRow(`
SELECT email, login, username, avatar_img_path
FROM users WHERE email = $1;
`, email)

	profile := &models.Profile{}
	if err := row.Scan(&profile.Email, &profile.Login, &profile.Username, &profile.AvatarImgPath); err != nil {
		if err == sql.ErrNoRows {
			ppsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return nil, repositoryToUsecaseErrors.ProfileRepositoryEmailDontExistsError
		}
		ppsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.ProfileRepositoryError
	}

	ppsr.logger.LogrusLoggerWithContext(ctx).Debugf("Got user profile: %#v", profile)

	return profile, nil
}
