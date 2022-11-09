package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/articleComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
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
