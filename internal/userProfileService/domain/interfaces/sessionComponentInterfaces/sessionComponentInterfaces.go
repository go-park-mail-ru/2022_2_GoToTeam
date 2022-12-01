package sessionComponentInterfaces

import (
	"2022_2_GoTo_team/internal/userProfileService/domain/models"
	"context"
)

type SessionRepositoryInterface interface {
	UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error
}
