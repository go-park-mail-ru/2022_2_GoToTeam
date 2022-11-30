package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
)

type sessionUsecase struct {
	sessionRepository sessionComponentInterfaces.SessionRepositoryInterface
	logger            *logger.Logger
}

func NewSessionUsecase(sessionRepository sessionComponentInterfaces.SessionRepositoryInterface, logger *logger.Logger) sessionComponentInterfaces.SessionUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewSessionUsecase function.")

	sessionUsecase := &sessionUsecase{
		sessionRepository: sessionRepository,
		logger:            logger,
	}

	logger.LogrusLogger.Info("sessionUsecase has created.")

	return sessionUsecase
}

func (su *sessionUsecase) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the SessionExists function.")

	exists, err := su.sessionRepository.SessionExists(ctx, session)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}

	return exists, err
}

func (su *sessionUsecase) CreateSessionForUser(ctx context.Context, email string, password string) (*models.Session, error) {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the CreateSessionForUser function.")

	session, err := su.sessionRepository.CreateSessionForUser(ctx, email, password)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}

	return session, err
}

func (su *sessionUsecase) RemoveSession(ctx context.Context, session *models.Session) error {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the RemoveSession function.")

	err := su.sessionRepository.RemoveSession(ctx, session)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}

	return err
}

func (su *sessionUsecase) GetUserInfoBySession(ctx context.Context, session *models.Session) (*models.User, error) {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetUserInfoByEmail function.")

	user, err := su.sessionRepository.GetUserInfoBySession(ctx, session)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}

	return user, err
}

func (su *sessionUsecase) GetUserEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetUserEmailBySession function.")

	email, err := su.sessionRepository.GetUserEmailBySession(ctx, session)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}

	return email, err
}
