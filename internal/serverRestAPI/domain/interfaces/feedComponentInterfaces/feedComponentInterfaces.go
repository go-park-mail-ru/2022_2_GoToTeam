package feedComponentInterfaces

import "2022_2_GoTo_team/internal/serverRestAPI/domain/models"

type FeedUsecaseInterface interface {
	GetArticles() []*models.Article
}

type FeedRepositoryInterface interface {
	PrintArticles()
	GetArticles() []*models.Article
}
