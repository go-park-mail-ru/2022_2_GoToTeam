package profileComponentInterfaces

import (
	"2022_2_GoTo_team/internal/userProfileService/domain/models"
	"context"
)

type ProfileUsecaseInterface interface {
	GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error)
	UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error
}

type ProfileRepositoryInterface interface {
	GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error)
	UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string) error
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
	UserExistsByLoginWithIgnoringRowsWithEmail(ctx context.Context, login string, emailToIgnore string) (bool, error)
}
