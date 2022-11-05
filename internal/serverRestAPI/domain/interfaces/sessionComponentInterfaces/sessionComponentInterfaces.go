package sessionComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
)

type SessionUsecaseInterface interface {
	SessionExists(session *models.Session) (bool, error)
	CreateSessionForUser(email string, password string) (*models.Session, error)
	RemoveSession(session *models.Session) error
	GetUserBySession(session *models.Session) (*models.User, error)
}

type SessionRepositoryInterface interface {
	CreateSessionForUser(email string) (*models.Session, error)
	GetEmailBySession(session *models.Session) (string, error)
	RemoveSession(session *models.Session) error
	SessionExists(session *models.Session) (bool, error)
}
