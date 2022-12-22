package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/commentaryComponent/delivery/modelsRestApi/createCommentary"
	"2022_2_GoTo_team/internal/serverRestAPI/commentaryComponent/delivery/modelsRestApi/getAllCommentariesForArticle"
	"2022_2_GoTo_team/internal/serverRestAPI/commentaryComponent/delivery/modelsRestApi/likeData"
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
	ctx := c.Request().Context()
	cc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the CreateCommentaryHandler function.")
	defer c.Request().Body.Close()

	parsedInputCommentary := new(createCommentary.Commentary)
	if err := c.Bind(parsedInputCommentary); err != nil {
		cc.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	cc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed parsedInputCommentary: %#v", parsedInputCommentary)

	err := cc.commentaryUsecase.AddCommentaryBySession(ctx, &models.Commentary{
		Content:             parsedInputCommentary.Content,
		ArticleId:           parsedInputCommentary.ArticleId,
		CommentForCommentId: parsedInputCommentary.CommentForCommentId,
	})
	if err != nil {
		cc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func (cc *CommentaryController) GetAllCommentariesForArticle(c echo.Context) error {
	ctx := c.Request().Context()
	cc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetAllCommentariesForArticle function.")

	parsedInputArticleIdStr := c.QueryParam("article")
	cc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed parsedInputArticleIdStr: %#v", parsedInputArticleIdStr)
	if parsedInputArticleIdStr == "" {
		parsedInputArticleIdStr = "0"
	}
	parsedInputArticleId, err := strconv.Atoi(parsedInputArticleIdStr)
	if err != nil {
		cc.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}
	if parsedInputArticleId <= 0 {
		cc.logger.LogrusLoggerWithContext(ctx).Warn(errors.New(fmt.Sprintf("parsedInputArticleId = %d <= 0", parsedInputArticleId)))
		return c.NoContent(http.StatusBadRequest)
	}

	cc.logger.LogrusLoggerWithContext(ctx).Debugf("Converted parsedInputArticleId: %#v", parsedInputArticleId)

	commentaries, err := cc.commentaryUsecase.GetAllCommentariesForArticle(ctx, parsedInputArticleId)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		default:
			cc.logger.LogrusLoggerWithContext(ctx).Error(err)
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
			Liked: v.Liked,
		}
		allCommentariesForArticle.Commentaries = append(allCommentariesForArticle.Commentaries, commentary)
	}
	cc.logger.LogrusLoggerWithContext(ctx).Debug("Formed allCommentariesForArticle: ", allCommentariesForArticle)

	jsonBytes, err := allCommentariesForArticle.MarshalJSON()
	if err != nil {
		cc.logger.LogrusLoggerWithContext(ctx).Error(err)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (cc *CommentaryController) LikeHandler(c echo.Context) error {
	ctx := c.Request().Context()
	cc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the LikeHandler function.")
	defer c.Request().Body.Close()

	parsedInputLikeData := new(likeData.LikeData)
	if err := c.Bind(parsedInputLikeData); err != nil {
		cc.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}
	sign := parsedInputLikeData.Sign
	if sign != -1 && sign != 0 && sign != 1 {
		cc.logger.LogrusLoggerWithContext(ctx).Warnf("Incorrect sign value = %#v, should be -1 or 0 or 1", sign)
		return c.NoContent(http.StatusBadRequest)
	}

	cc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed parsedInputLikeData: %#v", parsedInputLikeData)

	updatedRating, err := cc.commentaryUsecase.ProcessLike(ctx, &models.LikeData{Id: parsedInputLikeData.Id, Sign: parsedInputLikeData.Sign})
	if err != nil {
		cc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	likeResponse := likeData.LikeResponse{
		Rating: updatedRating,
	}

	jsonBytes, err := likeResponse.MarshalJSON()
	if err != nil {
		cc.logger.LogrusLoggerWithContext(ctx).Error(err)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}
