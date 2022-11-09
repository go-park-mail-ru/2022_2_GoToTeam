package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/articleComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
)

type articleUsecase struct {
	articleRepository articleComponentInterfaces.ArticleRepositoryInterface
	logger            *logger.Logger
}

func NewArticleUsecase(articleRepository articleComponentInterfaces.ArticleRepositoryInterface, logger *logger.Logger) articleComponentInterfaces.ArticleUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewArticleUsecase function.")

	articleUsecase := &articleUsecase{
		articleRepository: articleRepository,
		logger:            logger,
	}

	logger.LogrusLogger.Info("articleUsecase has created.")

	return articleUsecase
}
