package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/feedComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/feedComponent/delivery/modelsRestApi"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

const ARTICLE_NUMBER_IN_FEED = 10

type FeedController struct {
	feedUsecase feedComponentInterfaces.FeedUsecaseInterface
	logger      *logger.Logger
}

func NewFeedController(feedUsecase feedComponentInterfaces.FeedUsecaseInterface, logger *logger.Logger) *FeedController {
	feedController := &FeedController{
		feedUsecase: feedUsecase,
		logger:      logger,
	}

	return feedController
}

func (fc *FeedController) FeedHandler(c echo.Context) error {
	startFromArticleOfNumberStr := c.QueryParam("startFromArticleOfNumber")
	if startFromArticleOfNumberStr == "" {
		startFromArticleOfNumberStr = "0"
	}

	startFromArticleOfNumber, err := strconv.Atoi(startFromArticleOfNumberStr)
	if err != nil {
		//c.LogrusLogger().Printf("Error: %s", err.Error())
		fc.logger.LogrusLogger.Error(err)
		return c.NoContent(http.StatusBadRequest)
	}
	if startFromArticleOfNumber < 0 {
		//c.LogrusLogger().Printf("Error: startFromArticleOfNumber = %d < 0", startFromArticleOfNumber)
		fc.logger.LogrusLogger.Error("startFromArticleOfNumber = ", startFromArticleOfNumber, " < 0")
		return c.NoContent(http.StatusBadRequest)
	}

	articles := fc.feedUsecase.GetArticles()

	if startFromArticleOfNumber+ARTICLE_NUMBER_IN_FEED <= len(articles) {
		articles = articles[startFromArticleOfNumber : startFromArticleOfNumber+ARTICLE_NUMBER_IN_FEED]
	} else if startFromArticleOfNumber < len(articles) {
		articles = articles[startFromArticleOfNumber:]
	} else {
		var startTmp = len(articles) - ARTICLE_NUMBER_IN_FEED
		if startTmp < 0 {
			startTmp = 0
		}
		articles = articles[startTmp:]
	}

	feed := modelsRestApi.Feed{}
	for _, v := range articles {
		article := modelsRestApi.Article{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			Tags:        v.Tags,
			Category:    v.Category,
			Rating:      v.Rating,
			Authors:     v.Authors,
			Content:     v.Content,
		}
		feed.Articles = append(feed.Articles, article)
	}
	fc.logger.LogrusLogger.Info("Formed feed = ", feed)

	return c.JSON(http.StatusOK, feed)
}
