package userComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type UserUsecaseInterface interface {
	AddNewUser(ctx context.Context, email string, login string, username string, password string) error
	GetUserInfo(ctx context.Context, login string) (*models.User, error)
	IsUserSubscribedOnUser(ctx context.Context, login string) (bool, error)
}

type UserRepositoryInterface interface {
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	AddUser(ctx context.Context, email string, login string, username string, password string) (int, error)
	GetUserInfo(ctx context.Context, login string) (*models.User, error)
	IsUserSubscribedOnUser(ctx context.Context, sessionEmail string, login string) (bool, error)
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
	UserExistsByLogin(ctx context.Context, login string) (bool, error)
}
