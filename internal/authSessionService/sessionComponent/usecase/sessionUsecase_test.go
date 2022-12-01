package usecase

import (
	"2022_2_GoTo_team/internal/authSessionService/domain/customErrors/sessionComponentErrors/repositoryToUsecaseErrors"
	repositoryToUsecaseErrors2 "2022_2_GoTo_team/internal/authSessionService/domain/customErrors/userComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/authSessionService/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/authSessionService/domain/interfaces/userComponentInterfaces"
	"2022_2_GoTo_team/internal/authSessionService/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"testing"
)

type sessionRepositoryMock struct {
}

func (ss *sessionRepositoryMock) getSessionsInStorageString() string {
	return "sess1"
}

func (ss *sessionRepositoryMock) CreateSessionForUser(ctx context.Context, email string) (*models.Session, error) {
	return &models.Session{
		SessionId: "sess1",
	}, nil
}

func (ss *sessionRepositoryMock) GetEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	return "asd@asd.asd", nil
}

func (ss *sessionRepositoryMock) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return nil
}

func (ss *sessionRepositoryMock) RemoveSession(ctx context.Context, session *models.Session) error {
	return nil
}

func (ss *sessionRepositoryMock) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	return true, nil
}

type userRepositoryMock struct {
}

func (upsr *userRepositoryMock) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	users := make([]*models.User, 0, 10)
	users = append(users, &models.User{})

	return users, nil
}

func (upsr *userRepositoryMock) CheckUserEmailAndPassword(ctx context.Context, email string, password string) (bool, error) {
	return true, nil
}

func (upsr *userRepositoryMock) GetUserInfoForSessionComponentByEmail(ctx context.Context, email string) (*models.User, error) {
	return &models.User{Email: "asd@asd.asd"}, nil
}

func TestSessionUsecase(t *testing.T) {
	var sessionRepository sessionComponentInterfaces.SessionRepositoryInterface = &sessionRepositoryMock{}
	var userRepository userComponentInterfaces.UserRepositoryInterface = &userRepositoryMock{}
	sessionUsecase := NewSessionUsecase(sessionRepository, userRepository, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	r1, err1 := sessionUsecase.SessionExists(context.Background(), &models.Session{SessionId: "sess1"})
	if r1 != true || err1 != nil {
		t.Error("SessionExists positive test error")
	}
	r2, err2 := sessionUsecase.CreateSessionForUser(context.Background(), "asd@asd.asd", "asd123")
	if r2 == nil || err2 != nil {
		t.Error("CreateSessionForUser positive test error")
	}
	err3 := sessionUsecase.RemoveSession(context.Background(), &models.Session{SessionId: "sess1"})
	if err3 != nil {
		t.Error("RemoveSession positive test error")
	}
	r4, err4 := sessionUsecase.GetUserInfoBySession(context.Background(), &models.Session{SessionId: "sess1"})
	if r4 == nil || err4 != nil {
		t.Error("GetUserInfoBySession positive test error")
	}
	r5, err5 := sessionUsecase.GetUserEmailBySession(context.Background(), &models.Session{SessionId: "sess1"})
	if r5 == "" || err5 != nil {
		t.Error("GetUserEmailBySession positive test error")
	}
	err6 := sessionUsecase.UpdateEmailBySession(context.Background(), &models.Session{SessionId: "sess1"}, "asd@asd.asd")
	if err6 != nil {
		t.Error("UpdateEmailBySession positive test error")
	}
}

type sessionRepositoryNegativeMock struct {
}

func (ss *sessionRepositoryNegativeMock) getSessionsInStorageString() string {
	return ""
}

func (ss *sessionRepositoryNegativeMock) CreateSessionForUser(ctx context.Context, email string) (*models.Session, error) {
	return nil, errors.New("err")
}

func (ss *sessionRepositoryNegativeMock) GetEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	return "", errors.New("err")
}

func (ss *sessionRepositoryNegativeMock) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return errors.New("err")
}

func (ss *sessionRepositoryNegativeMock) RemoveSession(ctx context.Context, session *models.Session) error {
	return errors.New("err")
}

func (ss *sessionRepositoryNegativeMock) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	return false, errors.New("err")
}

type userRepositoryNegativeMock struct {
}

func (upsr *userRepositoryNegativeMock) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	users := make([]*models.User, 0, 10)
	users = append(users, &models.User{})

	return nil, errors.New("err")
}

func (upsr *userRepositoryNegativeMock) CheckUserEmailAndPassword(ctx context.Context, email string, password string) (bool, error) {
	return false, errors.New("err")
}

func (upsr *userRepositoryNegativeMock) GetUserInfoForSessionComponentByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, errors.New("err")
}

func TestSessionUsecaseNegative(t *testing.T) {
	var sessionRepository sessionComponentInterfaces.SessionRepositoryInterface = &sessionRepositoryNegativeMock{}
	var userRepository userComponentInterfaces.UserRepositoryInterface = &userRepositoryNegativeMock{}
	sessionUsecase := NewSessionUsecase(sessionRepository, userRepository, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	r1, err1 := sessionUsecase.SessionExists(context.Background(), &models.Session{SessionId: "sess1"})
	if r1 == true || err1 == nil {
		t.Error("SessionExists negative test error")
	}
	r2, err2 := sessionUsecase.CreateSessionForUser(context.Background(), "asd@asd.asd", "asd123")
	if r2 != nil || err2 == nil {
		t.Error("CreateSessionForUser negative test error")
	}
	err3 := sessionUsecase.RemoveSession(context.Background(), &models.Session{SessionId: "sess1"})
	if err3 == nil {
		t.Error("RemoveSession negative test error")
	}
	r4, err4 := sessionUsecase.GetUserInfoBySession(context.Background(), &models.Session{SessionId: "sess1"})
	if r4 != nil || err4 == nil {
		t.Error("GetUserInfoBySession negative test error")
	}
	r5, err5 := sessionUsecase.GetUserEmailBySession(context.Background(), &models.Session{SessionId: "sess1"})
	if r5 != "" || err5 == nil {
		t.Error("GetUserEmailBySession negative test error")
	}
	err6 := sessionUsecase.UpdateEmailBySession(context.Background(), &models.Session{SessionId: "sess1"}, "asd@asd.asd")
	if err6 == nil {
		t.Error("UpdateEmailBySession negative test error")
	}

	_, err7 := sessionUsecase.CreateSessionForUser(context.Background(), "asd@asd.asd", "a")
	if err7 == nil {
		t.Error("UpdateEmailBySession negative test error")
	}
	_, err8 := sessionUsecase.CreateSessionForUser(context.Background(), "asd", "aasd123")
	if err8 == nil {
		t.Error("UpdateEmailBySession negative test error")
	}
}

type sessionRepositoryNegativeMock2 struct {
}

func (ss *sessionRepositoryNegativeMock2) getSessionsInStorageString() string {
	return ""
}

func (ss *sessionRepositoryNegativeMock2) CreateSessionForUser(ctx context.Context, email string) (*models.Session, error) {
	return nil, errors.New("err")
}

func (ss *sessionRepositoryNegativeMock2) GetEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	return "", repositoryToUsecaseErrors.SessionRepositoryEmailDoesntExistError
}

func (ss *sessionRepositoryNegativeMock2) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return errors.New("err")
}

func (ss *sessionRepositoryNegativeMock2) RemoveSession(ctx context.Context, session *models.Session) error {
	return errors.New("err")
}

func (ss *sessionRepositoryNegativeMock2) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	return false, errors.New("err")
}

type userRepositoryNegativeMock2 struct {
}

func (upsr *userRepositoryNegativeMock2) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	users := make([]*models.User, 0, 10)
	users = append(users, &models.User{})

	return nil, errors.New("err")
}

func (upsr *userRepositoryNegativeMock2) CheckUserEmailAndPassword(ctx context.Context, email string, password string) (bool, error) {
	return false, nil
}

func (upsr *userRepositoryNegativeMock2) GetUserInfoForSessionComponentByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, errors.New("err")
}

func TestSessionUsecaseNegative2(t *testing.T) {
	var sessionRepository sessionComponentInterfaces.SessionRepositoryInterface = &sessionRepositoryNegativeMock2{}
	var userRepository userComponentInterfaces.UserRepositoryInterface = &userRepositoryNegativeMock2{}
	sessionUsecase := NewSessionUsecase(sessionRepository, userRepository, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	r1, err1 := sessionUsecase.SessionExists(context.Background(), &models.Session{SessionId: "sess1"})
	if r1 == true || err1 == nil {
		t.Error("SessionExists negative test error")
	}
	r2, err2 := sessionUsecase.CreateSessionForUser(context.Background(), "asd@asd.asd", "asd123")
	if r2 != nil || err2 == nil {
		t.Error("CreateSessionForUser negative test error")
	}
	err3 := sessionUsecase.RemoveSession(context.Background(), &models.Session{SessionId: "sess1"})
	if err3 == nil {
		t.Error("RemoveSession negative test error")
	}
	r4, err4 := sessionUsecase.GetUserInfoBySession(context.Background(), &models.Session{SessionId: "sess1"})
	if r4 != nil || err4 == nil {
		t.Error("GetUserInfoBySession negative test error")
	}
	r5, err5 := sessionUsecase.GetUserEmailBySession(context.Background(), &models.Session{SessionId: "sess1"})
	if r5 != "" || err5 == nil {
		t.Error("GetUserEmailBySession negative test error")
	}
	err6 := sessionUsecase.UpdateEmailBySession(context.Background(), &models.Session{SessionId: "sess1"}, "asd@asd.asd")
	if err6 == nil {
		t.Error("UpdateEmailBySession negative test error")
	}
}

type sessionRepositoryNegativeMock3 struct {
}

func (ss *sessionRepositoryNegativeMock3) getSessionsInStorageString() string {
	return ""
}

func (ss *sessionRepositoryNegativeMock3) CreateSessionForUser(ctx context.Context, email string) (*models.Session, error) {
	return nil, errors.New("err")
}

func (ss *sessionRepositoryNegativeMock3) GetEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	return "", nil
}

func (ss *sessionRepositoryNegativeMock3) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return errors.New("err")
}

func (ss *sessionRepositoryNegativeMock3) RemoveSession(ctx context.Context, session *models.Session) error {
	return errors.New("err")
}

func (ss *sessionRepositoryNegativeMock3) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	return false, errors.New("err")
}

type userRepositoryNegativeMock3 struct {
}

func (upsr *userRepositoryNegativeMock3) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	users := make([]*models.User, 0, 10)
	users = append(users, &models.User{})

	return nil, errors.New("err")
}

func (upsr *userRepositoryNegativeMock3) CheckUserEmailAndPassword(ctx context.Context, email string, password string) (bool, error) {
	return true, nil
}

func (upsr *userRepositoryNegativeMock3) GetUserInfoForSessionComponentByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, repositoryToUsecaseErrors2.UserRepositoryEmailDoesntExistError
}

func TestSessionUsecaseNegative3(t *testing.T) {
	var sessionRepository sessionComponentInterfaces.SessionRepositoryInterface = &sessionRepositoryNegativeMock3{}
	var userRepository userComponentInterfaces.UserRepositoryInterface = &userRepositoryNegativeMock3{}
	sessionUsecase := NewSessionUsecase(sessionRepository, userRepository, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	r1, err1 := sessionUsecase.SessionExists(context.Background(), &models.Session{SessionId: "sess1"})
	if r1 == true || err1 == nil {
		t.Error("SessionExists negative test error")
	}
	r2, err2 := sessionUsecase.CreateSessionForUser(context.Background(), "asd@asd.asd", "asd123")
	if r2 != nil || err2 == nil {
		t.Error("CreateSessionForUser negative test error")
	}
	err3 := sessionUsecase.RemoveSession(context.Background(), &models.Session{SessionId: "sess1"})
	if err3 == nil {
		t.Error("RemoveSession negative test error")
	}
	r4, err4 := sessionUsecase.GetUserInfoBySession(context.Background(), &models.Session{SessionId: "sess1"})
	if r4 != nil || err4 == nil {
		t.Error("GetUserInfoBySession negative test error")
	}
	r5, err5 := sessionUsecase.GetUserEmailBySession(context.Background(), &models.Session{SessionId: "sess1"})
	if r5 != "" || err5 != nil {
		t.Error("GetUserEmailBySession negative test error")
	}
	err6 := sessionUsecase.UpdateEmailBySession(context.Background(), &models.Session{SessionId: "sess1"}, "asd@asd.asd")
	if err6 == nil {
		t.Error("UpdateEmailBySession negative test error")
	}
}

type sessionRepositoryNegativeMock4 struct {
}

func (ss *sessionRepositoryNegativeMock4) getSessionsInStorageString() string {
	return ""
}

func (ss *sessionRepositoryNegativeMock4) CreateSessionForUser(ctx context.Context, email string) (*models.Session, error) {
	return nil, errors.New("err")
}

func (ss *sessionRepositoryNegativeMock4) GetEmailBySession(ctx context.Context, session *models.Session) (string, error) {
	return "", nil
}

func (ss *sessionRepositoryNegativeMock4) UpdateEmailBySession(ctx context.Context, session *models.Session, newEmail string) error {
	return errors.New("err")
}

func (ss *sessionRepositoryNegativeMock4) RemoveSession(ctx context.Context, session *models.Session) error {
	return errors.New("err")
}

func (ss *sessionRepositoryNegativeMock4) SessionExists(ctx context.Context, session *models.Session) (bool, error) {
	return false, errors.New("err")
}

type userRepositoryNegativeMock4 struct {
}

func (upsr *userRepositoryNegativeMock4) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	users := make([]*models.User, 0, 10)
	users = append(users, &models.User{})

	return nil, errors.New("err")
}

func (upsr *userRepositoryNegativeMock4) CheckUserEmailAndPassword(ctx context.Context, email string, password string) (bool, error) {
	return true, nil
}

func (upsr *userRepositoryNegativeMock4) GetUserInfoForSessionComponentByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, errors.New("err")
}

func TestSessionUsecaseNegative4(t *testing.T) {
	var sessionRepository sessionComponentInterfaces.SessionRepositoryInterface = &sessionRepositoryNegativeMock4{}
	var userRepository userComponentInterfaces.UserRepositoryInterface = &userRepositoryNegativeMock4{}
	sessionUsecase := NewSessionUsecase(sessionRepository, userRepository, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	r1, err1 := sessionUsecase.SessionExists(context.Background(), &models.Session{SessionId: "sess1"})
	if r1 == true || err1 == nil {
		t.Error("SessionExists negative test error")
	}
	r2, err2 := sessionUsecase.CreateSessionForUser(context.Background(), "asd@asd.asd", "asd123")
	if r2 != nil || err2 == nil {
		t.Error("CreateSessionForUser negative test error")
	}
	err3 := sessionUsecase.RemoveSession(context.Background(), &models.Session{SessionId: "sess1"})
	if err3 == nil {
		t.Error("RemoveSession negative test error")
	}
	r4, err4 := sessionUsecase.GetUserInfoBySession(context.Background(), &models.Session{SessionId: "sess1"})
	if r4 != nil || err4 == nil {
		t.Error("GetUserInfoBySession negative test error")
	}
	r5, err5 := sessionUsecase.GetUserEmailBySession(context.Background(), &models.Session{SessionId: "sess1"})
	if r5 != "" || err5 != nil {
		t.Error("GetUserEmailBySession negative test error")
	}
	err6 := sessionUsecase.UpdateEmailBySession(context.Background(), &models.Session{SessionId: "sess1"}, "asd@asd.asd")
	if err6 == nil {
		t.Error("UpdateEmailBySession negative test error")
	}
}
