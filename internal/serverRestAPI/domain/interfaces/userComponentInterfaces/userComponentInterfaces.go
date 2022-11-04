package userComponentInterfaces

import "2022_2_GoTo_team/internal/serverRestAPI/domain/models"

type UserUsecaseInterface interface {
	AddNewUser(email string, login string, username string, password string) error
}

type UserRepositoryInterface interface {
	PrintUsers()
	AddUser(user *models.User) error
	UserExistsByLogin(login string) bool
	UserExistsByEmail(email string) bool
	GetUserByLogin(login string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}
