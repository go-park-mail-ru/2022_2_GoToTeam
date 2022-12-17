package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/categoryComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/domain"
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

type categoryRepositoryMock struct {
}

func (crm *categoryRepositoryMock) GetCategoryInfo(ctx context.Context, category string) (*models.Category, error) {
	return &models.Category{}, nil
}

func (crm *categoryRepositoryMock) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	return []*models.Category{}, nil
}

func (crm *categoryRepositoryMock) IsUserSubscribedOnCategory(ctx context.Context, userEmail string, categoryName string) (bool, error) {
	return true, nil
}

func (crm *categoryRepositoryMock) SubscribeOnCategory(ctx context.Context, email string, categoryName string) error {
	return nil
}

func (crm *categoryRepositoryMock) UnsubscribeFromCategory(ctx context.Context, email string, categoryName string) (int64, error) {
	return 1, nil
}

func TestGetCategoryInfo(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock{}, loggerMock)

	res, err := cu.GetCategoryInfo(context.Background(), "programming")
	assert.NotEqual(t, nil, res)
	assert.Equal(t, nil, err)
}

func TestGetCategoryList(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock{}, loggerMock)

	res, err := cu.GetCategoryList(context.Background())
	assert.NotEqual(t, nil, res)
	assert.Equal(t, nil, err)
}

func TestIsUserSubscribedOnCategory(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock{}, loggerMock)

	res, err := cu.IsUserSubscribedOnCategory(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), "category")
	assert.NotEqual(t, nil, res)
	assert.Equal(t, nil, err)
}

func TestSubscribeOnCategory(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock{}, loggerMock)

	err := cu.SubscribeOnCategory(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), "category")
	assert.Equal(t, nil, err)
}

func TestUnsubscribeOnCategory(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock{}, loggerMock)

	err := cu.UnsubscribeFromCategory(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), "category")
	assert.Equal(t, nil, err)
}

type categoryRepositoryMock2 struct {
}

func (crm *categoryRepositoryMock2) GetCategoryInfo(ctx context.Context, category string) (*models.Category, error) {
	return &models.Category{}, errors.New("err")
}

func (crm *categoryRepositoryMock2) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	return []*models.Category{}, errors.New("err")
}

func (crm *categoryRepositoryMock2) IsUserSubscribedOnCategory(ctx context.Context, userEmail string, categoryName string) (bool, error) {
	return true, nil
}

func (crm *categoryRepositoryMock2) SubscribeOnCategory(ctx context.Context, email string, categoryName string) error {
	return nil
}

func (crm *categoryRepositoryMock2) UnsubscribeFromCategory(ctx context.Context, email string, categoryName string) (int64, error) {
	return 1, nil
}

func TestGetCategoryInfoNegativeUnknownError(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock2{}, loggerMock)

	_, err := cu.GetCategoryInfo(context.Background(), "programming")
	assert.NotEqual(t, nil, err)
}

func TestGetCategoryListNegative(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock2{}, loggerMock)

	_, err := cu.GetCategoryList(context.Background())
	assert.NotEqual(t, nil, err)
}

func TestIsUserSubscribedOnCategoryNegativeEmptyContextEmail(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock2{}, loggerMock)

	res, err := cu.IsUserSubscribedOnCategory(context.Background(), "category")
	assert.Equal(t, false, res)
	assert.NotEqual(t, nil, err)
}

func TestSubscribeOnCategoryNegativeEmptyContextEmail(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock2{}, loggerMock)

	err := cu.SubscribeOnCategory(context.Background(), "category")
	assert.NotEqual(t, nil, err)
}

func TestUnsubscribeOnCategoryNegativeEmptyContextEmail(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock2{}, loggerMock)

	err := cu.UnsubscribeFromCategory(context.Background(), "category")
	assert.NotEqual(t, nil, err)
}

type categoryRepositoryMock3 struct {
}

func (crm *categoryRepositoryMock3) GetCategoryInfo(ctx context.Context, category string) (*models.Category, error) {
	return &models.Category{}, repositoryToUsecaseErrors.CategoryRepositoryCategoryDoesntExistError
}

func (crm *categoryRepositoryMock3) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	return []*models.Category{}, errors.New("err")
}

func (crm *categoryRepositoryMock3) IsUserSubscribedOnCategory(ctx context.Context, userEmail string, categoryName string) (bool, error) {
	return false, errors.New("err")
}

func (crm *categoryRepositoryMock3) SubscribeOnCategory(ctx context.Context, email string, categoryName string) error {
	return errors.New("err")
}

func (crm *categoryRepositoryMock3) UnsubscribeFromCategory(ctx context.Context, email string, categoryName string) (int64, error) {
	return 0, errors.New("err")
}

func TestGetCategoryInfoNegativeKnownError(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock3{}, loggerMock)

	_, err := cu.GetCategoryInfo(context.Background(), "programming")
	assert.NotEqual(t, nil, err)
}

func TestIsUserSubscribedOnCategoryNegativeUnknownError(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock3{}, loggerMock)

	res, err := cu.IsUserSubscribedOnCategory(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), "category")
	assert.Equal(t, false, res)
	assert.NotEqual(t, nil, err)
}

func TestSubscribeOnCategoryNegativeError(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock3{}, loggerMock)

	err := cu.SubscribeOnCategory(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), "category")
	assert.NotEqual(t, nil, err)
}

func TestUnsubscribeOnCategoryNegativeError(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock3{}, loggerMock)

	err := cu.UnsubscribeFromCategory(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), "category")
	assert.NotEqual(t, nil, err)
}
