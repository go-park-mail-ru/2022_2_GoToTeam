package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"testing"
)

type sessionRepositoryMock struct {
}

func (srm *sessionRepositoryMock) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	return true, nil
}

func (prm *sessionRepositoryMock) CreateSessionForUser(ctx context.Context, email string, password string) (*models.Session, error) {
	return &models.Session{SessionId: "sess1"}, nil
}

func (prm *sessionRepositoryMock) RemoveSession(ctx context.Context, session *models.Session) error {
	return nil
}

func (prm *sessionRepositoryMock) GetUserInfoBySession(ctx context.Context, session *models.Session) (*models.User, error) {
	return &models.User{Email: "asd@ads.asd"}, nil
}

func (prm *sessionRepositoryMock) GetUserEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	return "", nil
}

func (prm *sessionRepositoryMock) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return nil
}

func TestSessionRepository(t *testing.T) {
	sessionUsecase := NewSessionUsecase(&sessionRepositoryMock{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := sessionUsecase.SessionExists(context.Background(), &models.Session{SessionId: "sess1"})
	if err != nil {
		t.Error(err)
	}

	_, err = sessionUsecase.CreateSessionForUser(context.Background(), "asd@asd.asd", "pswd")
	if err != nil {
		t.Error(err)
	}

	err = sessionUsecase.RemoveSession(context.Background(), &models.Session{SessionId: "sess1"})
	if err != nil {
		t.Error(err)
	}

	_, err = sessionUsecase.GetUserInfoBySession(context.Background(), &models.Session{SessionId: "sess1"})
	if err != nil {
		t.Error(err)
	}

	_, err = sessionUsecase.GetUserEmailBySession(context.Background(), &models.Session{SessionId: "sess1"})
	if err != nil {
		t.Error(err)
	}
}

type sessionRepositoryMock2 struct {
}

func (srm *sessionRepositoryMock2) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	return true, errors.New("err")
}

func (prm *sessionRepositoryMock2) CreateSessionForUser(ctx context.Context, email string, password string) (*models.Session, error) {
	return &models.Session{SessionId: "sess1"}, errors.New("err")
}

func (prm *sessionRepositoryMock2) RemoveSession(ctx context.Context, session *models.Session) error {
	return errors.New("err")
}

func (prm *sessionRepositoryMock2) GetUserInfoBySession(ctx context.Context, session *models.Session) (*models.User, error) {
	return &models.User{Email: "asd@ads.asd"}, errors.New("err")
}

func (prm *sessionRepositoryMock2) GetUserEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	return "", errors.New("err")
}

func (prm *sessionRepositoryMock2) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return errors.New("err")
}

func TestSessionRepositoryNegative(t *testing.T) {
	sessionUsecase := NewSessionUsecase(&sessionRepositoryMock2{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := sessionUsecase.SessionExists(context.Background(), &models.Session{SessionId: "sess1"})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionUsecase.CreateSessionForUser(context.Background(), "asd@asd.asd", "pswd")
	if err == nil {
		t.Error(err)
	}

	err = sessionUsecase.RemoveSession(context.Background(), &models.Session{SessionId: "sess1"})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionUsecase.GetUserInfoBySession(context.Background(), &models.Session{SessionId: "sess1"})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionUsecase.GetUserEmailBySession(context.Background(), &models.Session{SessionId: "sess1"})
	if err == nil {
		t.Error(err)
	}
}
