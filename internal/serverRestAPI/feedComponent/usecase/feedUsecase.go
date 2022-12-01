package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/feedComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/feedComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/utils/errorsUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"2022_2_GoTo_team/pkg/utils/validators"
	"context"
	"fmt"
)

type feedUsecase struct {
	feedRepository feedComponentInterfaces.FeedRepositoryInterface
	logger         *logger.Logger
}

func NewFeedUsecase(feedRepository feedComponentInterfaces.FeedRepositoryInterface, logger *logger.Logger) feedComponentInterfaces.FeedUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewFeedUsecase function.")

	feedUsecase := &feedUsecase{
		feedRepository: feedRepository,
		logger:         logger,
	}

	logger.LogrusLogger.Info("feedUsecase has created.")

	return feedUsecase
}

func (fu *feedUsecase) GetFeed(ctx context.Context) ([]*models.Article, error) {
	fu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetFeed function.")

	wrappingErrorMessage := "error while getting articles"

	articles, err := fu.feedRepository.GetFeed(ctx)
	if err != nil {
		fu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return articles, nil
}

func (fu *feedUsecase) GetFeedForUserByLogin(ctx context.Context, login string) ([]*models.Article, error) {
	fu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetFeedForUserByLogin function.")

	wrappingErrorMessage := "error while getting articles for user by login"

	if !validators.LoginIsValidByRegExp(login) {
		fu.logger.LogrusLoggerWithContext(ctx).Infof("Login %s is not valid.", login)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.LoginIsNotValidError{Err: fmt.Errorf("login is not valid %#v", login)})
	}

	exists, err := fu.feedRepository.UserExistsByLogin(ctx, login)
	if err != nil {
		fu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}
	if !exists {
		fu.logger.LogrusLoggerWithContext(ctx).Infof("Login %s doesnt exist", login)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.LoginDoesntExistError{Err: fmt.Errorf("login %#v doesnt exist", login)})
	}

	articles, err := fu.feedRepository.GetFeedForUserByLogin(ctx, login)
	if err != nil {
		fu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return articles, nil
}

func (fu *feedUsecase) GetFeedForCategory(ctx context.Context, category string) ([]*models.Article, error) {
	fu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetFeedForCategory function.")

	wrappingErrorMessage := "error while getting articles for category"

	exists, err := fu.feedRepository.CategoryExists(ctx, category)
	if err != nil {
		fu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}
	if !exists {
		fu.logger.LogrusLoggerWithContext(ctx).Infof("Category %s doesnt exists", category)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.CategoryDoesntExistError{Err: fmt.Errorf("category %#v doesnt exist", category)})
	}

	articles, err := fu.feedRepository.GetFeedForCategory(ctx, category)
	if err != nil {
		fu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return articles, nil
}
