package delivery

import (
	"2022_2_GoTo_team/internal/authSessionService/domain/customErrors/sessionComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/authSessionService/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/authSessionService/domain/models"
	"2022_2_GoTo_team/pkg/domain/grpcCustomErrors/authSessionServiceErrors"
	"2022_2_GoTo_team/pkg/domain/grpcProtos/authSessionServiceGrpcProtos"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
	"google.golang.org/grpc/status"
)

type SessionDelivery struct {
	authSessionServiceGrpcProtos.UnimplementedAuthSessionServiceServer

	sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface
	logger         *logger.Logger
}

func NewSessionDelivery(sessionUsecase sessionComponentInterfaces.SessionUsecaseInterface, logger *logger.Logger) *SessionDelivery {
	logger.LogrusLogger.Debug("Enter to the NewSessionDelivery function.")

	sessionDelivery := &SessionDelivery{
		sessionUsecase: sessionUsecase,
		logger:         logger,
	}

	logger.LogrusLogger.Info("SessionDelivery has created.")

	return sessionDelivery
}

func (sd *SessionDelivery) SessionExists(ctx context.Context, session *authSessionServiceGrpcProtos.Session) (*authSessionServiceGrpcProtos.Exists, error) {
	sd.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the SessionExists function.")
	sd.logger.LogrusLoggerWithContext(ctx).Debugf("Input session: %#v", session)

	emailExists, err := sd.sessionUsecase.SessionExists(ctx, &models.Session{SessionId: session.SessionId})
	if err != nil {
		switch errors.Unwrap(err).(type) {
		default:
			sd.logger.LogrusLoggerWithContext(ctx).Error(err)
			//return c.NoContent(http.StatusInternalServerError)
			return nil, status.Errorf(500, "")
		}
	}

	exists := authSessionServiceGrpcProtos.Exists{
		Exists: emailExists,
	}
	sd.logger.LogrusLoggerWithContext(ctx).Debug("Formed exists = ", exists.Exists)

	return &exists, nil
}

func (sd *SessionDelivery) CreateSessionForUser(ctx context.Context, userAccountData *authSessionServiceGrpcProtos.UserAccountData) (*authSessionServiceGrpcProtos.Session, error) {
	sd.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the CreateSessionForUser function.")
	sd.logger.LogrusLoggerWithContext(ctx).Debugf("Input userAccountData: %#v", userAccountData)

	session, err := sd.sessionUsecase.CreateSessionForUser(ctx, userAccountData.Email, userAccountData.Password)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailIsNotValidError:
			sd.logger.LogrusLoggerWithContext(ctx).Warn(err)
			//return c.JSON(http.StatusBadRequest, "email is not valid")
			return nil, status.Errorf(400, authSessionServiceErrors.EmailIsNotValidError.Error())
		case *usecaseToDeliveryErrors.PasswordIsNotValidError:
			sd.logger.LogrusLoggerWithContext(ctx).Warn(err)
			//return c.JSON(http.StatusBadRequest, "password is not valid")
			return nil, status.Errorf(400, authSessionServiceErrors.PasswordIsNotValidError.Error())
		case *usecaseToDeliveryErrors.IncorrectEmailOrPasswordError:
			sd.logger.LogrusLoggerWithContext(ctx).Warn(err)
			//return c.JSON(http.StatusBadRequest, "incorrect email or password")
			return nil, status.Errorf(400, authSessionServiceErrors.IncorrectEmailOrPasswordError.Error())
		default:
			sd.logger.LogrusLoggerWithContext(ctx).Error(err)
			//return c.NoContent(http.StatusInternalServerError)
			return nil, status.Errorf(500, "")
		}
	}

	sd.logger.LogrusLoggerWithContext(ctx).Infof("User email %#v auth success!", userAccountData.Email)

	return &authSessionServiceGrpcProtos.Session{
		SessionId: session.SessionId,
	}, nil
}

func (sd *SessionDelivery) RemoveSession(ctx context.Context, session *authSessionServiceGrpcProtos.Session) (*authSessionServiceGrpcProtos.Nothing, error) {
	sd.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the RemoveSession function.")
	sd.logger.LogrusLoggerWithContext(ctx).Debugf("Input session: %#v", session)

	if err := sd.sessionUsecase.RemoveSession(ctx, &models.Session{SessionId: session.SessionId}); err != nil {
		sd.logger.LogrusLoggerWithContext(ctx).Error(err)
		//return c.NoContent(http.StatusInternalServerError)
		return nil, status.Errorf(500, "")
	}

	sd.logger.LogrusLoggerWithContext(ctx).Infof("User session %#v removed successfully.", session.SessionId)

	return &authSessionServiceGrpcProtos.Nothing{Ok: true}, nil
}

func (sd *SessionDelivery) GetUserInfoBySession(ctx context.Context, session *authSessionServiceGrpcProtos.Session) (*authSessionServiceGrpcProtos.UserInfoBySession, error) {
	sd.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetUserInfoBySession function.")
	sd.logger.LogrusLoggerWithContext(ctx).Debugf("Input session: %#v", session)

	user, err := sd.sessionUsecase.GetUserInfoBySession(ctx, &models.Session{SessionId: session.SessionId})
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailForSessionDoesntExistError:
			sd.logger.LogrusLoggerWithContext(ctx).Warn(err)
			//return c.NoContent(http.StatusNotFound)
			return nil, status.Errorf(404, "")
		case *usecaseToDeliveryErrors.UserForSessionDoesntExistError:
			sd.logger.LogrusLoggerWithContext(ctx).Warn(err)
			//return c.NoContent(http.StatusNotFound)
			return nil, status.Errorf(404, "")
		default:
			sd.logger.LogrusLoggerWithContext(ctx).Error(err)
			//return c.NoContent(http.StatusInternalServerError)
			return nil, status.Errorf(500, "")
		}
	}

	userInfoBySession := authSessionServiceGrpcProtos.UserInfoBySession{
		Username:      user.Username,
		AvatarImgPath: user.AvatarImgPath,
	}
	sd.logger.LogrusLoggerWithContext(ctx).Debug("Formed userInfoBySession = ", userInfoBySession.Username, userInfoBySession.AvatarImgPath)

	return &userInfoBySession, nil
}

func (sd *SessionDelivery) GetUserEmailBySession(ctx context.Context, session *authSessionServiceGrpcProtos.Session) (*authSessionServiceGrpcProtos.UserEmail, error) {
	sd.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetUserEmailBySession function.")
	sd.logger.LogrusLoggerWithContext(ctx).Debugf("Input session: %#v", session)

	email, err := sd.sessionUsecase.GetUserEmailBySession(ctx, &models.Session{SessionId: session.SessionId})
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailForSessionDoesntExistError:
			sd.logger.LogrusLoggerWithContext(ctx).Warn(err)
			//return c.NoContent(http.StatusNotFound)
			return nil, status.Errorf(404, "")
		default:
			sd.logger.LogrusLoggerWithContext(ctx).Error(err)
			//return c.NoContent(http.StatusInternalServerError)
			return nil, status.Errorf(500, "")
		}
	}

	userEmail := authSessionServiceGrpcProtos.UserEmail{
		Email: email,
	}
	sd.logger.LogrusLoggerWithContext(ctx).Debug("Formed userEmail = ", userEmail.Email)

	return &userEmail, nil
}

func (sd *SessionDelivery) UpdateEmailBySession(ctx context.Context, updateEmailData *authSessionServiceGrpcProtos.UpdateEmailData) (*authSessionServiceGrpcProtos.Nothing, error) {
	sd.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UpdateEmailBySession function.")
	sd.logger.LogrusLoggerWithContext(ctx).Debugf("Input updateEmailData: %#v, %#v.", updateEmailData.Session.SessionId, updateEmailData.Email)

	if err := sd.sessionUsecase.UpdateEmailBySession(ctx, &models.Session{SessionId: updateEmailData.Session.SessionId}, updateEmailData.Email); err != nil {
		sd.logger.LogrusLoggerWithContext(ctx).Error(err)
		//return c.NoContent(http.StatusInternalServerError)
		return nil, status.Errorf(500, "")
	}

	sd.logger.LogrusLoggerWithContext(ctx).Info("User session updated successfully.")

	return &authSessionServiceGrpcProtos.Nothing{Ok: true}, nil
}
