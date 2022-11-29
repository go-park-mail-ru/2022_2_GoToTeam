package repository

import (
	"2022_2_GoTo_team/internal/authSessionService/domain/customErrors/userComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/authSessionService/domain/interfaces/userComponentInterfaces"
	"2022_2_GoTo_team/internal/authSessionService/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"database/sql"
	"fmt"
)

type userPostgreSQLRepository struct {
	database *sql.DB
	logger   *logger.Logger
}

func NewUserPostgreSQLRepository(database *sql.DB, logger *logger.Logger) userComponentInterfaces.UserRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewUserPostgreSQLRepository function.")

	userRepository := &userPostgreSQLRepository{
		database: database,
		logger:   logger,
	}

	logger.LogrusLogger.Debug("All users in storage:  \n" + func() string {
		allUsers, err := userRepository.GetAllUsers(context.Background())
		if err != nil {
			return repositoryToUsecaseErrors.UserRepositoryError.Error()
		}
		return userRepository.getUsersString(allUsers)
	}())

	logger.LogrusLogger.Info("NewUserPostgreSQLRepository has created.")

	return userRepository
}

func (upsr *userPostgreSQLRepository) getUsersString(users []*models.User) string {
	usersString := ""
	for _, v := range users {
		usersString += fmt.Sprintf("%#v\n", v)
	}

	return usersString
}

func (upsr *userPostgreSQLRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetAllUsers function.")

	users := make([]*models.User, 0, 10)

	rows, err := upsr.database.Query(`
SELECT U.user_id,
       U.email,
       U.login,
       U.password,
       COALESCE(U.username, U.login),
       COALESCE(U.sex, 'U'),
       COALESCE(U.date_of_birth, '2000-01-01'),
       COALESCE(U.avatar_img_path, ''),
       U.registration_date,
       U.subscribers_count,
       U.subscriptions_count
FROM users U;
`)
	if err != nil {
		upsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.UserRepositoryError
	}
	defer rows.Close()

	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.UserId, &user.Email, &user.Login, &user.Password, &user.Username, &user.Sex, &user.DateOfBirth, &user.AvatarImgPath, &user.RegistrationDate, &user.SubscribersCount, &user.SubscriptionsCount); err != nil {
			upsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.UserRepositoryError
		}

		users = append(users, user)
	}

	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Got users: \n" + upsr.getUsersString(users))

	return users, nil
}

func (upsr *userPostgreSQLRepository) CheckUserEmailAndPassword(ctx context.Context, email string, password string) (bool, error) {
	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the CheckUserEmailAndPassword function.")

	row := upsr.database.QueryRow(`
SELECT U.email
FROM users U WHERE U.email = $1 AND U.password = $2;
`, email, password)

	emailTmp := ""
	if err := row.Scan(&emailTmp); err != nil {
		if err == sql.ErrNoRows {
			upsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return false, nil
		}
		upsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return false, repositoryToUsecaseErrors.UserRepositoryError
	}

	return true, nil
}

func (upsr *userPostgreSQLRepository) GetUserInfoForSessionComponentByEmail(ctx context.Context, email string) (*models.User, error) {
	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetUserInfoForSessionComponentByEmail function.")

	row := upsr.database.QueryRow(`
SELECT 
    COALESCE(U.username, U.login), 
    COALESCE(U.avatar_img_path, '')
FROM users U WHERE U.email = $1;
`, email)

	user := &models.User{}
	if err := row.Scan(&user.Username, &user.AvatarImgPath); err != nil {
		if err == sql.ErrNoRows {
			upsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return nil, repositoryToUsecaseErrors.UserRepositoryEmailDoesntExistError
		}
		upsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.UserRepositoryError
	}

	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Got user: %#v", user)

	return user, nil
}
