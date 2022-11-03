package interfaces

import "2022_2_GoTo_team/internal/serverRestAPI/domain/models"

type UserUsecaseInterface interface {
}

type UserRepositoryInterface interface {
	PrintUsers()
	AddUser(user *models.User) error
	UserIsExistByLogin(login string) bool
	UserIsExistByEmail(email string) bool
	GetUserByLogin(login string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	CreateUserInstanceFromData(username string, email string, login string, password string) *models.User
}
