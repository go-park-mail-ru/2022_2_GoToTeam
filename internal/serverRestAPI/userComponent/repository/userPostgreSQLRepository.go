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

func (upsr *userPostgreSQLRepository) IsUserSubscribedOnUser(ctx context.Context, sessionEmail string, login string) (bool, error) {
	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the IsUserSubscribedOnUser function.")

	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Input sessionEmail = ", sessionEmail, " login = ", login)

	row := upsr.database.QueryRow(`
SELECT COUNT(*) count
FROM users U1
JOIN subscriptions S ON U1.user_id = S.subscribed_to_id
JOIN users U2 ON U2.user_id = S.user_id 
WHERE U2.email = $1 AND U1.login = $2;
`, sessionEmail, login)

	entriesFound := 0
	if err := row.Scan(&entriesFound); err != nil {
		upsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return false, repositoryToUsecaseErrors.UserRepositoryError
	}

	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Got entriesFound: ", entriesFound)

	result := false
	if entriesFound == 1 {
		result = true
	}

	return result, nil
}

func (upsr *userPostgreSQLRepository) SubscribeOnUser(ctx context.Context, email string, subscribeToLogin string) error {
	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the SubscribeOnUser function.")

	row := upsr.database.QueryRow(`
INSERT INTO subscriptions (user_id, subscribed_to_id) VALUES 
       ((SELECT user_id FROM users WHERE email = $1), (SELECT user_id FROM users WHERE login = $2)) RETURNING user_id, subscribed_to_id;
`, email, subscribeToLogin)

	var lastInsertedUserId int
	var lastInsertedSubscribedToId int
	if err := row.Scan(&lastInsertedUserId, &lastInsertedSubscribedToId); err != nil {
		upsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return repositoryToUsecaseErrors.UserRepositoryError
	}

	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Got lastInsertedUserId: ", lastInsertedUserId, " lastInsertedSubscribedToId:", lastInsertedSubscribedToId)

	return nil
}

func (upsr *userPostgreSQLRepository) UnsubscribeFromUser(ctx context.Context, email string, unsubscribeFromLogin string) (int64, error) {
	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UnsubscribeFromUser function.")

	result, err := upsr.database.Exec(`
DELETE FROM subscriptions WHERE 
                              user_id IN (SELECT user_id FROM users WHERE email = $1) AND 
                              subscribed_to_id IN (SELECT user_id FROM users WHERE login = $2)
							RETURNING *;
`, email, unsubscribeFromLogin)

	if err != nil {
		upsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.UserRepositoryError
	}

	removedRowsCount, err := result.RowsAffected()
	if err != nil {
		upsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, repositoryToUsecaseErrors.UserRepositoryError
	}
	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Removed subscriptions count: ", removedRowsCount)

	return removedRowsCount, nil
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
