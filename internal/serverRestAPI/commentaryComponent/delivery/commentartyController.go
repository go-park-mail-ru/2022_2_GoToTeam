package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/commentaryComponent/delivery/modelsRestApi/createCommentary"
	"2022_2_GoTo_team/internal/serverRestAPI/commentaryComponent/delivery/modelsRestApi/getAllCommentariesForArticle"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/commentaryComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CommentaryController struct {
	commentaryUsecase commentaryComponentInterfaces.CommentaryUsecaseInterface
	logger            *logger.Logger
}

func NewCommentaryController(commentaryUsecase commentaryComponentInterfaces.CommentaryUsecaseInterface, logger *logger.Logger) *CommentaryController {
	logger.LogrusLogger.Debug("Enter to the NewCommentaryController function.")

	commentaryController := &CommentaryController{
		commentaryUsecase: commentaryUsecase,
		logger:            logger,
	}

	logger.LogrusLogger.Info("CommentaryController has created.")

	return commentaryController
}

func (cc *CommentaryController) CreateCommentaryHandler(c echo.Context) error {
	cc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the CreateCommentaryHandler function.")
	defer c.Request().Body.Close()

	parsedInputCommentary := new(createCommentary.Commentary)
	if err := c.Bind(parsedInputCommentary); err != nil {
		cc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	cc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed parsedInputCommentary: %#v", parsedInputCommentary)

	err := cc.commentaryUsecase.AddCommentaryBySession(c.Request().Context(), &models.Commentary{
		Content:             parsedInputCommentary.Content,
		ArticleId:           parsedInputCommentary.ArticleId,
		CommentForCommentId: parsedInputCommentary.CommentForCommentId,
	})
	if err != nil {
		cc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func (cc *CommentaryController) GetAllCommentariesForArticle(c echo.Context) error {
	cc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the GetAllCommentariesForArticle function.")

	parsedInputArticleIdStr := c.QueryParam("article")
	cc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed parsedInputArticleIdStr: %#v", parsedInputArticleIdStr)
	if parsedInputArticleIdStr == "" {
		parsedInputArticleIdStr = "0"
	}
	parsedInputArticleId, err := strconv.Atoi(parsedInputArticleIdStr)
	if err != nil {
		cc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}
	if parsedInputArticleId <= 0 {
		cc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(errors.New(fmt.Sprintf("parsedInputArticleId = %d <= 0", parsedInputArticleId)))
		return c.NoContent(http.StatusBadRequest)
	}

	cc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Converted parsedInputArticleId: %#v", parsedInputArticleId)

	commentaries, err := cc.commentaryUsecase.GetAllCommentariesForArticle(c.Request().Context(), parsedInputArticleId)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		default:
			cc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	allCommentariesForArticle := getAllCommentariesForArticle.CommentariesForArticle{}
	for _, v := range commentaries {
		commentary := getAllCommentariesForArticle.Commentary{
			CommentId:           v.CommentId,
			Content:             v.Content,
			Rating:              v.Rating,
			ArticleId:           v.ArticleId,
			CommentForCommentId: v.CommentForCommentId,
			Publisher: getAllCommentariesForArticle.Publisher{
				Username: v.Publisher.Username,
				Login:    v.Publisher.Login,
			},
		}
		allCommentariesForArticle.Commentaries = append(allCommentariesForArticle.Commentaries, commentary)
	}
	cc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Formed allCommentariesForArticle: ", allCommentariesForArticle)

	return c.JSON(http.StatusOK, allCommentariesForArticle)
}
