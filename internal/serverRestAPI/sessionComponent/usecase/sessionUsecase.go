package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/userComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/errorsUtils"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
)

type sessionUsecase struct {
	sessionRepository sessionComponentInterfaces.SessionRepositoryInterface
	userRepository    userComponentInterfaces.UserRepositoryInterface
	logger            *logger.Logger
}

func NewSessionUsecase(sessionRepository sessionComponentInterfaces.SessionRepositoryInterface, userRepository userComponentInterfaces.UserRepositoryInterface, logger *logger.Logger) sessionComponentInterfaces.SessionUsecaseInterface {
	return &sessionUsecase{
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
		logger:            logger,
	}
}

func (su *sessionUsecase) SessionExists(session *models.Session) (bool, error) {
	wrappingErrorMessage := "error while checking session exists:"

	exists, err := su.sessionRepository.SessionExists(session)
	if err != nil {
		return false, errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	return exists, nil
}

func (su *sessionUsecase) CreateSessionForUser(email string, password string) (*models.Session, error) {
	wrappingErrorMessage := "error while creating session for user"

	user, err := su.userRepository.GetUserByEmail(email)
	if err != nil {
		// TODO logger
		return nil, errorsUtils.WrapError(wrappingErrorMessage, err)
	}
	if user.Password != password {
		// TODO logger
		//o.logger.Error("Invalid password")
		return nil, errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	session, err := su.sessionRepository.CreateSessionForUser(email)
	if err != nil {
		return nil, errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	return session, nil
}

func (su *sessionUsecase) RemoveSession(session *models.Session) error {
	wrappingErrorMessage := "error while removing session"

	if err := su.sessionRepository.RemoveSession(session); err != nil {
		return errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	return nil
}

func (su *sessionUsecase) GetUserBySession(session *models.Session) (*models.User, error) {
	wrappingErrorMessage := "error while getting user by session"

	email, err := su.sessionRepository.GetEmailBySession(session)
	if err != nil {
		return nil, errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	user, err := su.userRepository.GetUserByEmail(email)
	if err != nil {
		// TODO logger
		//api.logger.Error(err.Error())
		_ = su.RemoveSession(session) // We should try to remove "garbage" session
		return nil, errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	return user, nil
}
