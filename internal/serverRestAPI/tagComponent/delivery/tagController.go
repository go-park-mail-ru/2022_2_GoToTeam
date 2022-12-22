package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/tagComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/tagComponent/delivery/modelsRestApi"
	"2022_2_GoTo_team/pkg/utils/logger"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TagController struct {
	tagUsecase tagComponentInterfaces.TagUsecaseInterface
	logger     *logger.Logger
}

func NewTagController(tagUsecase tagComponentInterfaces.TagUsecaseInterface, logger *logger.Logger) *TagController {
	logger.LogrusLogger.Debug("Enter to the NewTagController function.")

	tagController := &TagController{
		tagUsecase: tagUsecase,
		logger:     logger,
	}

	logger.LogrusLogger.Info("TagController has created.")

	return tagController
}

func (tc *TagController) TagsListHandler(c echo.Context) error {
	ctx := c.Request().Context()
	tc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the TagsListHandler function.")

	tags, err := tc.tagUsecase.GetTagsList(ctx)
	if err != nil {
		tc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	tagsList := modelsRestApi.TagsList{}
	for _, v := range tags {
		tagsList.TagsNames = append(tagsList.TagsNames, v.TagName)
	}

	tc.logger.LogrusLoggerWithContext(ctx).Debug("Formed tagsList: ", tagsList)

	jsonBytes, err := tagsList.MarshalJSON()
	if err != nil {
		tc.logger.LogrusLoggerWithContext(ctx).Error(err)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}
