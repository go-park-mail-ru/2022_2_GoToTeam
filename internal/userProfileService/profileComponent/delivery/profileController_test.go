package delivery

import (
	"2022_2_GoTo_team/internal/userProfileService/domain/customErrors/profileComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/userProfileService/domain/models"
	"2022_2_GoTo_team/pkg/domain/grpcProtos/userProfileServiceGrpcProtos"
	"2022_2_GoTo_team/pkg/utils/errorsUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"testing"
)

type profileUsecaseMock struct {
}

func (pum *profileUsecaseMock) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	return &models.Profile{Email: "asd@asd.asd"}, nil
}

func (pum *profileUsecaseMock) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error {
	return nil
}

func TestProfileDelivery(t *testing.T) {
	profileDelivery := NewProfileDelivery(&profileUsecaseMock{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := profileDelivery.GetProfileByEmail(context.Background(), &userProfileServiceGrpcProtos.UserEmail{Email: "asd@asd.asd"})
	if err != nil {
		t.Error(err)
	}

	_, err = profileDelivery.UpdateProfileByEmail(context.Background(), &userProfileServiceGrpcProtos.UpdateProfileData{
		Profile: &userProfileServiceGrpcProtos.Profile{
			Email:         "",
			Login:         "",
			Password:      "",
			Username:      "",
			AvatarImgPath: "",
		},
	})
	if err != nil {
		t.Error(err)
	}
}

type profileUsecaseMock2 struct {
}

func (pum *profileUsecaseMock2) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	return &models.Profile{Email: "asd@asd.asd"}, errors.New("err")
}

func (pum *profileUsecaseMock2) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error {
	return errors.New("err")
}

func TestProfileDeliveryNegative(t *testing.T) {
	profileDelivery := NewProfileDelivery(&profileUsecaseMock2{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := profileDelivery.GetProfileByEmail(context.Background(), &userProfileServiceGrpcProtos.UserEmail{Email: "asd@asd.asd"})
	if err == nil {
		t.Error(err)
	}

	_, err = profileDelivery.UpdateProfileByEmail(context.Background(), &userProfileServiceGrpcProtos.UpdateProfileData{
		Profile: &userProfileServiceGrpcProtos.Profile{
			Email:         "",
			Login:         "",
			Password:      "",
			Username:      "",
			AvatarImgPath: "",
		},
	})
	if err == nil {
		t.Error(err)
	}
}

type profileUsecaseMock3 struct {
}

func (pum *profileUsecaseMock3) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	return &models.Profile{Email: "asd@asd.asd"}, errorsUtils.WrapError("err", &usecaseToDeliveryErrors.EmailDoesntExistError{})
}

func (pum *profileUsecaseMock3) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error {
	return errorsUtils.WrapError("err", &usecaseToDeliveryErrors.EmailIsNotValidError{})
}

func TestProfileDeliveryNegative2(t *testing.T) {
	profileDelivery := NewProfileDelivery(&profileUsecaseMock3{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := profileDelivery.GetProfileByEmail(context.Background(), &userProfileServiceGrpcProtos.UserEmail{Email: "asd@asd.asd"})
	if err == nil {
		t.Error(err)
	}

	_, err = profileDelivery.UpdateProfileByEmail(context.Background(), &userProfileServiceGrpcProtos.UpdateProfileData{
		Profile: &userProfileServiceGrpcProtos.Profile{
			Email:         "",
			Login:         "",
			Password:      "",
			Username:      "",
			AvatarImgPath: "",
		},
	})
	if err == nil {
		t.Error(err)
	}
}

type profileUsecaseMock4 struct {
}

func (pum *profileUsecaseMock4) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	return &models.Profile{Email: "asd@asd.asd"}, errorsUtils.WrapError("err", &usecaseToDeliveryErrors.EmailDoesntExistError{})
}

func (pum *profileUsecaseMock4) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error {
	return errorsUtils.WrapError("err", &usecaseToDeliveryErrors.LoginIsNotValidError{})
}

func TestProfileDeliveryNegative3(t *testing.T) {
	profileDelivery := NewProfileDelivery(&profileUsecaseMock4{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := profileDelivery.GetProfileByEmail(context.Background(), &userProfileServiceGrpcProtos.UserEmail{Email: "asd@asd.asd"})
	if err == nil {
		t.Error(err)
	}

	_, err = profileDelivery.UpdateProfileByEmail(context.Background(), &userProfileServiceGrpcProtos.UpdateProfileData{
		Profile: &userProfileServiceGrpcProtos.Profile{
			Email:         "",
			Login:         "",
			Password:      "",
			Username:      "",
			AvatarImgPath: "",
		},
	})
	if err == nil {
		t.Error(err)
	}
}

type profileUsecaseMock5 struct {
}

func (pum *profileUsecaseMock5) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	return &models.Profile{Email: "asd@asd.asd"}, errorsUtils.WrapError("err", &usecaseToDeliveryErrors.PasswordIsNotValidError{})
}

func (pum *profileUsecaseMock5) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error {
	return errorsUtils.WrapError("err", &usecaseToDeliveryErrors.PasswordIsNotValidError{})
}

func TestProfileDeliveryNegative4(t *testing.T) {
	profileDelivery := NewProfileDelivery(&profileUsecaseMock5{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := profileDelivery.GetProfileByEmail(context.Background(), &userProfileServiceGrpcProtos.UserEmail{Email: "asd@asd.asd"})
	if err == nil {
		t.Error(err)
	}

	_, err = profileDelivery.UpdateProfileByEmail(context.Background(), &userProfileServiceGrpcProtos.UpdateProfileData{
		Profile: &userProfileServiceGrpcProtos.Profile{
			Email:         "",
			Login:         "",
			Password:      "",
			Username:      "",
			AvatarImgPath: "",
		},
	})
	if err == nil {
		t.Error(err)
	}
}

type profileUsecaseMock6 struct {
}

func (pum *profileUsecaseMock6) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	return &models.Profile{Email: "asd@asd.asd"}, errorsUtils.WrapError("err", &usecaseToDeliveryErrors.EmailExistsError{})
}

func (pum *profileUsecaseMock6) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error {
	return errorsUtils.WrapError("err", &usecaseToDeliveryErrors.EmailExistsError{})
}

func TestProfileDeliveryNegative5(t *testing.T) {
	profileDelivery := NewProfileDelivery(&profileUsecaseMock6{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := profileDelivery.GetProfileByEmail(context.Background(), &userProfileServiceGrpcProtos.UserEmail{Email: "asd@asd.asd"})
	if err == nil {
		t.Error(err)
	}

	_, err = profileDelivery.UpdateProfileByEmail(context.Background(), &userProfileServiceGrpcProtos.UpdateProfileData{
		Profile: &userProfileServiceGrpcProtos.Profile{
			Email:         "",
			Login:         "",
			Password:      "",
			Username:      "",
			AvatarImgPath: "",
		},
	})
	if err == nil {
		t.Error(err)
	}
}

type profileUsecaseMock7 struct {
}

func (pum *profileUsecaseMock7) GetProfileByEmail(ctx context.Context, email string) (*models.Profile, error) {
	return &models.Profile{Email: "asd@asd.asd"}, errorsUtils.WrapError("err", &usecaseToDeliveryErrors.LoginExistsError{})
}

func (pum *profileUsecaseMock7) UpdateProfileByEmail(ctx context.Context, newProfile *models.Profile, email string, session *models.Session) error {
	return errorsUtils.WrapError("err", &usecaseToDeliveryErrors.LoginExistsError{})
}

func TestProfileDeliveryNegative6(t *testing.T) {
	profileDelivery := NewProfileDelivery(&profileUsecaseMock7{}, &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
		"requestId": "asd",
		"userEmail": "asd",
	})})

	_, err := profileDelivery.GetProfileByEmail(context.Background(), &userProfileServiceGrpcProtos.UserEmail{Email: "asd@asd.asd"})
	if err == nil {
		t.Error(err)
	}

	_, err = profileDelivery.UpdateProfileByEmail(context.Background(), &userProfileServiceGrpcProtos.UpdateProfileData{
		Profile: &userProfileServiceGrpcProtos.Profile{
			Email:         "",
			Login:         "",
			Password:      "",
			Username:      "",
			AvatarImgPath: "",
		},
	})
	if err == nil {
		t.Error(err)
	}
}
