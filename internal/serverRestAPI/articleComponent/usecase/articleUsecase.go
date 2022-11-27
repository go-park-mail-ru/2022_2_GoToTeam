package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/articleComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/articleComponentErrors/usecaseToDeliveryErrors"
	repositoryToUsecaseErrors_sessionComponent "2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/sessionComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/articleComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/sessionComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/errorsUtils"
	"2022_2_GoTo_team/pkg/logger"
	"context"
	"errors"
)

type articleUsecase struct {
	articleRepository articleComponentInterfaces.ArticleRepositoryInterface
	sessionRepository sessionComponentInterfaces.SessionRepositoryInterface
	logger            *logger.Logger
}

func NewArticleUsecase(articleRepository articleComponentInterfaces.ArticleRepositoryInterface, sessionRepository sessionComponentInterfaces.SessionRepositoryInterface, logger *logger.Logger) articleComponentInterfaces.ArticleUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewArticleUsecase function.")

	articleUsecase := &articleUsecase{
		articleRepository: articleRepository,
		sessionRepository: sessionRepository,
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

func (au *articleUsecase) RemoveArticleById(ctx context.Context, id int) error {
	au.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the RemoveArticleById function.")

	wrappingErrorMessage := "error while getting article by id"

	removedRowsCount, err := au.articleRepository.DeleteArticleById(ctx, id)
	if err != nil {
		au.logger.LogrusLoggerWithContext(ctx).Error(err)
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})

	}
	if removedRowsCount <= 0 {
		au.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.ArticleDontExistsError{Err: errors.New("article dont exists")})
	}

	return nil
}

func (au *articleUsecase) AddArticleBySession(ctx context.Context, article *models.Article, session *models.Session) error {
	au.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the AddArticleBySession function.")

	wrappingErrorMessage := "error while adding new article by session"

	authorEmail, err := au.sessionRepository.GetEmailBySession(ctx, session)
	if err != nil {
		switch err {
		case repositoryToUsecaseErrors_sessionComponent.SessionRepositoryEmailDontExistsError:
			au.logger.LogrusLoggerWithContext(ctx).Error(err)
			return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDontFoundError{Err: err})
		default:
			au.logger.LogrusLoggerWithContext(ctx).Error(err)
			return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}
	article.Publisher.Email = authorEmail

	_, err = au.articleRepository.AddArticle(ctx, article)
	if err != nil {
		au.logger.LogrusLoggerWithContext(ctx).Error(err)
		return errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	return nil
}
