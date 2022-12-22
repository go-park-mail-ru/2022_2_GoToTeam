package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/feedComponentInterfaces/mock"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/domain"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

var loggerMock = &logger.Logger{LogrusLogger: logrus.New().WithFields(logrus.Fields{
	"requestId": "qwerty",
	"userEmail": "asd@asd.asd",
})}

func TestGetFeed(t *testing.T) {
	retArticles := []*models.Article{&models.Article{ArticleId: 1}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	feedRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	feedRepositoryMock.EXPECT().GetFeed(ctx, "asd@asd.asd").Times(1).Return(retArticles, nil)

	feedUsecase := NewFeedUsecase(feedRepositoryMock, loggerMock)

	res, err := feedUsecase.GetFeed(ctx)
	assert.Equal(t, retArticles, res)
	assert.Equal(t, nil, err)
}

func TestGetFeedNegativeEmptyEmailContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	feedRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.Background()
	feedRepositoryMock.EXPECT().GetFeed(ctx, "").Times(1).Return(nil, errors.New("err"))

	feedUsecase := NewFeedUsecase(feedRepositoryMock, loggerMock)

	res, err := feedUsecase.GetFeed(ctx)
	if res != nil {
		t.Error("err")
	}
	assert.NotEqual(t, nil, err)
}

func TestGetFeedForUserByLogin(t *testing.T) {
	retArticles := []*models.Article{&models.Article{ArticleId: 1}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	feedRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	feedRepositoryMock.EXPECT().GetFeedForUserByLogin(ctx, "login123", "asd@asd.asd").Times(1).Return(retArticles, nil)
	feedRepositoryMock.EXPECT().UserExistsByLogin(ctx, "login123").Times(1).Return(true, nil)

	feedUsecase := NewFeedUsecase(feedRepositoryMock, loggerMock)

	res, err := feedUsecase.GetFeedForUserByLogin(ctx, "login123")
	assert.Equal(t, retArticles, res)
	assert.Equal(t, nil, err)
}

func TestGetFeedForUserByLoginNegativeEmptyEmailContext(t *testing.T) {
	retArticles := []*models.Article{&models.Article{ArticleId: 1}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	feedRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.Background()
	feedRepositoryMock.EXPECT().GetFeedForUserByLogin(ctx, "login123", "asd@asd.asd").Times(0).Return(retArticles, nil)
	feedRepositoryMock.EXPECT().UserExistsByLogin(ctx, "login123").Times(0).Return(true, nil)

	feedUsecase := NewFeedUsecase(feedRepositoryMock, loggerMock)

	res, err := feedUsecase.GetFeedForUserByLogin(ctx, "")
	if res != nil {
		t.Error("err")
	}
	assert.NotEqual(t, nil, err)
}

func TestGetFeedForUserByLoginNegativeUserExistsByLoginError(t *testing.T) {
	retArticles := []*models.Article{&models.Article{ArticleId: 1}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	feedRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	feedRepositoryMock.EXPECT().GetFeedForUserByLogin(ctx, "login123", "asd@asd.asd").Times(0).Return(retArticles, nil)
	feedRepositoryMock.EXPECT().UserExistsByLogin(ctx, "login123").Times(1).Return(true, errors.New("err"))

	feedUsecase := NewFeedUsecase(feedRepositoryMock, loggerMock)

	res, err := feedUsecase.GetFeedForUserByLogin(ctx, "login123")
	if res != nil {
		t.Error("err")
	}
	assert.NotEqual(t, nil, err)
}

func TestGetFeedForUserByLoginNegative2UserExistsByLoginError(t *testing.T) {
	retArticles := []*models.Article{&models.Article{ArticleId: 1}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	feedRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	feedRepositoryMock.EXPECT().GetFeedForUserByLogin(ctx, "login123", "asd@asd.asd").Times(0).Return(retArticles, nil)
	feedRepositoryMock.EXPECT().UserExistsByLogin(ctx, "login123").Times(1).Return(false, nil)

	feedUsecase := NewFeedUsecase(feedRepositoryMock, loggerMock)

	res, err := feedUsecase.GetFeedForUserByLogin(ctx, "login123")
	if res != nil {
		t.Error("err")
	}
	assert.NotEqual(t, nil, err)
}

func TestGetFeedForUserByLoginNegativeGetFeedForUserByLogin(t *testing.T) {
	retArticles := []*models.Article{&models.Article{ArticleId: 1}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	feedRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	feedRepositoryMock.EXPECT().GetFeedForUserByLogin(ctx, "login123", "asd@asd.asd").Times(1).Return(retArticles, errors.New("err"))
	feedRepositoryMock.EXPECT().UserExistsByLogin(ctx, "login123").Times(1).Return(true, nil)

	feedUsecase := NewFeedUsecase(feedRepositoryMock, loggerMock)

	res, err := feedUsecase.GetFeedForUserByLogin(ctx, "login123")
	if res != nil {
		t.Error("err")
	}
	assert.NotEqual(t, nil, err)
}

func TestGetGetFeedForCategory(t *testing.T) {
	retArticles := []*models.Article{&models.Article{ArticleId: 1}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	feedRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	feedRepositoryMock.EXPECT().GetFeedForCategory(ctx, "asd", "asd@asd.asd").Times(1).Return(retArticles, nil)
	feedRepositoryMock.EXPECT().CategoryExists(ctx, "asd").Times(1).Return(true, nil)

	feedUsecase := NewFeedUsecase(feedRepositoryMock, loggerMock)

	res, err := feedUsecase.GetFeedForCategory(ctx, "asd")
	assert.Equal(t, retArticles, res)
	assert.Equal(t, nil, err)
}

func TestGetGetFeedForCategoryNegative(t *testing.T) {
	retArticles := []*models.Article{&models.Article{ArticleId: 1}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	feedRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.Background()
	feedRepositoryMock.EXPECT().GetFeedForCategory(ctx, "asd", "asd@asd.asd").Times(0).Return(retArticles, nil)
	feedRepositoryMock.EXPECT().CategoryExists(ctx, "asd").Times(1).Return(false, errors.New("err"))

	feedUsecase := NewFeedUsecase(feedRepositoryMock, loggerMock)

	res, err := feedUsecase.GetFeedForCategory(ctx, "asd")
	if res != nil {
		t.Error("err")
	}
	assert.NotEqual(t, nil, err)
}

func TestGetGetFeedForCategoryNegative2(t *testing.T) {
	retArticles := []*models.Article{&models.Article{ArticleId: 1}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	feedRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	feedRepositoryMock.EXPECT().GetFeedForCategory(ctx, "asd", "asd@asd.asd").Times(0).Return(retArticles, nil)
	feedRepositoryMock.EXPECT().CategoryExists(ctx, "asd").Times(1).Return(false, nil)

	feedUsecase := NewFeedUsecase(feedRepositoryMock, loggerMock)

	res, err := feedUsecase.GetFeedForCategory(ctx, "asd")
	if res != nil {
		t.Error("err")
	}
	assert.NotEqual(t, nil, err)
}

func TestGetGetFeedForCategoryNegative3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	feedRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	feedRepositoryMock.EXPECT().GetFeedForCategory(ctx, "asd", "asd@asd.asd").Times(1).Return(nil, errors.New("err"))
	feedRepositoryMock.EXPECT().CategoryExists(ctx, "asd").Times(1).Return(true, nil)

	feedUsecase := NewFeedUsecase(feedRepositoryMock, loggerMock)

	res, err := feedUsecase.GetFeedForCategory(ctx, "asd")
	if res != nil {
		t.Error("err")
	}
	assert.NotEqual(t, nil, err)
}

func TestGetNewArticlesFromIdForSubscriber(t *testing.T) {
	retArticles := []int{1}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	feedRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.WithValue(context.Background(), domain.USER_EMAIL_KEY_FOR_CONTEXT, "asd@asd.asd")
	feedRepositoryMock.EXPECT().GetNewArticlesFromIdForSubscriber(ctx, 1, "asd@asd.asd").Times(1).Return(retArticles, nil)

	feedUsecase := NewFeedUsecase(feedRepositoryMock, loggerMock)

	res, err := feedUsecase.GetNewArticlesFromIdForSubscriber(ctx, 1)
	assert.Equal(t, retArticles, res)
	assert.Equal(t, nil, err)
}

func TestGetNewArticlesFromIdForSubscriberNegative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	feedRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.Background()
	feedRepositoryMock.EXPECT().GetNewArticlesFromIdForSubscriber(ctx, 1, "").Times(1).Return(nil, errors.New("err"))

	feedUsecase := NewFeedUsecase(feedRepositoryMock, loggerMock)

	res, err := feedUsecase.GetNewArticlesFromIdForSubscriber(ctx, 1)
	if res != nil {
		t.Error("err")
	}
	assert.NotEqual(t, nil, err)
}

/*
func TestGetFeedNegativeEmptyEmailContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	feedRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.Background()
	feedRepositoryMock.EXPECT().GetFeed(ctx, "").Times(1).Return(nil, errors.New("err"))

	feedUsecase := NewFeedUsecase(feedRepositoryMock, loggerMock)

	res, err := feedUsecase.GetFeed(ctx)
	if res != nil {
		t.Error("err")
	}
	assert.NotEqual(t, nil, err)
}

*/
