package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/domain/grpcProtos/authSessionServiceGrpcProtos"
	"2022_2_GoTo_team/pkg/utils/grpcUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"google.golang.org/grpc"
)

type authSessionServiceRepository struct {
	authSessionServiceClient authSessionServiceGrpcProtos.AuthSessionServiceClient
	logger                   *logger.Logger
}

func NewAuthSessionServiceRepository(grpcConnection *grpc.ClientConn, logger *logger.Logger) sessionComponentInterfaces.SessionRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewAuthSessionServiceRepository function.")

	authSessionServiceRepository := &authSessionServiceRepository{
		authSessionServiceClient: authSessionServiceGrpcProtos.NewAuthSessionServiceClient(grpcConnection),
		logger:                   logger,
	}

	logger.LogrusLogger.Info("authSessionServiceRepository has created.")

	return authSessionServiceRepository
}

func (assr *authSessionServiceRepository) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	assr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the SessionExists function.")

	exists, err := assr.authSessionServiceClient.SessionExists(grpcUtils.MakeNewContextWithGrpcMetadataBasedOnContext(ctx), &authSessionServiceGrpcProtos.Session{
		SessionId: session.SessionId,
	})
	if err != nil {
		assr.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}
	assr.logger.LogrusLoggerWithContext(ctx).Debug("Got exists: ", exists)
	if exists == nil {
		return false, err
	}

	return exists.Exists, err
}

func (assr *authSessionServiceRepository) CreateSessionForUser(ctx context.Context, email string, password string) (*models.Session, error) {
	assr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the CreateSessionForUser function.")

	session, err := assr.authSessionServiceClient.CreateSessionForUser(grpcUtils.MakeNewContextWithGrpcMetadataBasedOnContext(ctx), &authSessionServiceGrpcProtos.UserAccountData{
		Email:    email,
		Password: password,
	})
	if err != nil {
		assr.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}
	assr.logger.LogrusLoggerWithContext(ctx).Debug("Got session: ", session)
	if session == nil {
		return nil, err
	}

	return &models.Session{SessionId: session.SessionId}, err
}

func (assr *authSessionServiceRepository) RemoveSession(ctx context.Context, session *models.Session) error {
	assr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the RemoveSession function.")

	_, err := assr.authSessionServiceClient.RemoveSession(grpcUtils.MakeNewContextWithGrpcMetadataBasedOnContext(ctx), &authSessionServiceGrpcProtos.Session{
		SessionId: session.SessionId,
	})
	if err != nil {
		assr.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}

	return err
}

func (assr *authSessionServiceRepository) GetUserInfoBySession(ctx context.Context, session *models.Session) (*models.User, error) {
	assr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetUserInfoBySession function.")

	userInfo, err := assr.authSessionServiceClient.GetUserInfoBySession(grpcUtils.MakeNewContextWithGrpcMetadataBasedOnContext(ctx), &authSessionServiceGrpcProtos.Session{
		SessionId: session.SessionId,
	})
	if err != nil {
		assr.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}
	assr.logger.LogrusLoggerWithContext(ctx).Debug("Got userInfo: ", userInfo)
	if userInfo == nil {
		return nil, err
	}

	return &models.User{Username: userInfo.Username, AvatarImgPath: userInfo.AvatarImgPath}, err
}

func (assr *authSessionServiceRepository) GetUserEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	assr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetUserEmailBySession function.")

	userEmail, err := assr.authSessionServiceClient.GetUserEmailBySession(grpcUtils.MakeNewContextWithGrpcMetadataBasedOnContext(ctx), &authSessionServiceGrpcProtos.Session{
		SessionId: session.SessionId,
	})
	if err != nil {
		assr.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}
	assr.logger.LogrusLoggerWithContext(ctx).Debug("Got userEmail: ", userEmail)
	if userEmail == nil {
		return "", err
	}

	return userEmail.Email, err
}

func (assr *authSessionServiceRepository) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	assr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UpdateEmailBySession function.")

	_, err := assr.authSessionServiceClient.UpdateEmailBySession(grpcUtils.MakeNewContextWithGrpcMetadataBasedOnContext(ctx), &authSessionServiceGrpcProtos.UpdateEmailData{
		Session: &authSessionServiceGrpcProtos.Session{SessionId: session.SessionId},
		Email:   newEmail,
	})
	if err != nil {
		assr.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}

	return err
}
