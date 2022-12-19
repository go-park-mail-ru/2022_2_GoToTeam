package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/feedComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/feedComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/feedComponent/delivery/modelsRestApi"
	"2022_2_GoTo_team/pkg/utils/logger"
	"errors"
	"fmt"
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
	logger.LogrusLogger.Debug("Enter to the NewFeedController function.")

	feedController := &FeedController{
		feedUsecase: feedUsecase,
		logger:      logger,
	}

	logger.LogrusLogger.Info("FeedController has created.")

	return feedController
}

func (fc *FeedController) FeedHandler(c echo.Context) error {
	fc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the FeedHandler function.")

	startFromArticleOfNumberStr := c.QueryParam("startFromArticleOfNumber")
	fc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed startFromArticleOfNumberStr: %#v", startFromArticleOfNumberStr)

	if startFromArticleOfNumberStr == "" {
		startFromArticleOfNumberStr = "0"
	}

	startFromArticleOfNumber, err := strconv.Atoi(startFromArticleOfNumberStr)
	if err != nil {
		fc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}
	if startFromArticleOfNumber < 0 {
		fc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(errors.New(fmt.Sprintf("startFromArticleOfNumber = %d < 0", startFromArticleOfNumber)))
		return c.NoContent(http.StatusBadRequest)
	}

	articles, err := fc.feedUsecase.GetFeed(c.Request().Context())
	if err != nil {
		fc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

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
			Id:           v.ArticleId,
			Title:        v.Title,
			Description:  v.Description,
			Tags:         v.Tags,
			Category:     v.CategoryName,
			Rating:       v.Rating,
			Comments:     v.CommentsCount,
			CoverImgPath: v.CoverImgPath,
			Publisher: modelsRestApi.Publisher{
				Username: v.Publisher.Username,
				Login:    v.Publisher.Login,
			},
			CoAuthor: modelsRestApi.CoAuthor{
				Username: v.CoAuthor.Username,
				Login:    v.CoAuthor.Login,
			},
			Liked: v.Liked,
		}
		feed.Articles = append(feed.Articles, article)
	}
	fc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Formed feed: ", feed)

	return c.JSON(http.StatusOK, feed)
}

func (fc *FeedController) FeedUserHandler(c echo.Context) error {
	fc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the FeedUserHandler function.")

	login := c.QueryParam("login")
	fc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed login: %#v", login)
	if login == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	startFromArticleOfNumberStr := c.QueryParam("startFromArticleOfNumber")
	fc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed startFromArticleOfNumberStr: %#v", startFromArticleOfNumberStr)

	if startFromArticleOfNumberStr == "" {
		startFromArticleOfNumberStr = "0"
	}

	startFromArticleOfNumber, err := strconv.Atoi(startFromArticleOfNumberStr)
	if err != nil {
		fc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}
	if startFromArticleOfNumber < 0 {
		fc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(errors.New(fmt.Sprintf("startFromArticleOfNumber = %d < 0", startFromArticleOfNumber)))
		return c.NoContent(http.StatusBadRequest)
	}

	articles, err := fc.feedUsecase.GetFeedForUserByLogin(c.Request().Context(), login)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.LoginIsNotValidError:
			fc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.NoContent(http.StatusBadRequest)
		case *usecaseToDeliveryErrors.LoginDoesntExistError:
			fc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.NoContent(http.StatusNotFound)
		default:
			fc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

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
			Id:           v.ArticleId,
			Title:        v.Title,
			Description:  v.Description,
			Tags:         v.Tags,
			Category:     v.CategoryName,
			Rating:       v.Rating,
			Comments:     v.CommentsCount,
			CoverImgPath: v.CoverImgPath,
			Publisher: modelsRestApi.Publisher{
				Username: v.Publisher.Username,
				Login:    v.Publisher.Login,
			},
			CoAuthor: modelsRestApi.CoAuthor{
				Username: v.CoAuthor.Username,
				Login:    v.CoAuthor.Login,
			},
		}
		feed.Articles = append(feed.Articles, article)
	}
	fc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Formed feed: ", feed)

	return c.JSON(http.StatusOK, feed)
}

func (fc *FeedController) FeedCategoryHandler(c echo.Context) error {
	fc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the FeedCategoryHandler function.")

	category := c.QueryParam("category")
	fc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed category: %#v", category)
	if category == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	startFromArticleOfNumberStr := c.QueryParam("startFromArticleOfNumber")
	fc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed startFromArticleOfNumberStr: %#v", startFromArticleOfNumberStr)

	if startFromArticleOfNumberStr == "" {
		startFromArticleOfNumberStr = "0"
	}

	startFromArticleOfNumber, err := strconv.Atoi(startFromArticleOfNumberStr)
	if err != nil {
		fc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}
	if startFromArticleOfNumber < 0 {
		fc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(errors.New(fmt.Sprintf("startFromArticleOfNumber = %d < 0", startFromArticleOfNumber)))
		return c.NoContent(http.StatusBadRequest)
	}

	articles, err := fc.feedUsecase.GetFeedForCategory(c.Request().Context(), category)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.CategoryDoesntExistError:
			fc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.NoContent(http.StatusNotFound)
		default:
			fc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

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
			Id:           v.ArticleId,
			Title:        v.Title,
			Description:  v.Description,
			Tags:         v.Tags,
			Category:     v.CategoryName,
			Rating:       v.Rating,
			Comments:     v.CommentsCount,
			CoverImgPath: v.CoverImgPath,
			Publisher: modelsRestApi.Publisher{
				Username: v.Publisher.Username,
				Login:    v.Publisher.Login,
			},
			CoAuthor: modelsRestApi.CoAuthor{
				Username: v.CoAuthor.Username,
				Login:    v.CoAuthor.Login,
			},
		}
		feed.Articles = append(feed.Articles, article)
	}
	fc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Formed feed: ", feed)

	return c.JSON(http.StatusOK, feed)
}
