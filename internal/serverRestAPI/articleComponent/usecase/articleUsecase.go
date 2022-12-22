package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/articleComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/articleComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/articleComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/sessionUtils"
	"2022_2_GoTo_team/pkg/domain"
	"2022_2_GoTo_team/pkg/utils/errorsUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
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

	email, err := sessionUtils.GetEmailFromContext(ctx, au.logger)
	if err != nil {
		au.logger.LogrusLoggerWithContext(ctx).Error(err)
		email = ""
	}

	article, err := au.articleRepository.GetArticleById(ctx, id, email)
	if err != nil {
		switch err {
		case repositoryToUsecaseErrors.ArticleRepositoryArticleDoesntExistError:
			au.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.ArticleDoesntExistError{Err: err})
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
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.ArticleDoesntExistError{Err: errors.New("article doesnt exist")})
	}

	return nil
}

func (au *articleUsecase) AddArticleBySession(ctx context.Context, article *models.Article) error {
	au.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the AddArticleBySession function.")

	wrappingErrorMessage := "error while adding new article by session"

	email := ctx.Value(domain.USER_EMAIL_KEY_FOR_CONTEXT)
	au.logger.LogrusLoggerWithContext(ctx).Debug("Email from context = ", email)

	if email == nil || email.(string) == "" {
		au.logger.LogrusLoggerWithContext(ctx).Error("Email from context is empty.")
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{Err: errors.New("email from context is empty")})
	}
	article.Publisher.Email = email.(string)

	_, err := au.articleRepository.AddArticle(ctx, article)
	if err != nil {
		au.logger.LogrusLoggerWithContext(ctx).Error(err)
		return errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	return nil
}

func (au *articleUsecase) UpdateArticle(ctx context.Context, article *models.Article) error {
	au.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UpdateArticle function.")

	wrappingErrorMessage := "error while updating new article by session"

	email := ctx.Value(domain.USER_EMAIL_KEY_FOR_CONTEXT)
	au.logger.LogrusLoggerWithContext(ctx).Debug("Email from context = ", email)

	if email == nil || email.(string) == "" {
		au.logger.LogrusLoggerWithContext(ctx).Error("Email from context is empty.")
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{Err: errors.New("email from context is empty")})
	}

	authorEmail, err := au.articleRepository.GetAuthorEmailForArticle(ctx, article.ArticleId)
	if err != nil {
		au.logger.LogrusLoggerWithContext(ctx).Error(err)
		return errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	if email != authorEmail {
		au.logger.LogrusLoggerWithContext(ctx).Error("Email is not author fot the article.")
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailIsNotAuthorError{Err: errors.New("email is not author fot the article")})
	}

	err = au.articleRepository.UpdateArticle(ctx, article)
	if err != nil {
		au.logger.LogrusLoggerWithContext(ctx).Error(err)
		return errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	return nil
}

func (au *articleUsecase) ProcessLike(ctx context.Context, likeData *models.LikeData) (int, error) {
	au.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the ProcessLike function.")

	wrappingErrorMessage := "error while liking the article by session"

	email, err := sessionUtils.GetEmailFromContext(ctx, au.logger)
	if err != nil {
		au.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{Err: err})
	}

	sign := likeData.Sign
	if sign == 1 { // Like
		_, err = au.articleRepository.AddLike(ctx, true, likeData.Id, email)
	} else if sign == -1 { // Dislike
		_, err = au.articleRepository.AddLike(ctx, false, likeData.Id, email)
	} else if sign == 0 { // Remove like if exist
		_, err = au.articleRepository.RemoveLike(ctx, likeData.Id, email)
	}
	if err != nil {
		au.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	updatedRating, err := au.articleRepository.GetArticleRating(ctx, likeData.Id)
	if err != nil {
		au.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return updatedRating, err
}
