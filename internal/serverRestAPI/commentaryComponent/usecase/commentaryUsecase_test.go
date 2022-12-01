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

type commentaryRepositoryMock struct {
}

func (crm *commentaryRepositoryMock) AddCommentaryByEmail(ctx context.Context, commentary *models.Commentary) (int, error) {
	return 0, nil
}

func (crm *commentaryRepositoryMock) GetAllCommentsForArticle(ctx context.Context, articleId int) ([]*models.Commentary, error) {
	return []*models.Commentary{}, nil
}

func TestAddCommentaryBySession(t *testing.T) {
	cu := NewCommentaryUsecase(&commentaryRepositoryMock{}, loggerMock)

	err := cu.AddCommentaryBySession(context.Background(), &models.Commentary{})
	assert.NotEqual(t, nil, err)
}

func TestGetAllCommentariesForArticle(t *testing.T) {
	cu := NewCommentaryUsecase(&commentaryRepositoryMock{}, loggerMock)

	res, err := cu.GetAllCommentariesForArticle(context.Background(), 2)
	assert.NotEqual(t, nil, res)
	assert.Equal(t, nil, err)
}

type commentaryRepositoryMock2 struct {
}

func (crm *commentaryRepositoryMock2) AddCommentaryByEmail(ctx context.Context, commentary *models.Commentary) (int, error) {
	return 0, errors.New("err")
}

func (crm *commentaryRepositoryMock2) GetAllCommentsForArticle(ctx context.Context, articleId int) ([]*models.Commentary, error) {
	return []*models.Commentary{}, errors.New("err")
}

func TestAddCommentaryBySessionNegative(t *testing.T) {
	cu := NewCommentaryUsecase(&commentaryRepositoryMock2{}, loggerMock)

	err := cu.AddCommentaryBySession(context.Background(), &models.Commentary{})
	assert.NotEqual(t, nil, err)
}

func TestGetAllCommentariesForArticleNegative(t *testing.T) {
	cu := NewCommentaryUsecase(&commentaryRepositoryMock2{}, loggerMock)

	_, err := cu.GetAllCommentariesForArticle(context.Background(), 2)
	assert.NotEqual(t, nil, err)
}
