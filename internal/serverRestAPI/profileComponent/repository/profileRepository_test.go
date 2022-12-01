package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/domain/grpcProtos/userProfileServiceGrpcProtos"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
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

	resk := NewUserProfileServiceRepository(&grpc.ClientConn{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})
	if resk == nil {
		t.Error("nil")
	}
}

type profileRepositoryMock2 struct {
}

func (ss *profileRepositoryMock2) GetProfileByEmail(ctx context.Context, in *userProfileServiceGrpcProtos.UserEmail, opts ...grpc.CallOption) (*userProfileServiceGrpcProtos.Profile, error) {
	return &userProfileServiceGrpcProtos.Profile{Login: "asd"}, errors.New("err")
}

func (ss *profileRepositoryMock2) UpdateProfileByEmail(ctx context.Context, in *userProfileServiceGrpcProtos.UpdateProfileData, opts ...grpc.CallOption) (*userProfileServiceGrpcProtos.Nothing, error) {
	return &userProfileServiceGrpcProtos.Nothing{}, errors.New("err")
}

func TestSessionRepositoryNegative(t *testing.T) {
	var profileServiceClient userProfileServiceGrpcProtos.UserProfileServiceClient = &profileRepositoryMock2{}
	profileRepository := NewUserProfileServiceTest(profileServiceClient, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	res, err := profileRepository.GetProfileByEmail(context.Background(), "asd@asd.asd")
	if res == nil || err == nil {
		t.Error(err)
	}

	err = profileRepository.UpdateProfileByEmail(context.Background(), &models.Profile{Email: "ads@ads.asd"},
		"asd@asd.asd", &models.Session{SessionId: "sess"})
	if err == nil {
		t.Error(err)
	}
}

type profileRepositoryMock3 struct {
}

func (ss *profileRepositoryMock3) GetProfileByEmail(ctx context.Context, in *userProfileServiceGrpcProtos.UserEmail, opts ...grpc.CallOption) (*userProfileServiceGrpcProtos.Profile, error) {
	return nil, nil
}

func (ss *profileRepositoryMock3) UpdateProfileByEmail(ctx context.Context, in *userProfileServiceGrpcProtos.UpdateProfileData, opts ...grpc.CallOption) (*userProfileServiceGrpcProtos.Nothing, error) {
	return &userProfileServiceGrpcProtos.Nothing{}, errors.New("err")
}

func TestSessionRepositoryNegative2(t *testing.T) {
	var profileServiceClient userProfileServiceGrpcProtos.UserProfileServiceClient = &profileRepositoryMock3{}
	profileRepository := NewUserProfileServiceTest(profileServiceClient, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	res, err := profileRepository.GetProfileByEmail(context.Background(), "asd@asd.asd")
	if res != nil {
		t.Error(err)
	}

}
