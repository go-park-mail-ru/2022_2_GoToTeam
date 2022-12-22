package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/feedComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/feedComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/feedComponent/delivery/modelsRestApi"
	"2022_2_GoTo_team/pkg/utils/logger"
	"errors"
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
	ctx := c.Request().Context()
	fc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the FeedHandler function.")

	articles, err := fc.feedUsecase.GetFeed(ctx)
	if err != nil {
		fc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
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
	fc.logger.LogrusLoggerWithContext(ctx).Debug("Formed feed: ", feed)

	jsonBytes, err := feed.MarshalJSON()
	if err != nil {
		fc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (fc *FeedController) FeedUserHandler(c echo.Context) error {
	ctx := c.Request().Context()
	fc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the FeedUserHandler function.")

	login := c.QueryParam("login")
	fc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed login: %#v", login)
	if login == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	articles, err := fc.feedUsecase.GetFeedForUserByLogin(ctx, login)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.LoginIsNotValidError:
			fc.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusBadRequest)
		case *usecaseToDeliveryErrors.LoginDoesntExistError:
			fc.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusNotFound)
		default:
			fc.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
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
	fc.logger.LogrusLoggerWithContext(ctx).Debug("Formed feed: ", feed)

	jsonBytes, err := feed.MarshalJSON()
	if err != nil {
		fc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (fc *FeedController) FeedCategoryHandler(c echo.Context) error {
	ctx := c.Request().Context()
	fc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the FeedCategoryHandler function.")

	category := c.QueryParam("category")
	fc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed category: %#v", category)
	if category == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	articles, err := fc.feedUsecase.GetFeedForCategory(ctx, category)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.CategoryDoesntExistError:
			fc.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusNotFound)
		default:
			fc.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
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
	fc.logger.LogrusLoggerWithContext(ctx).Debug("Formed feed: ", feed)

	jsonBytes, err := feed.MarshalJSON()
	if err != nil {
		fc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (fc *FeedController) GetNewArticlesFromIdForSubscriber(c echo.Context) error {
	ctx := c.Request().Context()
	fc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetNewArticlesFromIdForSubscriber function.")
	defer c.Request().Body.Close()

	articleIdStr := c.QueryParam("articleId")
	fc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed articleIdStr: %#v", articleIdStr)
	if articleIdStr == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	articleId, err := strconv.Atoi(articleIdStr)
	if err != nil {
		fc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusBadRequest)
	}

	newArticlesIds, err := fc.feedUsecase.GetNewArticlesFromIdForSubscriber(ctx, articleId)
	if err != nil {
		fc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	newArticlesIdsResponse := modelsRestApi.NewArticlesIds{
		Ids: newArticlesIds,
	}

	jsonBytes, err := newArticlesIdsResponse.MarshalJSON()
	if err != nil {
		fc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}
