package usecase

import (
	repositoryToUsecaseErrors2 "2022_2_GoTo_team/internal/authSessionService/domain/customErrors/sessionComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/authSessionService/domain/customErrors/sessionComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/authSessionService/domain/customErrors/userComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/authSessionService/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/authSessionService/domain/interfaces/userComponentInterfaces"
	"2022_2_GoTo_team/internal/authSessionService/domain/models"
	"2022_2_GoTo_team/pkg/errorsUtils"
	"2022_2_GoTo_team/pkg/logger"
	"2022_2_GoTo_team/pkg/validators"
	"context"
	"errors"
)

type sessionUsecase struct {
	sessionRepository sessionComponentInterfaces.SessionRepositoryInterface
	userRepository    userComponentInterfaces.UserRepositoryInterface
	logger            *logger.Logger
}

func NewSessionUsecase(sessionRepository sessionComponentInterfaces.SessionRepositoryInterface, userRepository userComponentInterfaces.UserRepositoryInterface, logger *logger.Logger) sessionComponentInterfaces.SessionUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewSessionUsecase function.")

	sessionUsecase := &sessionUsecase{
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
		logger:            logger,
	}

	logger.LogrusLogger.Info("sessionUsecase has created.")

	return sessionUsecase
}

func (su *sessionUsecase) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the SessionExists function.")

	wrappingErrorMessage := "error while checking session exists"

	exists, err := su.sessionRepository.SessionExists(ctx, session)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Error(err)
		return false, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return exists, nil
}

func (su *sessionUsecase) CreateSessionForUser(ctx context.Context, email string, password string) (*models.Session, error) {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the CreateSessionForUser function.")

	wrappingErrorMessage := "error while creating session for user"

	if err := su.validateUserData(ctx, email, password); err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	exists, err := su.userRepository.CheckUserEmailAndPassword(ctx, email, password)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}
	if !exists {
		su.logger.LogrusLoggerWithContext(ctx).Warn("Incorrect email or password.")
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.IncorrectEmailOrPasswordError{Err: errors.New("incorrect email or password")})
	}

	session, err := su.sessionRepository.CreateSessionForUser(ctx, email)
	if err != nil {
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return session, nil
}

func (su *sessionUsecase) RemoveSession(ctx context.Context, session *models.Session) error {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the RemoveSession function.")

	wrappingErrorMessage := "error while removing session"

	if err := su.sessionRepository.RemoveSession(ctx, session); err != nil {
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return nil
}

func (su *sessionUsecase) GetUserInfoBySession(ctx context.Context, session *models.Session) (*models.User, error) {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetUserInfoByEmail function.")

	wrappingErrorMessage := "error while getting user info by session"

	email, err := su.sessionRepository.GetEmailBySession(ctx, session)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Error(err)
		switch err {
		case repositoryToUsecaseErrors2.SessionRepositoryEmailDoesntExistError:
			su.logger.LogrusLoggerWithContext(ctx).Debug("Trying to remove the garbage session: %#v", session)
			_ = su.RemoveSession(ctx, session) // We should try to remove "garbage" session
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{Err: err})
		default:
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	user, err := su.userRepository.GetUserInfoForSessionComponentByEmail(ctx, email)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Error(err)
		switch err {
		case repositoryToUsecaseErrors.UserRepositoryEmailDoesntExistError:
			su.logger.LogrusLoggerWithContext(ctx).Debug("Trying to remove the garbage session: %#v", session)
			_ = su.RemoveSession(ctx, session) // We should try to remove "garbage" session
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.UserForSessionDoesntExistError{Err: err})
		default:
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	return user, nil
}

func (su *sessionUsecase) GetUserEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetUserEmailBySession function.")

	wrappingErrorMessage := "error while getting email by session"

	email, err := su.sessionRepository.GetEmailBySession(ctx, session)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Error(err)
		switch err {
		case repositoryToUsecaseErrors2.SessionRepositoryEmailDoesntExistError:
			su.logger.LogrusLoggerWithContext(ctx).Debug("Trying to remove the garbage session: %#v", session)
			_ = su.RemoveSession(ctx, session) // We should try to remove "garbage" session
			return "", errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{Err: err})
		default:
			return "", errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	return email, nil
}

func (su *sessionUsecase) validateUserData(ctx context.Context, email string, password string) error {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the validateUserData function.")

	if !validators.EmailIsValidByCustomValidation(email) {
		su.logger.LogrusLoggerWithContext(ctx).Debugf("Email %s is not valid.", email)
		return &usecaseToDeliveryErrors.EmailIsNotValidError{Err: errors.New("email is not valid")}
	}
	if !validators.PasswordIsValidByRegExp(password) {
		su.logger.LogrusLoggerWithContext(ctx).Debug("Password is not valid.")
		return &usecaseToDeliveryErrors.PasswordIsNotValidError{Err: errors.New("password is not valid")}
	}

	return nil
}
