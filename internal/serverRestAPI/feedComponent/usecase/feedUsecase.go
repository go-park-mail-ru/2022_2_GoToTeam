package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/feedComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
)

type feedUsecase struct {
	feedRepository feedComponentInterfaces.FeedRepositoryInterface
	logger         *logger.Logger
}

func NewFeedUsecase(feedRepository feedComponentInterfaces.FeedRepositoryInterface, logger *logger.Logger) feedComponentInterfaces.FeedUsecaseInterface {
	feedUsecase := &feedUsecase{
		feedRepository: feedRepository,
		logger:         logger,
	}
	// TODO logger
	feedUsecase.feedRepository.PrintArticles()

	return feedUsecase
}

func (fu *feedUsecase) GetArticles() []*models.Article {
	return fu.feedRepository.GetArticles()
}
