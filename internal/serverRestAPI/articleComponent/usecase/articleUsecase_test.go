package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/articleComponentErrors/repositoryToUsecaseErrors"
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

type articleRepositoryMock struct {
}

func (arm *articleRepositoryMock) GetArticleById(ctx context.Context, id int, email string) (*models.Article, error) {
	return &models.Article{ArticleId: id}, nil
}

func (arm *articleRepositoryMock) GetTagsForArticle(ctx context.Context, articleId int) ([]string, error) {
	return []string{}, nil
}

func (arm *articleRepositoryMock) AddArticle(ctx context.Context, article *models.Article) (int, error) {
	return 123, nil
}

func (arm *articleRepositoryMock) DeleteArticleById(ctx context.Context, articleId int) (int64, error) {
	return 1, nil
}

func (arm *articleRepositoryMock) UpdateArticle(ctx context.Context, article *models.Article) error {
	return nil
}

func (arm *articleRepositoryMock) GetAuthorEmailForArticle(ctx context.Context, articleId int) (string, error) {
	return "qwe@qwe.qwe", nil
}

func (arm *articleRepositoryMock) AddLike(ctx context.Context, isLike bool, articleId int, email string) (int, error) {
	return 3, nil
}

func (arm *articleRepositoryMock) RemoveLike(ctx context.Context, articleId int, email string) (int64, error) {
	return 3, nil
}

func (arm *articleRepositoryMock) GetArticleRating(ctx context.Context, articleId int) (int, error) {
	return 55, nil
}

func TestGetArticleById(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock{}, loggerMock)

	id := 13
	res, err := au.GetArticleById(context.Background(), id)
	assert.Equal(t, id, res.ArticleId)
	assert.Equal(t, nil, err)
}

func TestRemoveArticleById(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock{}, loggerMock)

	id := 13
	err := au.RemoveArticleById(context.Background(), id)
	assert.Equal(t, nil, err)
}

func TestAddArticleBySession(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock{}, loggerMock)

	err := au.AddArticleBySession(context.Background(), &models.Article{})
	assert.NotEqual(t, nil, err)
}

func TestUpdateArticle(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock{}, loggerMock)

	err := au.UpdateArticle(context.Background(), &models.Article{})
	assert.NotEqual(t, nil, err)
}

func TestProcessLike(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock{}, loggerMock)

	_, err := au.ProcessLike(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), &models.LikeData{Id: 1, Sign: 1})
	assert.Equal(t, nil, err)
}

func TestProcessLike2(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock{}, loggerMock)

	_, err := au.ProcessLike(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), &models.LikeData{Id: 1, Sign: -1})
	assert.Equal(t, nil, err)
}

func TestProcessLike3(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock{}, loggerMock)

	_, err := au.ProcessLike(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), &models.LikeData{Id: 1, Sign: 0})
	assert.Equal(t, nil, err)
}

type articleRepositoryMock2 struct {
}

func (arm *articleRepositoryMock2) GetArticleById(ctx context.Context, id int, email string) (*models.Article, error) {
	return nil, errors.New("unknown err")
}

func (arm *articleRepositoryMock2) GetTagsForArticle(ctx context.Context, articleId int) ([]string, error) {
	return []string{}, nil
}

func (arm *articleRepositoryMock2) AddArticle(ctx context.Context, article *models.Article) (int, error) {
	return 123, errors.New("unknown err")
}

func (arm *articleRepositoryMock2) DeleteArticleById(ctx context.Context, articleId int) (int64, error) {
	return 1, errors.New("unknown err")
}

func (arm *articleRepositoryMock2) UpdateArticle(ctx context.Context, article *models.Article) error {
	return errors.New("unknown err")
}

func (arm *articleRepositoryMock2) GetAuthorEmailForArticle(ctx context.Context, articleId int) (string, error) {
	return "asd@asd.asd", nil
}

func (arm *articleRepositoryMock2) AddLike(ctx context.Context, isLike bool, articleId int, email string) (int, error) {
	return 3, nil
}

func (arm *articleRepositoryMock2) RemoveLike(ctx context.Context, articleId int, email string) (int64, error) {
	return 3, nil
}

func (arm *articleRepositoryMock2) GetArticleRating(ctx context.Context, articleId int) (int, error) {
	return 55, nil
}

func TestGetArticleByIdNegativeUnknownError(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock2{}, loggerMock)

	id := 13
	_, err := au.GetArticleById(context.Background(), id)
	assert.NotEqual(t, nil, err)
}

func TestRemoveArticleByIdNegative(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock2{}, loggerMock)

	id := 13
	err := au.RemoveArticleById(context.Background(), id)
	assert.NotEqual(t, nil, err)
}

func TestAddArticleBySessionNegative(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock2{}, loggerMock)

	err := au.AddArticleBySession(context.Background(), &models.Article{})
	assert.NotEqual(t, nil, err)
}

func TestUpdateArticleNegative(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock2{}, loggerMock)

	err := au.UpdateArticle(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), &models.Article{})
	assert.NotEqual(t, nil, err)
}

func TestProcessLikeNoEmailInContext(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock2{}, loggerMock)

	_, err := au.ProcessLike(context.Background(), &models.LikeData{Id: 1, Sign: 0})
	assert.NotEqual(t, nil, err)
}

type articleRepositoryMock3 struct {
}

func (arm *articleRepositoryMock3) GetArticleById(ctx context.Context, id int, email string) (*models.Article, error) {
	return nil, repositoryToUsecaseErrors.ArticleRepositoryArticleDoesntExistError
}

func (arm *articleRepositoryMock3) GetTagsForArticle(ctx context.Context, articleId int) ([]string, error) {
	return []string{}, nil
}

func (arm *articleRepositoryMock3) AddArticle(ctx context.Context, article *models.Article) (int, error) {
	return 123, errors.New("unknown err")
}

func (arm *articleRepositoryMock3) DeleteArticleById(ctx context.Context, articleId int) (int64, error) {
	return -1, nil
}

func (arm *articleRepositoryMock3) UpdateArticle(ctx context.Context, article *models.Article) error {
	return nil
}

func (arm *articleRepositoryMock3) GetAuthorEmailForArticle(ctx context.Context, articleId int) (string, error) {
	return "qwe@qwe.qwe", nil
}

func (arm *articleRepositoryMock3) AddLike(ctx context.Context, isLike bool, articleId int, email string) (int, error) {
	return 3, nil
}

func (arm *articleRepositoryMock3) RemoveLike(ctx context.Context, articleId int, email string) (int64, error) {
	return 3, errors.New("err")
}

func (arm *articleRepositoryMock3) GetArticleRating(ctx context.Context, articleId int) (int, error) {
	return 55, nil
}

func TestGetArticleByIdNegativeUnknownError2(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock3{}, loggerMock)

	id := 13
	_, err := au.GetArticleById(context.Background(), id)
	assert.NotEqual(t, nil, err)
}

func TestRemoveArticleByIdNegative2(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock3{}, loggerMock)

	id := 13
	err := au.RemoveArticleById(context.Background(), id)
	assert.NotEqual(t, nil, err)
}

func TestAddArticleBySessionNegative2(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock3{}, loggerMock)

	err := au.AddArticleBySession(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), &models.Article{})
	assert.NotEqual(t, nil, err)
}

func TestUpdateArticleNegative2(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock3{}, loggerMock)

	err := au.UpdateArticle(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), &models.Article{})
	assert.NotEqual(t, nil, err)
}

func TestProcessLikeNegative(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock3{}, loggerMock)

	_, err := au.ProcessLike(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), &models.LikeData{Id: 1, Sign: 0})
	assert.NotEqual(t, nil, err)
}

type articleRepositoryMock4 struct {
}

func (arm *articleRepositoryMock4) GetArticleById(ctx context.Context, id int, email string) (*models.Article, error) {
	return nil, repositoryToUsecaseErrors.ArticleRepositoryArticleDoesntExistError
}

func (arm *articleRepositoryMock4) GetTagsForArticle(ctx context.Context, articleId int) ([]string, error) {
	return []string{}, nil
}

func (arm *articleRepositoryMock4) AddArticle(ctx context.Context, article *models.Article) (int, error) {
	return 123, nil
}

func (arm *articleRepositoryMock4) DeleteArticleById(ctx context.Context, articleId int) (int64, error) {
	return -1, nil
}

func (arm *articleRepositoryMock4) UpdateArticle(ctx context.Context, article *models.Article) error {
	return nil
}

func (arm *articleRepositoryMock4) GetAuthorEmailForArticle(ctx context.Context, articleId int) (string, error) {
	return "qwe@qwe.qwe", nil
}

func (arm *articleRepositoryMock4) AddLike(ctx context.Context, isLike bool, articleId int, email string) (int, error) {
	return 3, nil
}

func (arm *articleRepositoryMock4) RemoveLike(ctx context.Context, articleId int, email string) (int64, error) {
	return 3, nil
}

func (arm *articleRepositoryMock4) GetArticleRating(ctx context.Context, articleId int) (int, error) {
	return 55, errors.New("err")
}

func TestAddArticleBySession2(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock4{}, loggerMock)

	err := au.AddArticleBySession(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), &models.Article{})
	assert.Equal(t, nil, err)
}

func TestUpdateArticle2(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock4{}, loggerMock)

	err := au.UpdateArticle(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "qwe@qwe.qwe"), &models.Article{})
	assert.Equal(t, nil, err)
}

func TestProcessLikeNegative2(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock4{}, loggerMock)

	_, err := au.ProcessLike(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), &models.LikeData{Id: 1, Sign: 0})
	assert.NotEqual(t, nil, err)
}

type articleRepositoryMock5 struct {
}

func (arm *articleRepositoryMock5) GetArticleById(ctx context.Context, id int, email string) (*models.Article, error) {
	return nil, repositoryToUsecaseErrors.ArticleRepositoryArticleDoesntExistError
}

func (arm *articleRepositoryMock5) GetTagsForArticle(ctx context.Context, articleId int) ([]string, error) {
	return []string{}, nil
}

func (arm *articleRepositoryMock5) AddArticle(ctx context.Context, article *models.Article) (int, error) {
	return 123, nil
}

func (arm *articleRepositoryMock5) DeleteArticleById(ctx context.Context, articleId int) (int64, error) {
	return -1, nil
}

func (arm *articleRepositoryMock5) UpdateArticle(ctx context.Context, article *models.Article) error {
	return nil
}

func (arm *articleRepositoryMock5) GetAuthorEmailForArticle(ctx context.Context, articleId int) (string, error) {
	return "", errors.New("err")
}

func (arm *articleRepositoryMock5) AddLike(ctx context.Context, isLike bool, articleId int, email string) (int, error) {
	return 3, nil
}

func (arm *articleRepositoryMock5) RemoveLike(ctx context.Context, articleId int, email string) (int64, error) {
	return 3, nil
}

func (arm *articleRepositoryMock5) GetArticleRating(ctx context.Context, articleId int) (int, error) {
	return 55, nil
}

func TestUpdateArticleNegative3(t *testing.T) {
	au := NewArticleUsecase(&articleRepositoryMock5{}, loggerMock)

	err := au.UpdateArticle(context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd"), &models.Article{})
	assert.NotEqual(t, nil, err)
}
