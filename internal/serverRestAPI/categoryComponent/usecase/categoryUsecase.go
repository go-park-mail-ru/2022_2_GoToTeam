package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/categoryComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/categoryComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/categoryComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/errorsUtils"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/logger"
	"context"
)

type categoryUsecase struct {
	categoryRepository categoryComponentInterfaces.CategoryRepositoryInterface
	logger             *logger.Logger
}

func NewCategoryUsecase(categoryRepository categoryComponentInterfaces.CategoryRepositoryInterface, logger *logger.Logger) categoryComponentInterfaces.CategoryUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewCategoryUsecase function.")

	categoryUsecase := &categoryUsecase{
		categoryRepository: categoryRepository,
		logger:             logger,
	}

	logger.LogrusLogger.Info("categoryUsecase has created.")

	return categoryUsecase
}

func (cu *categoryUsecase) GetCategoryInfo(ctx context.Context, categoryName string) (*models.Category, error) {
	cu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetCategoryInfo function.")

	wrappingErrorMessage := "error while getting info for category"

	category, err := cu.categoryRepository.GetCategoryInfo(ctx, categoryName)
	if err != nil {
		if err == repositoryToUsecaseErrors.CategoryRepositoryCategoryDontExistsError {
			cu.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.CategoryDontExistsError{Err: err})
		}
		cu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return category, nil
}
