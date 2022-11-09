package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/articleComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/articleComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/articleComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/errorsUtils"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"context"
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

func (au *articleUsecase) GetArticleById(ctx context.Context, id int) (*models.Article, error) {
	au.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetArticleById function.")

	wrappingErrorMessage := "error while getting article by id"

	article, err := au.articleRepository.GetArticleById(ctx, id)
	if err != nil {
		switch err {
		case repositoryToUsecaseErrors.ArticleRepositoryArticleDontExistsError:
			au.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.ArticleDontExistsError{Err: err})
		default:
			au.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	return article, nil
}
