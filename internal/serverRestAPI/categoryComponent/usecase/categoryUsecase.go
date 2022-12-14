package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/categoryComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/categoryComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/categoryComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/domain"
	"2022_2_GoTo_team/pkg/utils/errorsUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
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
		switch err {
		case repositoryToUsecaseErrors.CategoryRepositoryCategoryDoesntExistError:
			cu.logger.LogrusLoggerWithContext(ctx).Warn(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.CategoryNotFoundError{Err: err})
		default:
			cu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	return category, nil
}

func (cu *categoryUsecase) IsUserSubscribedOnCategory(ctx context.Context, categoryName string) (bool, error) {
	cu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the IsUserSubscribedOnCategory function.")

	wrappingErrorMessage := "error while IsUserSubscribedOnCategory for category"

	email := ctx.Value(domain.USER_EMAIL_KEY_FOR_CONTEXT)
	cu.logger.LogrusLoggerWithContext(ctx).Debug("Email from context = ", email)

	if email == nil || email.(string) == "" {
		cu.logger.LogrusLoggerWithContext(ctx).Error("Email from context is empty.")
		return false, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{Err: errors.New("email from context is empty")})
	}

	isSubscribed, err := cu.categoryRepository.IsUserSubscribedOnCategory(ctx, email.(string), categoryName)
	if err != nil {
		switch err {
		default:
			cu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return false, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	return isSubscribed, nil
}

func (cu *categoryUsecase) GetCategoryList(ctx context.Context) ([]*models.Category, error) {
	cu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetCategoryList function.")

	wrappingErrorMessage := "error while getting category list"

	categories, err := cu.categoryRepository.GetAllCategories(ctx)
	if err != nil {
		switch err {
		default:
			cu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	return categories, nil
}

func (cu *categoryUsecase) SubscribeOnCategory(ctx context.Context, categoryName string) error {
	cu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the SubscribeOnCategory function.")

	wrappingErrorMessage := "error while subscribing on category"

	email := ctx.Value(domain.USER_EMAIL_KEY_FOR_CONTEXT)
	cu.logger.LogrusLoggerWithContext(ctx).Debug("Email from context = ", email)

	if email == nil || email.(string) == "" {
		cu.logger.LogrusLoggerWithContext(ctx).Error("Email from context is empty.")
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{Err: errors.New("email from context is empty")})
	}

	if err := cu.categoryRepository.SubscribeOnCategory(ctx, email.(string), categoryName); err != nil {
		cu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return nil
}

func (cu *categoryUsecase) UnsubscribeFromCategory(ctx context.Context, categoryName string) error {
	cu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UnsubscribeFromCategory function.")

	wrappingErrorMessage := "error while unsubscribing from the category"

	email := ctx.Value(domain.USER_EMAIL_KEY_FOR_CONTEXT)
	cu.logger.LogrusLoggerWithContext(ctx).Debug("Email from context = ", email)

	if email == nil || email.(string) == "" {
		cu.logger.LogrusLoggerWithContext(ctx).Error("Email from context is empty.")
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{Err: errors.New("email from context is empty")})
	}

	if _, err := cu.categoryRepository.UnsubscribeFromCategory(ctx, email.(string), categoryName); err != nil {
		cu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return nil
}
