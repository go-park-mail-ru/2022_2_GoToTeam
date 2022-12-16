package repository

import (
	"2022_2_GoTo_team/internal/userProfileService/domain/customErrors/profileComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/userProfileService/domain/interfaces/profileComponentInterfaces"
	"2022_2_GoTo_team/internal/userProfileService/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
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

func (ppsr *profilePostgreSQLRepository) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	ppsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetProfileByEmail function.")

	row := ppsr.database.QueryRow(`
SELECT email, login, COALESCE(username, ''), COALESCE(avatar_img_path, '')
FROM users WHERE email = $1;
`, email)

	profile := &models.Profile{}
	if err := row.Scan(&profile.Email, &profile.Login, &profile.Username, &profile.AvatarImgPath); err != nil {
		if err == sql.ErrNoRows {
			ppsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return nil, repositoryToUsecaseErrors.ProfileRepositoryEmailDoesntExistError
		}
		ppsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.ProfileRepositoryError
	}

	ppsr.logger.LogrusLoggerWithContext(ctx).Debugf("Got user profile: %#v", profile)

	return profile, nil
}

func (ppsr *profilePostgreSQLRepository) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string) error {
	ppsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UpdateProfileByEmail function.")

	if newProfile.Email != email {
		exists, err := ppsr.UserExistsByEmail(ctx, newProfile.Email)
		if err != nil {
			ppsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return repositoryToUsecaseErrors.ProfileRepositoryError
		}
		if exists {
			ppsr.logger.LogrusLoggerWithContext(ctx).Warnf("Email %s exists.", newProfile.Email)
			return repositoryToUsecaseErrors.ProfileRepositoryEmailExistsError
		}
	}

	exists, err := ppsr.UserExistsByLoginWithIgnoringRowsWithEmail(ctx, newProfile.Login, email)
	if err != nil {
		ppsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return repositoryToUsecaseErrors.ProfileRepositoryError
	}
	if exists {
		ppsr.logger.LogrusLoggerWithContext(ctx).Warnf("Login %s exists.", newProfile.Login)
		return repositoryToUsecaseErrors.ProfileRepositoryLoginExistsError
	}

	_, err = ppsr.database.Exec(`
UPDATE users SET email = $1, login = $2, username = (CASE WHEN $3 = '' THEN NULL ELSE $3 END), password = 
    (CASE WHEN $4 = '' THEN (SELECT password FROM users WHERE email = $6) ELSE $4 END)
WHERE email = $5;
`, newProfile.Email, newProfile.Login, newProfile.Username, newProfile.Password, email)

	if err != nil {
		ppsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return repositoryToUsecaseErrors.ProfileRepositoryError
	}

	ppsr.logger.LogrusLoggerWithContext(ctx).Debug("Profile has been updated successfully.")

	return nil
}

func (ppsr *profilePostgreSQLRepository) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	ppsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UserExistsByEmail function.")

	row := ppsr.database.QueryRow(`
SELECT U.email
FROM users U WHERE U.email = $1;
`, email)

	emailTmp := ""
	if err := row.Scan(&emailTmp); err != nil {
		if err == sql.ErrNoRows {
			ppsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return false, nil
		}
		ppsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return true, repositoryToUsecaseErrors.ProfileRepositoryError
	}

	ppsr.logger.LogrusLoggerWithContext(ctx).Debug("Got email: ", emailTmp)

	return true, nil
}

func (ppsr *profilePostgreSQLRepository) UserExistsByLoginWithIgnoringRowsWithEmail(ctx context.Context, login string, emailToIgnore string) (bool, error) {
	ppsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UserExistsByLoginWithIgnoringRowsWithEmail function.")

	row := ppsr.database.QueryRow(`
SELECT U.login
FROM users U WHERE U.login = $1 AND U.email != $2;
`, login, emailToIgnore)

	loginTmp := ""
	if err := row.Scan(&loginTmp); err != nil {
		if err == sql.ErrNoRows {
			ppsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return false, nil
		}
		ppsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return true, repositoryToUsecaseErrors.ProfileRepositoryError
	}

	ppsr.logger.LogrusLoggerWithContext(ctx).Debug("Got login: ", loginTmp)

	return true, nil
}
