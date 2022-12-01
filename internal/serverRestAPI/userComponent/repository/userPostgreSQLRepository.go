package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/userComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/userComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
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

	logger.LogrusLogger.Info("userPostgreSQLRepository has created.")

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

func (upsr *userPostgreSQLRepository) AddUser(ctx context.Context, email string, login string, username string, password string) (int, error) {
	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the AddUser function.")

	row := upsr.database.QueryRow(`
INSERT INTO users (email, login, username, password) VALUES ($1, $2, (CASE WHEN $3 = '' THEN NULL ELSE $3 END), $4) RETURNING user_id
`, email, login, username, password)

	var lastInsertId int
	if err := row.Scan(&lastInsertId); err != nil {
		upsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.UserRepositoryError
	}

	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Got lastInsertId: ", lastInsertId)

	return lastInsertId, nil
}

func (upsr *userPostgreSQLRepository) GetUserInfo(ctx context.Context, login string) (*models.User, error) {
	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetUserInfo function.")

	row := upsr.database.QueryRow(`
SELECT COALESCE(U.username, U.login), TO_CHAR(registration_date :: DATE, 'dd-mm-yyyy'), subscribers_count
FROM users U WHERE U.login = $1;
`, login)

	user := &models.User{}
	if err := row.Scan(&user.Username, &user.RegistrationDate, &user.SubscribersCount); err != nil {
		if err == sql.ErrNoRows {
			upsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return nil, repositoryToUsecaseErrors.UserRepositoryLoginDoesntExistError
		}
		upsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.UserRepositoryError
	}

	upsr.logger.LogrusLoggerWithContext(ctx).Debugf("Got user info: %#v", user)

	return user, nil
}

func (upsr *userPostgreSQLRepository) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UserExistsByEmail function.")

	row := upsr.database.QueryRow(`
SELECT U.email
FROM users U WHERE U.email = $1;
`, email)

	emailTmp := ""
	if err := row.Scan(&emailTmp); err != nil {
		if err == sql.ErrNoRows {
			upsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return false, nil
		}
		upsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return true, repositoryToUsecaseErrors.UserRepositoryError
	}

	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Got email: ", emailTmp)

	return true, nil
}

func (upsr *userPostgreSQLRepository) UserExistsByLogin(ctx context.Context, login string) (bool, error) {
	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UserExistsByLogin function.")

	row := upsr.database.QueryRow(`
SELECT U.login
FROM users U WHERE U.login = $1;
`, login)

	loginTmp := ""
	if err := row.Scan(&loginTmp); err != nil {
		if err == sql.ErrNoRows {
			upsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return false, nil
		}
		upsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return true, repositoryToUsecaseErrors.UserRepositoryError
	}

	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Got login: ", loginTmp)

	return true, nil
}
