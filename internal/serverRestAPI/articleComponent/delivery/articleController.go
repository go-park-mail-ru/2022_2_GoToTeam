package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/articleComponent/delivery/modelsRestApi/createArticle"
	"2022_2_GoTo_team/internal/serverRestAPI/articleComponent/delivery/modelsRestApi/getArticle"
	"2022_2_GoTo_team/internal/serverRestAPI/articleComponent/delivery/modelsRestApi/removeArticle"
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/articleComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/articleComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
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
		case *usecaseToDeliveryErrors.ArticleDoesntExistError:
			ac.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.NoContent(http.StatusNotFound)
		default:
			ac.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	articleOutput := getArticle.Article{
		Id:           article.ArticleId,
		Title:        article.Title,
		Description:  article.Description,
		Tags:         article.Tags,
		Category:     article.CategoryName,
		Rating:       article.Rating,
		Comments:     article.CommentsCount,
		Content:      article.Content,
		CoverImgPath: article.CoverImgPath,
		Publisher: getArticle.Publisher{
			Username: article.Publisher.Username,
			Login:    article.Publisher.Login,
		},
		CoAuthor: getArticle.CoAuthor{
			Username: article.CoAuthor.Username,
			Login:    article.CoAuthor.Login,
		},
	}
	ac.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Formed articleOutput: ", articleOutput)

	return c.JSON(http.StatusOK, articleOutput)
}

func (ac *ArticleController) CreateArticleHandler(c echo.Context) error {
	ac.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the CreateArticleHandler function.")
	defer c.Request().Body.Close()

	cookie, err := c.Cookie(domain.SESSION_COOKIE_HEADER_NAME)
	if err != nil {
		ac.logger.LogrusLoggerWithContext(c.Request().Context()).Info(err)
		return c.NoContent(http.StatusUnauthorized)
	}

	parsedInputArticle := new(createArticle.Article)
	if err := c.Bind(parsedInputArticle); err != nil {
		ac.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	ac.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed parsedInputArticle: %#v", parsedInputArticle)

	err = ac.articleUsecase.AddArticleBySession(c.Request().Context(), &models.Article{Title: parsedInputArticle.Title, Description: parsedInputArticle.Description, Tags: parsedInputArticle.Tags, CategoryName: parsedInputArticle.Category, CoverImgPath: parsedInputArticle.CoverImgPath, Content: parsedInputArticle.Content, CoAuthor: models.CoAuthor{Login: parsedInputArticle.CoAuthorLogin}}, &models.Session{SessionId: cookie.Value})
	if err != nil {
		ac.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func (ac *ArticleController) RemoveArticleHandler(c echo.Context) error {
	ac.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the RemoveArticleHandler function.")
	defer c.Request().Body.Close()

	parsedInputArticleId := new(removeArticle.ArticleId)
	if err := c.Bind(parsedInputArticleId); err != nil {
		ac.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	ac.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed parsedInputArticleId: %#v", parsedInputArticleId)

	err := ac.articleUsecase.RemoveArticleById(c.Request().Context(), parsedInputArticleId.Id)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.ArticleDoesntExistError:
			ac.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.NoContent(http.StatusNotFound)
		default:
			ac.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	return c.NoContent(http.StatusOK)
}
