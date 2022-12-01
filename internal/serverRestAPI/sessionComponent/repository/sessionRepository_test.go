package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
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
	return &authSessionServiceGrpcProtos.Exists{Exists: true}, nil
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
	return &authSessionServiceGrpcProtos.UserEmail{Email: "asd@asd.asd"}, nil
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

	res, err := sessionRepository.SessionExists(context.Background(), &models.Session{
		SessionId: "sessId",
	})
	if res != true || err != nil {
		t.Error(err)
	}

	res2, err := sessionRepository.CreateSessionForUser(context.Background(), "asd@asd.asd", "qwerty123")
	if res2 == nil || err != nil {
		t.Error(err)
	}

	err = sessionRepository.RemoveSession(context.Background(), &models.Session{
		SessionId: "sessId",
	})
	if err != nil {
		t.Error(err)
	}

	res3, err := sessionRepository.GetUserInfoBySession(context.Background(), &models.Session{
		SessionId: "sessId",
	})
	if res3 == nil || err != nil {
		t.Error(err)
	}

	res4, err := sessionRepository.GetUserEmailBySession(context.Background(), &models.Session{
		SessionId: "sessId",
	})
	if res4 == "" || err != nil {
		t.Error(err)
	}

	err = sessionRepository.UpdateEmailBySession(context.Background(), &models.Session{
		SessionId: "sessId",
	}, "asd@asd.asd")
	if err != nil {
		t.Error(err)
	}

	resk := NewAuthSessionServiceRepository(&grpc.ClientConn{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})
	if resk == nil {
		t.Error(err)
	}
}

type sessionRepositoryMockNegative struct {
}

func (ss *sessionRepositoryMockNegative) SessionExists(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Exists, error) {
	return &authSessionServiceGrpcProtos.Exists{Exists: true}, errors.New("err")
}

func (ss *sessionRepositoryMockNegative) CreateSessionForUser(ctx context.Context, in *authSessionServiceGrpcProtos.UserAccountData, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Session, error) {
	return &authSessionServiceGrpcProtos.Session{}, errors.New("err")
}

func (ss *sessionRepositoryMockNegative) RemoveSession(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Nothing, error) {
	return &authSessionServiceGrpcProtos.Nothing{}, errors.New("err")
}

func (ss *sessionRepositoryMockNegative) GetUserInfoBySession(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.UserInfoBySession, error) {
	return &authSessionServiceGrpcProtos.UserInfoBySession{}, errors.New("err")
}

func (ss *sessionRepositoryMockNegative) GetUserEmailBySession(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.UserEmail, error) {
	return &authSessionServiceGrpcProtos.UserEmail{}, errors.New("err")
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

	res, err := sessionRepository.SessionExists(context.Background(), &models.Session{
		SessionId: "sessId",
	})
	if res != true || err == nil {
		t.Error(err)
	}

	res2, err := sessionRepository.CreateSessionForUser(context.Background(), "asd@asd.asd", "qwerty123")
	if res2 == nil || err == nil {
		t.Error(err)
	}

	err = sessionRepository.RemoveSession(context.Background(), &models.Session{
		SessionId: "sessId",
	})
	if err == nil {
		t.Error(err)
	}

	res3, err := sessionRepository.GetUserInfoBySession(context.Background(), &models.Session{
		SessionId: "sessId",
	})
	if res3 == nil || err == nil {
		t.Error(err)
	}

	_, err = sessionRepository.GetUserEmailBySession(context.Background(), &models.Session{
		SessionId: "sessId",
	})
	if err == nil {
		t.Error(err)
	}

	err = sessionRepository.UpdateEmailBySession(context.Background(), &models.Session{
		SessionId: "sessId",
	}, "asd@asd.asd")
	if err == nil {
		t.Error(err)
	}
}

type sessionRepositoryMockNegative2 struct {
}

func (ss *sessionRepositoryMockNegative2) SessionExists(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Exists, error) {
	return nil, errors.New("err")
}

func (ss *sessionRepositoryMockNegative2) CreateSessionForUser(ctx context.Context, in *authSessionServiceGrpcProtos.UserAccountData, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Session, error) {
	return nil, errors.New("err")
}

func (ss *sessionRepositoryMockNegative2) RemoveSession(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Nothing, error) {
	return &authSessionServiceGrpcProtos.Nothing{}, errors.New("err")
}

func (ss *sessionRepositoryMockNegative2) GetUserInfoBySession(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.UserInfoBySession, error) {
	return nil, errors.New("err")
}

func (ss *sessionRepositoryMockNegative2) GetUserEmailBySession(ctx context.Context, in *authSessionServiceGrpcProtos.Session, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.UserEmail, error) {
	return nil, errors.New("err")
}

func (ss *sessionRepositoryMockNegative2) UpdateEmailBySession(ctx context.Context, in *authSessionServiceGrpcProtos.UpdateEmailData, opts ...grpc.CallOption) (*authSessionServiceGrpcProtos.Nothing, error) {
	return nil, errors.New("err")
}

func TestSessionRepositoryNegative2(t *testing.T) {
	var sessionServiceClient authSessionServiceGrpcProtos.AuthSessionServiceClient = &sessionRepositoryMockNegative2{}
	sessionRepository := NewAuthSessionServiceTest(sessionServiceClient, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	res, err := sessionRepository.SessionExists(context.Background(), &models.Session{
		SessionId: "sessId",
	})
	if res != false {
		t.Error(err)
	}

	res2, err := sessionRepository.CreateSessionForUser(context.Background(), "asd@asd.asd", "qwerty123")
	if res2 != nil {
		t.Error(err)
	}

	err = sessionRepository.RemoveSession(context.Background(), &models.Session{
		SessionId: "sessId",
	})
	if err == nil {
		t.Error(err)
	}

	res3, err := sessionRepository.GetUserInfoBySession(context.Background(), &models.Session{
		SessionId: "sessId",
	})
	if res3 != nil {
		t.Error(err)
	}

	res4, err := sessionRepository.GetUserEmailBySession(context.Background(), &models.Session{
		SessionId: "sessId",
	})
	if res4 != "" {
		t.Error(err)
	}

	err = sessionRepository.UpdateEmailBySession(context.Background(), &models.Session{
		SessionId: "sessId",
	}, "asd@asd.asd")
	if err == nil {
		t.Error(err)
	}
}
