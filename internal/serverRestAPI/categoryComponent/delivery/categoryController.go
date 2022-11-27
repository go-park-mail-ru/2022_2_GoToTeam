package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/categoryComponent/delivery/modelsRestApi"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/categoryComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/categoryComponentInterfaces"
	"2022_2_GoTo_team/pkg/logger"
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
	cc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the CategoryInfoHandler function.")

	categoryName := c.QueryParam("category")
	cc.logger.LogrusLoggerWithContext(c.Request().Context()).Debugf("Parsed category: %#v", categoryName)
	if categoryName == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	category, err := cc.categoryUsecase.GetCategoryInfo(c.Request().Context(), categoryName)
	if err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.CategoryNotFoundError:
			cc.logger.LogrusLoggerWithContext(c.Request().Context()).Warn(err)
			return c.NoContent(http.StatusNotFound)
		default:
			cc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	categoryInfo := modelsRestApi.CategoryInfo{
		CategoryName:     category.CategoryName,
		Description:      category.Description,
		SubscribersCount: category.SubscribersCount,
	}

	cc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Formed categoryInfo: ", categoryInfo)

	return c.JSON(http.StatusOK, categoryInfo)
}

func (cc *CategoryController) CategoryListHandler(c echo.Context) error {
	cc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Enter to the CategoryListHandler function.")

	categories, err := cc.categoryUsecase.GetCategoryList(c.Request().Context())
	if err != nil {
		cc.logger.LogrusLoggerWithContext(c.Request().Context()).Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	categoryList := modelsRestApi.CategoryList{}
	for _, v := range categories {
		categoryList.CategoryNames = append(categoryList.CategoryNames, v.CategoryName)
	}

	cc.logger.LogrusLoggerWithContext(c.Request().Context()).Debug("Formed categoryList: ", categoryList)

	return c.JSON(http.StatusOK, categoryList)
}
