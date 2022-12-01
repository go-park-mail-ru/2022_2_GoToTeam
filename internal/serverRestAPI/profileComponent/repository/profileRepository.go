package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/profileComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/domain/grpcProtos/userProfileServiceGrpcProtos"
	"2022_2_GoTo_team/pkg/utils/grpcUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"google.golang.org/grpc"
)

type userProfileServiceRepository struct {
	userProfileServiceClient userProfileServiceGrpcProtos.UserProfileServiceClient
	logger                   *logger.Logger
}

func NewUserProfileServiceRepository(grpcConnection *grpc.ClientConn, logger *logger.Logger) profileComponentInterfaces.ProfileRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewUserProfileServiceRepository function.")

	userProfileServiceRepository := &userProfileServiceRepository{
		userProfileServiceClient: userProfileServiceGrpcProtos.NewUserProfileServiceClient(grpcConnection),
		logger:                   logger,
	}

	logger.LogrusLogger.Info("userProfileServiceRepository has created.")

	return userProfileServiceRepository
}

func (upsr *userProfileServiceRepository) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetProfileByEmail function.")

	profile, err := upsr.userProfileServiceClient.GetProfileByEmail(grpcUtils.UpgradeContextByInjectedMetadata(ctx), &userProfileServiceGrpcProtos.UserEmail{
		Email: email,
	})
	if err != nil {
		upsr.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}
	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Got profile: ", profile)
	if profile == nil {
		return nil, err
	}

	return &models.Profile{
		Email:         profile.Email,
		Login:         profile.Login,
		Password:      profile.Password,
		Username:      profile.Username,
		AvatarImgPath: profile.AvatarImgPath,
	}, err
}

func (upsr *userProfileServiceRepository) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error {
	upsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UpdateProfileByEmail function.")

	_, err := upsr.userProfileServiceClient.UpdateProfileByEmail(grpcUtils.UpgradeContextByInjectedMetadata(ctx), &userProfileServiceGrpcProtos.UpdateProfileData{
		Profile: &userProfileServiceGrpcProtos.Profile{
			Email:         newProfile.Email,
			Login:         newProfile.Login,
			Password:      newProfile.Password,
			Username:      newProfile.Username,
			AvatarImgPath: newProfile.AvatarImgPath,
		},
		Email:     email,
		SessionId: session.SessionId,
	})
	if err != nil {
		upsr.logger.LogrusLoggerWithContext(ctx).Warn(err)
	}

	return err
}
