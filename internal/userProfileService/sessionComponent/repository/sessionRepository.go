package repository

import (
	"2022_2_GoTo_team/internal/userProfileService/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/userProfileService/domain/models"
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
