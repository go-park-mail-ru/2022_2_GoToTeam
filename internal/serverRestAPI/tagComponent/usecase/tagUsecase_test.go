package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/tagComponentInterfaces/mock"
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

func TestGetTagsList(t *testing.T) {
	retTags := []*models.Tag{&models.Tag{TagName: "java"}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tagRepositoryMock := mock.NewMockTagRepositoryInterface(ctrl)
	ctx := context.Background()
	tagRepositoryMock.EXPECT().GetAllTags(ctx).Times(1).Return(retTags, nil)

	tagUsecase := NewTagUsecase(tagRepositoryMock, loggerMock)

	res, err := tagUsecase.GetTagsList(ctx)
	assert.Equal(t, retTags, res)
	assert.Equal(t, nil, err)
}

func TestGetTagsListNegative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tagRepositoryMock := mock.NewMockTagRepositoryInterface(ctrl)
	ctx := context.Background()
	tagRepositoryMock.EXPECT().GetAllTags(ctx).Times(1).Return(nil, errors.New("err"))

	tagUsecase := NewTagUsecase(tagRepositoryMock, loggerMock)

	res, err := tagUsecase.GetTagsList(ctx)
	if res != nil {
		t.Error()
	}
	assert.NotEqual(t, nil, err)
}
