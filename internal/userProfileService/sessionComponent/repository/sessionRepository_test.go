package repository

import (
	"2022_2_GoTo_team/internal/userProfileService/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/userProfileService/domain/models"
	"2022_2_GoTo_team/pkg/domain/grpcProtos/authSessionServiceGrpcProtos"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func NewAuthSessionServiceTest(grpcConnection authSessionServiceGrpcProtos.AuthSessionServiceClient, logger *logger.Logger) sessionComponentInterfaces.SessionRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewAuthSessionServiceRepository function.")

	authSessionServiceRepository := &authSessionServiceRepository{
		authSessionServiceClient: grpcConnection,
		logger:                   logger,
	}

	logger.LogrusLogger.Info("authSessionServiceRepository has created.")

	return authSessionServiceRepository
}

type sessionRepositoryMock struct {
}

func (ss *sessionRepositoryMock) SessionExists(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Exists, error) {
	return &authSessionServiceGrpcProtos.Exists{}, nil
}

func (ss *sessionRepositoryMock) CreateSessionForUser(ctx context.Context, in *authSessionServiceGrpcProtos.UserAccountData, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Session, error) {
	return &authSessionServiceGrpcProtos.Session{}, nil
}

func (ss *sessionRepositoryMock) RemoveSession(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Nothing, error) {
	return &authSessionServiceGrpcProtos.Nothing{}, nil
}

func (ss *sessionRepositoryMock) GetUserInfoBySession(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.UserInfoBySession, error) {
	return &authSessionServiceGrpcProtos.UserInfoBySession{}, nil
}

func (ss *sessionRepositoryMock) GetUserEmailBySession(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.UserEmail, error) {
	return &authSessionServiceGrpcProtos.UserEmail{}, nil
}

func (ss *sessionRepositoryMock) UpdateEmailBySession(ctx context.Context, in *authSessionServiceGrpcProtos.UpdateEmailData, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Nothing, error) {
	return &authSessionServiceGrpcProtos.Nothing{}, nil
}

func TestSessionRepository(t *testing.T) {
	var sessionServiceClient authSessionServiceGrpcProtos.AuthSessionServiceClient = &sessionRepositoryMock{}
	sessionRepository := NewAuthSessionServiceTest(sessionServiceClient, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	err := sessionRepository.UpdateEmailBySession(context.Background(), &models.Session{
		SessionId: "sessId",
	}, "asd@asd.asd")
	if err != nil {
		t.Error(err)
	}

	res := NewAuthSessionServiceRepository(&grpc.ClientConn{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})
	if res == nil {
		t.Error(err)
	}
}

type sessionRepositoryMockNegative struct {
}

func (ss *sessionRepositoryMockNegative) SessionExists(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Exists, error) {
	return &authSessionServiceGrpcProtos.Exists{}, nil
}

func (ss *sessionRepositoryMockNegative) CreateSessionForUser(ctx context.Context, in *authSessionServiceGrpcProtos.UserAccountData, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Session, error) {
	return &authSessionServiceGrpcProtos.Session{}, nil
}

func (ss *sessionRepositoryMockNegative) RemoveSession(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Nothing, error) {
	return &authSessionServiceGrpcProtos.Nothing{}, nil
}

func (ss *sessionRepositoryMockNegative) GetUserInfoBySession(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.UserInfoBySession, error) {
	return &authSessionServiceGrpcProtos.UserInfoBySession{}, nil
}

func (ss *sessionRepositoryMockNegative) GetUserEmailBySession(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.UserEmail, error) {
	return &authSessionServiceGrpcProtos.UserEmail{}, nil
}

func (ss *sessionRepositoryMockNegative) UpdateEmailBySession(ctx context.Context, in *authSessionServiceGrpcProtos.UpdateEmailData, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Nothing, error) {
	return nil, errors.New("err")
}

func TestSessionRepositoryNegative(t *testing.T) {
	var sessionServiceClient authSessionServiceGrpcProtos.AuthSessionServiceClient = &sessionRepositoryMockNegative{}
	sessionRepository := NewAuthSessionServiceTest(sessionServiceClient, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	err := sessionRepository.UpdateEmailBySession(context.Background(), &models.Session{
		SessionId: "sessId",
	}, "asd@asd.asd")
	if err == nil {
		t.Error(err)
	}
}
