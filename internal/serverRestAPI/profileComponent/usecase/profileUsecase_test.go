package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"testing"
)

type profileRepositoryMock struct {
}

func (prm *profileRepositoryMock) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	return nil, nil
}

func (prm *profileRepositoryMock) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error {
	return nil
}

func TestSessionRepository(t *testing.T) {
	profileUsecase := NewProfileUsecase(&profileRepositoryMock{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := profileUsecase.GetProfileByEmail(context.Background(), "asd@asd.asd")
	if err != nil {
		t.Error(err)
	}

	err = profileUsecase.UpdateProfileByEmail(context.Background(), &models.Profile{Email: "asd@asd.asd"}, "asd@asd.asd", &models.Session{SessionId: "sess1"})
	if err != nil {
		t.Error(err)
	}
}

type profileRepositoryMock2 struct {
}

func (prm *profileRepositoryMock2) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	return nil, errors.New("err")
}

func (prm *profileRepositoryMock2) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error {
	return errors.New("err")
}

func TestSessionRepository2(t *testing.T) {
	profileUsecase := NewProfileUsecase(&profileRepositoryMock2{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := profileUsecase.GetProfileByEmail(context.Background(), "asd@asd.asd")
	if err == nil {
		t.Error(err)
	}

	err = profileUsecase.UpdateProfileByEmail(context.Background(), &models.Profile{Email: "asd@asd.asd"}, "asd@asd.asd", &models.Session{SessionId: "sess1"})
	if err == nil {
		t.Error(err)
	}
}
