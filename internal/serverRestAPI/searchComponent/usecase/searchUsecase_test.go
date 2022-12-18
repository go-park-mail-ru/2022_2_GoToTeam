package usecase

import (
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

type searchRepositoryMock struct {
}

func (srm *searchRepositoryMock) GetArticlesByTag(ctx context.Context, tag string) ([]*models.Article, error) {
	return []*models.Article{}, nil
}

func (srm *searchRepositoryMock) GetArticlesBySearchParameters(ctx context.Context, substringToSearch string, login string, categoryName string, tagName string) ([]*models.Article, error) {
	return []*models.Article{}, nil
}

func (srm *searchRepositoryMock) TagExists(ctx context.Context, tag string) (bool, error) {
	return true, nil
}

func TestGetArticlesByTag(t *testing.T) {
	su := NewSearchUsecase(&searchRepositoryMock{}, loggerMock)

	res, err := su.GetArticlesByTag(context.Background(), "java")
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, res)
}

func TestGetArticlesBySearchParameters(t *testing.T) {
	su := NewSearchUsecase(&searchRepositoryMock{}, loggerMock)

	res, err := su.GetArticlesBySearchParameters(context.Background(), "a", "b", "c", "d")
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, res)
}

type searchRepositoryMock2 struct {
}

func (srm *searchRepositoryMock2) GetArticlesByTag(ctx context.Context, tag string) ([]*models.Article, error) {
	return []*models.Article{}, nil
}

func (srm *searchRepositoryMock2) GetArticlesBySearchParameters(ctx context.Context, substringToSearch string, login string, categoryName string, tagName string) ([]*models.Article, error) {
	return []*models.Article{}, errors.New("err")
}

func (srm *searchRepositoryMock2) TagExists(ctx context.Context, tag string) (bool, error) {
	return false, nil
}

func TestGetArticlesByTagNegativeTagExistsFalse(t *testing.T) {
	su := NewSearchUsecase(&searchRepositoryMock2{}, loggerMock)

	res, err := su.GetArticlesByTag(context.Background(), "java")
	assert.NotEqual(t, nil, err)
	if res != nil {
		t.Error()
	}
}

func TestGetArticlesBySearchParametersNegative(t *testing.T) {
	su := NewSearchUsecase(&searchRepositoryMock2{}, loggerMock)

	res, err := su.GetArticlesBySearchParameters(context.Background(), "a", "b", "c", "d")
	assert.NotEqual(t, nil, err)
	if res != nil {
		t.Error()
	}
}

type searchRepositoryMock3 struct {
}

func (srm *searchRepositoryMock3) GetArticlesByTag(ctx context.Context, tag string) ([]*models.Article, error) {
	return []*models.Article{}, nil
}

func (srm *searchRepositoryMock3) GetArticlesBySearchParameters(ctx context.Context, substringToSearch string, login string, categoryName string, tagName string) ([]*models.Article, error) {
	return []*models.Article{}, errors.New("err")
}

func (srm *searchRepositoryMock3) TagExists(ctx context.Context, tag string) (bool, error) {
	return false, errors.New("err")
}

func TestGetArticlesByTagNegativeKnownError(t *testing.T) {
	su := NewSearchUsecase(&searchRepositoryMock3{}, loggerMock)

	res, err := su.GetArticlesByTag(context.Background(), "java")
	assert.NotEqual(t, nil, err)
	if res != nil {
		t.Error()
	}
}

type searchRepositoryMock4 struct {
}

func (srm *searchRepositoryMock4) GetArticlesByTag(ctx context.Context, tag string) ([]*models.Article, error) {
	return []*models.Article{}, errors.New("err")
}

func (srm *searchRepositoryMock4) GetArticlesBySearchParameters(ctx context.Context, substringToSearch string, login string, categoryName string, tagName string) ([]*models.Article, error) {
	return []*models.Article{}, errors.New("err")
}

func (srm *searchRepositoryMock4) TagExists(ctx context.Context, tag string) (bool, error) {
	return true, nil
}

func TestGetArticlesByTagNegativeKnownError2(t *testing.T) {
	su := NewSearchUsecase(&searchRepositoryMock4{}, loggerMock)

	res, err := su.GetArticlesByTag(context.Background(), "java")
	assert.NotEqual(t, nil, err)
	if res != nil {
		t.Error()
	}
}
