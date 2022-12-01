package usecase

import (
	"2022_2_GoTo_team/internal/userProfileService/domain/customErrors/profileComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/userProfileService/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

var loggerMock = &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
	"requestId": "qwerty",
	"userEmail": "asd@asd.asd",
})}

type profileRepositoryMock struct {
}

func (prm *profileRepositoryMock) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	return &models.Profile{Email: "asd@asd.asd"}, nil
}

func (prm *profileRepositoryMock) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string) error {
	return nil
}

type sessionRepositoryMock struct {
}

func (srp *sessionRepositoryMock) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return errors.New("err")
}

var pu = NewProfileUsecase(&profileRepositoryMock{}, &sessionRepositoryMock{}, loggerMock)

func TestGetProfileByEmail(t *testing.T) {
	res, err := pu.GetProfileByEmail(context.Background(), "asd@asd.asd")
	assert.Equal(t, &models.Profile{Email: "asd@asd.asd"}, res)
	assert.Equal(t, nil, err)
}

func TestUpdateProfileByEmaill(t *testing.T) {
	err := pu.UpdateProfileByEmail(context.Background(), &models.Profile{
		Email:    "asd@asd.asd",
		Login:    "Abracodabra",
		Password: "Qwerty123",
	}, "asd@asd.asd", &models.Session{
		SessionId: "sess",
	})
	assert.Equal(t, nil, err)
}

func TestUpdateProfileByEmailShouldUpdateSession(t *testing.T) {
	err := pu.UpdateProfileByEmail(context.Background(), &models.Profile{
		Email:    "asd@asd.asd",
		Login:    "Abracodabra",
		Password: "Qwerty123",
	}, "newEmail@asd.asd", &models.Session{
		SessionId: "sess",
	})
	assert.Equal(t, nil, err)
}

func TestUpdateProfileByEmaillNegativeValidationEmail(t *testing.T) {
	err := pu.UpdateProfileByEmail(context.Background(), &models.Profile{
		Email:    "a",
		Login:    "Abracodabra",
		Password: "Qwerty123",
	}, "asd@asd.asd", &models.Session{
		SessionId: "sess",
	})
	assert.NotEqual(t, nil, err)
}

func TestUpdateProfileByEmaillNegativeValidationLogin(t *testing.T) {
	err := pu.UpdateProfileByEmail(context.Background(), &models.Profile{
		Email:    "asd@asd.asd",
		Login:    "123",
		Password: "Qwerty123",
	}, "asd@asd.asd", &models.Session{
		SessionId: "sess",
	})
	assert.NotEqual(t, nil, err)
}

func TestUpdateProfileByEmaillNegativeValidationPassword(t *testing.T) {
	err := pu.UpdateProfileByEmail(context.Background(), &models.Profile{
		Email:    "asd@asd.asd",
		Login:    "Abracodabra",
		Password: "1",
	}, "asd@asd.asd", &models.Session{
		SessionId: "sess",
	})
	assert.NotEqual(t, nil, err)
}

type profileRepositoryMock2 struct {
}

func (prm *profileRepositoryMock2) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	return nil, errors.New("err")
}

func (prm *profileRepositoryMock2) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string) error {
	return errors.New("err")
}

type sessionRepositoryMock2 struct {
}

func (srp *sessionRepositoryMock2) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return nil
}

var pu2 = NewProfileUsecase(&profileRepositoryMock2{}, &sessionRepositoryMock2{}, loggerMock)

func TestUpdateProfileByEmailNegativeUnknownError(t *testing.T) {
	err := pu2.UpdateProfileByEmail(context.Background(), &models.Profile{
		Email:    "asd@asd.asd",
		Login:    "Abracodabra",
		Password: "Qwerty123",
	}, "newEmail@asd.asd", &models.Session{
		SessionId: "sess",
	})
	assert.NotEqual(t, nil, err)
}

func TestGetProfileByEmailNegativeUnknownError(t *testing.T) {
	_, err := pu2.GetProfileByEmail(context.Background(), "asd@asd.asd")
	assert.NotEqual(t, nil, err)
}

type profileRepositoryMock3 struct {
}

func (prm *profileRepositoryMock3) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	return nil, repositoryToUsecaseErrors.ProfileRepositoryEmailDoesntExistError
}

func (prm *profileRepositoryMock3) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string) error {
	return repositoryToUsecaseErrors.ProfileRepositoryEmailExistsError
}

type sessionRepositoryMock3 struct {
}

func (srp *sessionRepositoryMock3) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return nil
}

var pu3 = NewProfileUsecase(&profileRepositoryMock3{}, &sessionRepositoryMock3{}, loggerMock)

func TestUpdateProfileByEmailNegativeKnownError(t *testing.T) {
	err := pu3.UpdateProfileByEmail(context.Background(), &models.Profile{
		Email:    "asd@asd.asd",
		Login:    "Abracodabra",
		Password: "Qwerty123",
	}, "newEmail@asd.asd", &models.Session{
		SessionId: "sess",
	})
	assert.NotEqual(t, nil, err)
}

func TestGetProfileByEmailNegativeKnownError(t *testing.T) {
	_, err := pu3.GetProfileByEmail(context.Background(), "asd@asd.asd")
	assert.NotEqual(t, nil, err)
}

type profileRepositoryMock4 struct {
}

func (prm *profileRepositoryMock4) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	return nil, repositoryToUsecaseErrors.ProfileRepositoryEmailDoesntExistError
}

func (prm *profileRepositoryMock4) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string) error {
	return repositoryToUsecaseErrors.ProfileRepositoryLoginExistsError
}

type sessionRepositoryMock4 struct {
}

func (srp *sessionRepositoryMock4) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return nil
}

var pu4 = NewProfileUsecase(&profileRepositoryMock4{}, &sessionRepositoryMock4{}, loggerMock)

func TestUpdateProfileByEmailNegativeKnownError2(t *testing.T) {
	err := pu4.UpdateProfileByEmail(context.Background(), &models.Profile{
		Email:    "asd@asd.asd",
		Login:    "Abracodabra",
		Password: "Qwerty123",
	}, "newEmail@asd.asd", &models.Session{
		SessionId: "sess",
	})
	assert.NotEqual(t, nil, err)
}
