package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/commentaryComponent/delivery/modelsRestApi/createCommentary"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/commentaryComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/logger"
	"github.com/labstack/echo/v4"
	"net/http"
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
