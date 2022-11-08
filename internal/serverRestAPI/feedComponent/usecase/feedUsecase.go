package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/feedComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/feedComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/errorsUtils"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"context"
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

	logger.LogrusLogger.Info("FeedUsecase has created.")

	return feedUsecase
}

func (fu *feedUsecase) GetFeed(ctx context.Context) ([]*models.Article, error) {
	fu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetFeed function.")

	articles, err := fu.feedRepository.GetFeed(ctx)
	if err != nil {
		fu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, errorsUtils.WrapError("error while getting articles", &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return articles, nil
}

func (fu *feedUsecase) GetFeedForUserByLogin(ctx context.Context, login string) ([]*models.Article, error) {
	fu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetFeedForUserByLogin function.")

	// TODO login validation

	articles, err := fu.feedRepository.GetFeedForUserByLogin(ctx, login)
	if err != nil {
		fu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, errorsUtils.WrapError("error while getting articles", &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return articles, nil
}
