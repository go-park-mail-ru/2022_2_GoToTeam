package profileComponentInterfaces

import (
	"2022_2_GoTo_team/internal/userProfileService/domain/models"
	"context"
)

type ProfileUsecaseInterface interface {
	GetProfileBySession(ctx context.Context) (*models.Profile, error)
	UpdateProfileBySession(ctx context.Context, newProfile *models.Profile, session *models.Session) error
}

type ProfileRepositoryInterface interface {
	GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error)
	UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string) error
}
