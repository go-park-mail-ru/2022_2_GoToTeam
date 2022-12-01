package profileComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type ProfileUsecaseInterface interface {
	GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error)
	UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error
}

type ProfileRepositoryInterface interface {
	GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error)
	UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error
}
