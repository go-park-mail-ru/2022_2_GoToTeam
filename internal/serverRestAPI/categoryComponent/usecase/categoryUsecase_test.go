package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/categoryComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
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

type categoryRepositoryMock2 struct {
}

func (crm *categoryRepositoryMock2) GetCategoryInfo(ctx context.Context, category string) (*models.Category, error) {
	return &models.Category{}, errors.New("err")
}

func (crm *categoryRepositoryMock2) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	return []*models.Category{}, errors.New("err")
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

type categoryRepositoryMock3 struct {
}

func (crm *categoryRepositoryMock3) GetCategoryInfo(ctx context.Context, category string) (*models.Category, error) {
	return &models.Category{}, repositoryToUsecaseErrors.CategoryRepositoryCategoryDoesntExistError
}

func (crm *categoryRepositoryMock3) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	return []*models.Category{}, errors.New("err")
}

func TestGetCategoryInfoNegativeKnownError(t *testing.T) {
	cu := NewCategoryUsecase(&categoryRepositoryMock3{}, loggerMock)

	_, err := cu.GetCategoryInfo(context.Background(), "programming")
	assert.NotEqual(t, nil, err)
}
