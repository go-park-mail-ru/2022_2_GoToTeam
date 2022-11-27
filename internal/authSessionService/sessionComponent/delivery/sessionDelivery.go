package delivery

import (
	"2022_2_GoTo_team/internal/authSessionService/domain/customErrors/sessionComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/authSessionService/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/authSessionService/domain/models"
	"2022_2_GoTo_team/pkg/domain/grpcCustomErrors/authSessionServiceErrors"
	"2022_2_GoTo_team/pkg/domain/grpcProtos/authSessionService"
	"2022_2_GoTo_team/pkg/logger"
	"context"
	"errors"
	"google.golang.org/grpc/status"
)

type SessionDelivery struct {
	authSessionService.UnimplementedAuthSessionServiceServer

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

func (sd *SessionDelivery) CreateSessionForUser(ctx context.Context, userAccountData *authSessionService.UserAccountData) (*authSessionService.Session, error) {
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

	return &authSessionService.Session{
		SessionId: session.SessionId,
	}, nil
}

func (sd *SessionDelivery) RemoveSession(ctx context.Context, session *authSessionService.Session) (*authSessionService.Nothing, error) {
	sd.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the RemoveSession function.")
	sd.logger.LogrusLoggerWithContext(ctx).Debugf("Input session: %#v", session)

	if err := sd.sessionUsecase.RemoveSession(ctx, &models.Session{SessionId: session.SessionId}); err != nil {
		sd.logger.LogrusLoggerWithContext(ctx).Error(err)
		//return c.NoContent(http.StatusInternalServerError)
		return nil, status.Errorf(500, "")
	}

	sd.logger.LogrusLoggerWithContext(ctx).Infof("User session %#v removed successfully.", session.SessionId)

	return &authSessionService.Nothing{Ok: true}, nil
}

func (sd *SessionDelivery) GetUserInfoBySession(ctx context.Context, session *authSessionService.Session) (*authSessionService.UserInfoBySession, error) {
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

	userInfoBySession := authSessionService.UserInfoBySession{
		Username:      user.Username,
		AvatarImgPath: user.AvatarImgPath,
	}
	sd.logger.LogrusLoggerWithContext(ctx).Debug("Formed userInfoBySession = ", userInfoBySession.Username, userInfoBySession.AvatarImgPath)

	return &userInfoBySession, nil
}
