package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/feedComponentInterfaces/mock"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
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

	tagRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.Background()
	tagRepositoryMock.EXPECT().GetFeed(ctx).Times(1).Return(retArticles, nil)

	feedUsecase := NewFeedUsecase(tagRepositoryMock, loggerMock)

	res, err := feedUsecase.GetFeed(ctx)
	assert.Equal(t, retArticles, res)
	assert.Equal(t, nil, err)
}

func TestGetFeedNegative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tagRepositoryMock := mock.NewMockFeedRepositoryInterface(ctrl)
	ctx := context.Background()
	tagRepositoryMock.EXPECT().GetFeed(ctx).Times(1).Return(nil, errors.New("err"))

	feedUsecase := NewFeedUsecase(tagRepositoryMock, loggerMock)

	res, err := feedUsecase.GetFeed(ctx)
	if res != nil {
		t.Error()
	}
	assert.NotEqual(t, nil, err)
}
