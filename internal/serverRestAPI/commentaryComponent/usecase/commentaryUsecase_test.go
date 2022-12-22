package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	domainPkg "2022_2_GoTo_team/pkg/domain"
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

func (crm *commentaryRepositoryMock) GetAllCommentsForArticle(ctx context.Context, articleId int, email string) ([]*models.Commentary, error) {
	return []*models.Commentary{}, nil
}

func (crm *commentaryRepositoryMock) AddLike(ctx context.Context, isLike bool, commentId int, email string) (int, error) {
	return 0, nil
}

func (crm *commentaryRepositoryMock) RemoveLike(ctx context.Context, commentId int, email string) (int64, error) {
	return 1, nil
}

func (crm *commentaryRepositoryMock) GetCommentaryRating(ctx context.Context, commentaryId int) (int, error) {
	return 55, nil
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

func TestProcessLike(t *testing.T) {
	cu := NewCommentaryUsecase(&commentaryRepositoryMock{}, loggerMock)
	ctx := context.WithValue(context.Background(), domainPkg.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	_, err := cu.ProcessLike(ctx, &models.LikeData{Id: 1, Sign: 1})
	assert.Equal(t, nil, err)
}

func TestProcessLike2(t *testing.T) {
	cu := NewCommentaryUsecase(&commentaryRepositoryMock{}, loggerMock)
	ctx := context.WithValue(context.Background(), domainPkg.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	_, err := cu.ProcessLike(ctx, &models.LikeData{Id: 1, Sign: 0})
	assert.Equal(t, nil, err)
}

func TestProcessLike3(t *testing.T) {
	cu := NewCommentaryUsecase(&commentaryRepositoryMock{}, loggerMock)
	ctx := context.WithValue(context.Background(), domainPkg.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	_, err := cu.ProcessLike(ctx, &models.LikeData{Id: 1, Sign: -1})
	assert.Equal(t, nil, err)
}

type commentaryRepositoryMock2 struct {
}

func (crm *commentaryRepositoryMock2) AddCommentaryByEmail(ctx context.Context, commentary *models.Commentary) (int, error) {
	return 0, errors.New("err")
}

func (crm *commentaryRepositoryMock2) GetAllCommentsForArticle(ctx context.Context, articleId int, email string) ([]*models.Commentary, error) {
	return []*models.Commentary{}, errors.New("err")
}

func (crm *commentaryRepositoryMock2) AddLike(ctx context.Context, isLike bool, commentId int, email string) (int, error) {
	return 0, nil
}

func (crm *commentaryRepositoryMock2) RemoveLike(ctx context.Context, commentId int, email string) (int64, error) {
	return 1, nil
}

func (crm *commentaryRepositoryMock2) GetCommentaryRating(ctx context.Context, commentaryId int) (int, error) {
	return 55, nil
}

func TestAddCommentaryBySessionNegative(t *testing.T) {
	cu := NewCommentaryUsecase(&commentaryRepositoryMock2{}, loggerMock)
	email := "asd@asd.asd"
	ctx := context.WithValue(context.Background(), domainPkg.USER_EMAIL_KEY_FOR_CONTEXT, email)
	err := cu.AddCommentaryBySession(ctx, &models.Commentary{Publisher: models.Publisher{Email: email}})
	assert.NotEqual(t, nil, err)
}

func TestGetAllCommentariesForArticleNegative(t *testing.T) {
	cu := NewCommentaryUsecase(&commentaryRepositoryMock2{}, loggerMock)

	_, err := cu.GetAllCommentariesForArticle(context.Background(), 2)
	assert.NotEqual(t, nil, err)
}

func TestProcessLikeNegativeNoEmailContext(t *testing.T) {
	cu := NewCommentaryUsecase(&commentaryRepositoryMock2{}, loggerMock)
	_, err := cu.ProcessLike(context.Background(), &models.LikeData{Id: 1, Sign: -1})
	assert.NotEqual(t, nil, err)
}

type commentaryRepositoryMock3 struct {
}

func (crm *commentaryRepositoryMock3) AddCommentaryByEmail(ctx context.Context, commentary *models.Commentary) (int, error) {
	return 0, nil
}

func (crm *commentaryRepositoryMock3) GetAllCommentsForArticle(ctx context.Context, articleId int, email string) ([]*models.Commentary, error) {
	return []*models.Commentary{}, errors.New("err")
}

func (crm *commentaryRepositoryMock3) AddLike(ctx context.Context, isLike bool, commentId int, email string) (int, error) {
	return 0, nil
}

func (crm *commentaryRepositoryMock3) RemoveLike(ctx context.Context, commentId int, email string) (int64, error) {
	return 1, errors.New("err")
}

func (crm *commentaryRepositoryMock3) GetCommentaryRating(ctx context.Context, commentaryId int) (int, error) {
	return 55, errors.New("err")
}

func TestAddCommentaryBySessionOk(t *testing.T) {
	cu := NewCommentaryUsecase(&commentaryRepositoryMock3{}, loggerMock)
	email := "asd@asd.asd"
	ctx := context.WithValue(context.Background(), domainPkg.USER_EMAIL_KEY_FOR_CONTEXT, email)
	err := cu.AddCommentaryBySession(ctx, &models.Commentary{Publisher: models.Publisher{Email: email}})
	assert.Equal(t, nil, err)
}

func TestProcessNegative(t *testing.T) {
	cu := NewCommentaryUsecase(&commentaryRepositoryMock3{}, loggerMock)
	ctx := context.WithValue(context.Background(), domainPkg.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	_, err := cu.ProcessLike(ctx, &models.LikeData{Id: 1, Sign: 0})
	assert.NotEqual(t, nil, err)
}

func TestProcessNegative2(t *testing.T) {
	cu := NewCommentaryUsecase(&commentaryRepositoryMock3{}, loggerMock)
	ctx := context.WithValue(context.Background(), domainPkg.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	_, err := cu.ProcessLike(ctx, &models.LikeData{Id: 1, Sign: -1})
	assert.NotEqual(t, nil, err)
}
