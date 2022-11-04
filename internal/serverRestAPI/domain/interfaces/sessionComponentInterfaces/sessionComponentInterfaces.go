package sessionComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
)

type SessionUsecaseInterface interface {
	IsSessionExists(session *models.Session) bool
	CreateSessionForUser(email string, password string, sessionHeaderName string) (*models.Session, error)
	RemoveSession(session *models.Session)
	GetUserBySession(session *models.Session) (*models.User, error)
}

type SessionRepositoryInterface interface {
	PrintSessions()
	CreateSessionForUser(email string, sessionHeaderName string) *models.Session
	GetEmailBySession(session *models.Session) string
	RemoveSession(session *models.Session)
	SessionExists(session *models.Session) bool
	ExpireSession(session *models.Session)
}
