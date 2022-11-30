package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/domain/grpcProtos/userProfileServiceGrpcProtos"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func NewUserProfileServiceTest(grpcConnection userProfileServiceGrpcProtos.UserProfileServiceClient, logger *logger.Logger) *userProfileServiceRepository {
	logger.LogrusLogger.Debug("Enter to the NewAuthSessionServiceRepository function.")

	userProfileServiceRepository := &userProfileServiceRepository{
		userProfileServiceClient: grpcConnection,
		logger:                   logger,
	}

	logger.LogrusLogger.Info("authSessionServiceRepository has created.")

	return userProfileServiceRepository
}

type profileRepositoryMock struct {
}

func (ss *profileRepositoryMock) GetProfileByEmail(ctx context.Context, in *userProfileServiceGrpcProtos.UserEmail, opts ...grpc.CallOption) (*userProfileServiceGrpcProtos.Profile, error) {
	return &userProfileServiceGrpcProtos.Profile{Login: "asd"}, nil
}

func (ss *profileRepositoryMock) UpdateProfileByEmail(ctx context.Context, in *userProfileServiceGrpcProtos.UpdateProfileData, opts ...grpc.CallOption) (*userProfileServiceGrpcProtos.Nothing, error) {
	return &userProfileServiceGrpcProtos.Nothing{}, nil
}

func TestSessionRepository(t *testing.T) {
	var profileServiceClient userProfileServiceGrpcProtos.UserProfileServiceClient = &profileRepositoryMock{}
	profileRepository := NewUserProfileServiceTest(profileServiceClient, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	res, err := profileRepository.GetProfileByEmail(context.Background(), "asd@asd.asd")
	if res == nil || err != nil {
		t.Error(err)
	}

	err = profileRepository.UpdateProfileByEmail(context.Background(), &models.Profile{Email: "ads@ads.asd"},
		"asd@asd.asd", &models.Session{SessionId: "sess"})
	if err != nil {
		t.Error(err)
	}
}
