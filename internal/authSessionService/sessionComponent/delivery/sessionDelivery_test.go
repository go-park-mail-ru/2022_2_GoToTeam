package delivery

import (
	"2022_2_GoTo_team/internal/authSessionService/domain/customErrors/sessionComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/authSessionService/domain/models"
	"2022_2_GoTo_team/pkg/domain/grpcProtos/authSessionServiceGrpcProtos"
	"2022_2_GoTo_team/pkg/utils/errorsUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"testing"
)

type sessionUsecaseMock struct {
}

func (sum *sessionUsecaseMock) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	return false, nil
}

func (sum *sessionUsecaseMock) CreateSessionForUser(ctx context.Context, email string, password string) (*models.Session, error) {
	return &models.Session{
		SessionId: "sess1",
	}, nil
}

func (sum *sessionUsecaseMock) RemoveSession(ctx context.Context, session *models.Session) error {
	return nil
}

func (sum *sessionUsecaseMock) GetUserInfoBySession(ctx context.Context, session *models.Session) (*models.User, error) {
	return &models.User{Email: "asd@asd.asd"}, nil
}

func (sum *sessionUsecaseMock) GetUserEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	return "str", nil
}

func (sum *sessionUsecaseMock) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return nil
}

func TestSessionDelivery(t *testing.T) {
	sessionDelivery := NewSessionDelivery(&sessionUsecaseMock{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := sessionDelivery.SessionExists(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err != nil {
		t.Error(err)
	}

	_, err = sessionDelivery.CreateSessionForUser(context.Background(), &authSessionServiceGrpcProtos.UserAccountData{})
	if err != nil {
		t.Error(err)
	}

	_, err = sessionDelivery.RemoveSession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err != nil {
		t.Error(err)
	}

	_, err = sessionDelivery.GetUserInfoBySession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err != nil {
		t.Error(err)
	}

	_, err = sessionDelivery.GetUserEmailBySession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err != nil {
		t.Error(err)
	}

	_, err = sessionDelivery.UpdateEmailBySession(context.Background(), &authSessionServiceGrpcProtos.UpdateEmailData{Session: &authSessionServiceGrpcProtos.Session{
		SessionId: "sess1",
	}})
	if err != nil {
		t.Error(err)
	}
}

type sessionUsecaseMock2 struct {
}

func (sum *sessionUsecaseMock2) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	return false, errors.New("err")
}

func (sum *sessionUsecaseMock2) CreateSessionForUser(ctx context.Context, email string, password string) (*models.Session, error) {
	return &models.Session{
		SessionId: "sess1",
	}, errors.New("err")
}

func (sum *sessionUsecaseMock2) RemoveSession(ctx context.Context, session *models.Session) error {
	return errors.New("err")
}

func (sum *sessionUsecaseMock2) GetUserInfoBySession(ctx context.Context, session *models.Session) (*models.User, error) {
	return &models.User{Email: "asd@asd.asd"}, errors.New("err")
}

func (sum *sessionUsecaseMock2) GetUserEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	return "str", errors.New("err")
}

func (sum *sessionUsecaseMock2) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return errors.New("err")
}

func TestSessionDeliveryNegative(t *testing.T) {
	sessionDelivery := NewSessionDelivery(&sessionUsecaseMock2{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := sessionDelivery.SessionExists(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.CreateSessionForUser(context.Background(), &authSessionServiceGrpcProtos.UserAccountData{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.RemoveSession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.GetUserInfoBySession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.GetUserEmailBySession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.UpdateEmailBySession(context.Background(), &authSessionServiceGrpcProtos.UpdateEmailData{Session: &authSessionServiceGrpcProtos.Session{
		SessionId: "sess1",
	}})
	if err == nil {
		t.Error(err)
	}
}

type sessionUsecaseMock3 struct {
}

func (sum *sessionUsecaseMock3) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	return false, errors.New("err")
}

func (sum *sessionUsecaseMock3) CreateSessionForUser(ctx context.Context, email string, password string) (*models.Session, error) {
	return &models.Session{
		SessionId: "sess1",
	}, errorsUtils.WrapError("err", &usecaseToDeliveryErrors.EmailIsNotValidError{})
}

func (sum *sessionUsecaseMock3) RemoveSession(ctx context.Context, session *models.Session) error {
	return errors.New("err")
}

func (sum *sessionUsecaseMock3) GetUserInfoBySession(ctx context.Context, session *models.Session) (*models.User, error) {
	return &models.User{Email: "asd@asd.asd"}, errorsUtils.WrapError("err", &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{})
}

func (sum *sessionUsecaseMock3) GetUserEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	return "str", errorsUtils.WrapError("err", &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{})
}

func (sum *sessionUsecaseMock3) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return errors.New("err")
}

func TestSessionDeliveryNegative2(t *testing.T) {
	sessionDelivery := NewSessionDelivery(&sessionUsecaseMock3{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := sessionDelivery.SessionExists(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.CreateSessionForUser(context.Background(), &authSessionServiceGrpcProtos.UserAccountData{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.RemoveSession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.GetUserInfoBySession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.GetUserEmailBySession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.UpdateEmailBySession(context.Background(), &authSessionServiceGrpcProtos.UpdateEmailData{Session: &authSessionServiceGrpcProtos.Session{
		SessionId: "sess1",
	}})
	if err == nil {
		t.Error(err)
	}
}

type sessionUsecaseMock4 struct {
}

func (sum *sessionUsecaseMock4) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	return false, errors.New("err")
}

func (sum *sessionUsecaseMock4) CreateSessionForUser(ctx context.Context, email string, password string) (*models.Session, error) {
	return &models.Session{
		SessionId: "sess1",
	}, errorsUtils.WrapError("err", &usecaseToDeliveryErrors.PasswordIsNotValidError{})
}

func (sum *sessionUsecaseMock4) RemoveSession(ctx context.Context, session *models.Session) error {
	return errors.New("err")
}

func (sum *sessionUsecaseMock4) GetUserInfoBySession(ctx context.Context, session *models.Session) (*models.User, error) {
	return &models.User{Email: "asd@asd.asd"}, errorsUtils.WrapError("err", &usecaseToDeliveryErrors.UserForSessionDoesntExistError{})
}

func (sum *sessionUsecaseMock4) GetUserEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	return "str", errorsUtils.WrapError("err", &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{})
}

func (sum *sessionUsecaseMock4) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return errors.New("err")
}

func TestSessionDeliveryNegative3(t *testing.T) {
	sessionDelivery := NewSessionDelivery(&sessionUsecaseMock4{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := sessionDelivery.SessionExists(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.CreateSessionForUser(context.Background(), &authSessionServiceGrpcProtos.UserAccountData{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.RemoveSession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.GetUserInfoBySession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.GetUserEmailBySession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.UpdateEmailBySession(context.Background(), &authSessionServiceGrpcProtos.UpdateEmailData{Session: &authSessionServiceGrpcProtos.Session{
		SessionId: "sess1",
	}})
	if err == nil {
		t.Error(err)
	}
}

type sessionUsecaseMock5 struct {
}

func (sum *sessionUsecaseMock5) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	return false, errors.New("err")
}

func (sum *sessionUsecaseMock5) CreateSessionForUser(ctx context.Context, email string, password string) (*models.Session, error) {
	return &models.Session{
		SessionId: "sess1",
	}, errorsUtils.WrapError("err", &usecaseToDeliveryErrors.IncorrectEmailOrPasswordError{})
}

func (sum *sessionUsecaseMock5) RemoveSession(ctx context.Context, session *models.Session) error {
	return errors.New("err")
}

func (sum *sessionUsecaseMock5) GetUserInfoBySession(ctx context.Context, session *models.Session) (*models.User, error) {
	return &models.User{Email: "asd@asd.asd"}, errorsUtils.WrapError("err", &usecaseToDeliveryErrors.UserForSessionDoesntExistError{})
}

func (sum *sessionUsecaseMock5) GetUserEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	return "str", errorsUtils.WrapError("err", &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{})
}

func (sum *sessionUsecaseMock5) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return errors.New("err")
}

func TestSessionDeliveryNegative4(t *testing.T) {
	sessionDelivery := NewSessionDelivery(&sessionUsecaseMock5{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := sessionDelivery.SessionExists(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.CreateSessionForUser(context.Background(), &authSessionServiceGrpcProtos.UserAccountData{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.RemoveSession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.GetUserInfoBySession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.GetUserEmailBySession(context.Background(), &authSessionServiceGrpcProtos.Session{})
	if err == nil {
		t.Error(err)
	}

	_, err = sessionDelivery.UpdateEmailBySession(context.Background(), &authSessionServiceGrpcProtos.UpdateEmailData{Session: &authSessionServiceGrpcProtos.Session{
		SessionId: "sess1",
	}})
	if err == nil {
		t.Error(err)
	}
}
