package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/categoryComponent/delivery/modelsRestApi"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/categoryComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/categoryComponentInterfaces"
	"2022_2_GoTo_team/pkg/utils/logger"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CategoryController struct {
	categoryUsecase categoryComponentInterfaces.CategoryUsecaseInterface
	logger          *logger.Logger
}

func NewCategoryController(categoryUsecase categoryComponentInterfaces.CategoryUsecaseInterface, logger *logger.Logger) *CategoryController {
	logger.LogrusLogger.Debug("Enter to the NewCategoryController function.")

	categoryController := &CategoryController{
		categoryUsecase: categoryUsecase,
		logger:          logger,
	}

	logger.LogrusLogger.Info("CategoryController has created.")

	return categoryController
}

func (cc *CategoryController) CategoryInfoHandler(c echo.Context) error {
	ctx := c.Request().Context()
	cc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the CategoryInfoHandler function.")

	categoryName := c.QueryParam("category")
	cc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed category: %#v", categoryName)
	if categoryName == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	category, err := cc.categoryUsecase.GetCategoryInfo(ctx, categoryName)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.CategoryNotFoundError:
			cc.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusNotFound)
		default:
			cc.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	isSubscribed, err := cc.categoryUsecase.IsUserSubscribedOnCategory(ctx, categoryName)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailForSessionDoesntExistError:
			cc.logger.LogrusLoggerWithContext(ctx).Warn(err)
		default:
			cc.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	categoryInfo := modelsRestApi.CategoryInfo{
		CategoryName:     category.CategoryName,
		Description:      category.Description,
		SubscribersCount: category.SubscribersCount,
		Subscribed:       isSubscribed,
	}

	cc.logger.LogrusLoggerWithContext(ctx).Debug("Formed categoryInfo: ", categoryInfo)

	jsonBytes, err := categoryInfo.MarshalJSON()
	if err != nil {
		cc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (cc *CategoryController) CategoryListHandler(c echo.Context) error {
	ctx := c.Request().Context()
	cc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the CategoryListHandler function.")

	categories, err := cc.categoryUsecase.GetCategoryList(ctx)
	if err != nil {
		cc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	categoryList := modelsRestApi.CategoryList{}
	for _, v := range categories {
		categoryList.CategoryNames = append(categoryList.CategoryNames, v.CategoryName)
	}

	cc.logger.LogrusLoggerWithContext(ctx).Debug("Formed categoryList: ", categoryList)

	jsonBytes, err := categoryList.MarshalJSON()
	if err != nil {
		cc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSONBlob(http.StatusOK, jsonBytes)
}

func (cc *CategoryController) SubscribeHandler(c echo.Context) error {
	ctx := c.Request().Context()
	cc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the SubscribeHandler function.")
	defer c.Request().Body.Close()

	parsedInput := new(modelsRestApi.Subscribe)
	if err := c.Bind(parsedInput); err != nil {
		cc.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	cc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed input json data: %#v", parsedInput)

	if err := cc.categoryUsecase.SubscribeOnCategory(ctx, parsedInput.CategoryName); err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailForSessionDoesntExistError:
			cc.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusInternalServerError)
		default:
			cc.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	cc.logger.LogrusLoggerWithContext(ctx).Info("User subscribed successfully!")

	return c.NoContent(http.StatusOK)
}

func (cc *CategoryController) UnsubscribeHandler(c echo.Context) error {
	ctx := c.Request().Context()
	cc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UnsubscribeHandler function.")
	defer c.Request().Body.Close()

	parsedInput := new(modelsRestApi.Subscribe)
	if err := c.Bind(parsedInput); err != nil {
		cc.logger.LogrusLoggerWithContext(ctx).Warn(err)
		return c.NoContent(http.StatusBadRequest)
	}

	cc.logger.LogrusLoggerWithContext(ctx).Debugf("Parsed input json data: %#v", parsedInput)

	if err := cc.categoryUsecase.UnsubscribeFromCategory(ctx, parsedInput.CategoryName); err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailForSessionDoesntExistError:
			cc.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return c.NoContent(http.StatusInternalServerError)
		default:
			cc.logger.LogrusLoggerWithContext(ctx).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	cc.logger.LogrusLoggerWithContext(ctx).Info("User unsubscribed successfully!")

	return c.NoContent(http.StatusOK)
}
