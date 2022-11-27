package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/sessionComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/sessionUtils"
	"2022_2_GoTo_team/pkg/logger"
	"context"
	"sync"
)

type sessionsStorage struct {
	sessions map[string]string // K: sessionId, V: email
	mu       sync.RWMutex
	logger   *logger.Logger
}

func NewSessionCustomRepository(logger *logger.Logger) sessionComponentInterfaces.SessionRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewSessionCustomRepository function.")

	sessionsStorage := &sessionsStorage{
		sessions: make(map[string]string),
		mu:       sync.RWMutex{},
		logger:   logger,
	}

	logger.LogrusLogger.Debug("Sessions in storage: " + sessionsStorage.getSessionsInStorageString())
	logger.LogrusLogger.Info("sessionCustomRepository has created.")

	return sessionsStorage
}

func (ss *sessionsStorage) getSessionsInStorageString() string {
	sessionsInStorageString := ""
	for k, v := range ss.sessions {
		sessionsInStorageString += "session_id: " + k + ", for user email: " + v + "; "
	}

	return sessionsInStorageString
}

func (ss *sessionsStorage) CreateSessionForUser(ctx context.Context, email string) (*models.Session, error) {
	ss.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the CreateSessionForUser function.")

	sessionId := sessionUtils.GenerateRandomRunesString(domain.SESSION_ID_STRING_LENGTH)
	ss.sessions[sessionId] = email

	ss.logger.LogrusLoggerWithContext(ctx).Debug("Generated sessionId: ", sessionId, ", for email: ", email, ". Sessions in storage: ", ss.getSessionsInStorageString())
	ss.logger.LogrusLoggerWithContext(ctx).Info("For the email ", email, " created the session.")

	return &models.Session{
		SessionId: sessionId,
	}, nil
}

func (ss *sessionsStorage) GetEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	ss.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetEmailBySession function.")

	email, found := ss.sessions[session.SessionId]
	if !found {
		ss.logger.LogrusLoggerWithContext(ctx).Errorf("Email for the sessionId %s dont exists.", session.SessionId)
		return "", repositoryToUsecaseErrors.SessionRepositoryEmailDontExistsError
	}

	return email, nil
}

func (ss *sessionsStorage) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) {
	ss.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UpdateEmailBySession function.")
	ss.logger.LogrusLoggerWithContext(ctx).Debugf("Before updating email is: %s for the sessionId: %s", ss.sessions[session.SessionId], session.SessionId)

	ss.sessions[session.SessionId] = newEmail

	ss.logger.LogrusLoggerWithContext(ctx).Debugf("After updating email is: %s for the sessionId: %s", ss.sessions[session.SessionId], session.SessionId)
}

func (ss *sessionsStorage) RemoveSession(ctx context.Context, session *models.Session) error {
	ss.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the RemoveSession function.")

	delete(ss.sessions, session.SessionId)

	ss.logger.LogrusLoggerWithContext(ctx).Debug("Removing the session: ", session, ". Sessions in storage: "+ss.getSessionsInStorageString())
	ss.logger.LogrusLoggerWithContext(ctx).Infof("The session %s has been removed", session)

	return nil
}

func (ss *sessionsStorage) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	ss.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the SessionExists function.")

	_, exists := ss.sessions[session.SessionId]

	ss.logger.LogrusLoggerWithContext(ctx).Debugf("The session %v exists: %v", session.SessionId, exists)

	return exists, nil
}
