package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/searchComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/searchComponent/delivery/modelsRestApi"
	"2022_2_GoTo_team/pkg/utils/logger"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SearchController struct {
	searchUsecase searchComponentInterfaces.SearchUsecaseInterface
	logger        *logger.Logger
}

func NewSearchController(searchUsecase searchComponentInterfaces.SearchUsecaseInterface, logger *logger.Logger) *SearchController {
	logger.LogrusLogger.Debug("Enter to the NewSearchController function.")

	searchController := &SearchController{
		searchUsecase: searchUsecase,
		logger:        logger,
	}

	logger.LogrusLogger.Info("SearchController has created.")

	return searchController
}

func (sc *SearchController) SearchHandler(c echo.Context) error {
	ctx := c.Request().Context()
	sc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the SearchHandler function.")

	substringToSearch := c.QueryParam("substringToSearch")
	sc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed substringToSearch: %#v", substringToSearch)
	authorLogin := c.QueryParam("author")
	sc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed authorLogin: %#v", authorLogin)
	categoryName := c.QueryParam("category")
	sc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed categoryName: %#v", categoryName)
	tagName := c.QueryParam("tag")
	sc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed tagName: %#v", tagName)

	tags, err := sc.searchUsecase.GetArticlesBySearchParameters(ctx, substringToSearch, authorLogin, categoryName, tagName)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		default:
			sc.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	feed := modelsRestApi.Feed{}
	for _, v := range tags {
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
	sc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Formed feed: ", feed)

	jsonBytes, err := feed.MarshalJSON()
	if err != nil {
		sc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}

/*
func (sc *SearchController) SearchTagHandler(c echo.Context) error {
	ctx := c.Request().Context()
	sc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the SearchTagHandler function.")

	tag := c.QueryParam("tag")
	sc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed tag: %#v", tag)
	if tag == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	tags, err := sc.searchUsecase.GetArticlesByTag(ctx, tag)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.TagDoesntExistError:
			sc.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusNotFound)
		default:
			sc.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	feed := modelsRestApi.Feed{}
	for _, v := range tags {
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
	sc.logger.LogrusLoggerWithContext(ctx).Debug("Formed feed: ", feed)

	jsonBytes, err := feed.MarshalJSON()
	if err != nil {
		sc.logger.LogrusLoggerWithContext(ctx).Error(err)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}
*/
