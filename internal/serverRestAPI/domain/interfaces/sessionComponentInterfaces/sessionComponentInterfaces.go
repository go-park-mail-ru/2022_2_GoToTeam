package sessionComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type SessionUsecaseInterface interface {
	SessionExists(ctx context.Context, session *models.Session) (bool, error)
	CreateSessionForUser(ctx context.Context, email string, password string) (*models.Session, error)
	RemoveSession(ctx context.Context, session *models.Session) error
	GetUserBySession(ctx context.Context, session *models.Session) (*models.User, error)
}

type SessionRepositoryInterface interface {
	CreateSessionForUser(ctx context.Context, email string) (*models.Session, error)
	GetEmailBySession(ctx context.Context, session *models.Session) (string, error)
	RemoveSession(ctx context.Context, session *models.Session) error
	SessionExists(ctx context.Context, session *models.Session) (bool, error)
}
