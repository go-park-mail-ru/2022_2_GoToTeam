package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/articleComponent/delivery/modelsRestApi"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/articleComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/articleComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ArticleController struct {
	articleUsecase articleComponentInterfaces.ArticleUsecaseInterface
	logger         *logger.Logger
}

func NewArticleController(articleUsecase articleComponentInterfaces.ArticleUsecaseInterface, logger *logger.Logger) *ArticleController {
	logger.LogrusLogger.Debug("Enter to the NewArticleController function.")

	articleController := &ArticleController{
		articleUsecase: articleUsecase,
		logger:         logger,
	}

	logger.LogrusLogger.Info("ArticleController has created.")

	return articleController
}

func (ac *ArticleController) ArticleHandler(c echo.Context) error {
	ac.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the ArticleHandler function.")

	idStr := c.QueryParam("id")
	ac.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed id: %#v", idStr)
	if idStr == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ac.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}
	if id < 1 {
		ac.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(errors.New(fmt.Sprintf("id = %d < 1", id)))
		return c.NoContent(http.StatusBadRequest)
	}

	article, err := ac.articleUsecase.GetArticleById(c.Request().Context(), id)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.ArticleDontExistsError:
			ac.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.NoContent(http.StatusNotFound)
		default:
			ac.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	articleOutput := modelsRestApi.Article{
		Id:           article.ArticleId,
		Title:        article.Title,
		Description:  article.Description,
		Tags:         article.Tags,
		Category:     article.CategoryName,
		Rating:       article.Rating,
		Comments:     article.CommentsCount,
		Content:      article.Content,
		CoverImgPath: article.CoverImgPath,
		Publisher: modelsRestApi.Publisher{
			Username: article.Publisher.Username,
			Login:    article.Publisher.Login,
		},
		CoAuthor: modelsRestApi.CoAuthor{
			Username: article.CoAuthor.Username,
			Login:    article.CoAuthor.Login,
		},
	}
	ac.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Formed articleOutput: ", articleOutput)

	return c.JSON(http.StatusOK, articleOutput)
}
