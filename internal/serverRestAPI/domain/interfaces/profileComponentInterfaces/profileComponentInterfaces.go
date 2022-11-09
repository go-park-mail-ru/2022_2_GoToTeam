package profileComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type ProfileUsecaseInterface interface {
	GetProfileBySession(ctx context.Context, session *models.Session) (*models.Profile, error)
}

type ProfileRepositoryInterface interface {
	GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error)
}
