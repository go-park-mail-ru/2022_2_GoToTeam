package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/searchComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/searchComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/errorsUtils"
	"2022_2_GoTo_team/pkg/logger"
	"context"
	"fmt"
)

type searchUsecase struct {
	searchRepository searchComponentInterfaces.SearchRepositoryInterface
	logger           *logger.Logger
}

func NewSearchUsecase(searchRepository searchComponentInterfaces.SearchRepositoryInterface, logger *logger.Logger) searchComponentInterfaces.SearchUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewSearchUsecase function.")

	searchUsecase := &searchUsecase{
		searchRepository: searchRepository,
		logger:           logger,
	}

	logger.LogrusLogger.Info("searchUsecase has created.")

	return searchUsecase
}

func (su *searchUsecase) GetArticlesByTag(ctx context.Context, tag string) ([]*models.Article, error) {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetArticlesByTag function.")

	wrappingErrorMessage := "error while getting articles by tag"

	exists, err := su.searchRepository.TagExists(ctx, tag)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}
	if !exists {
		su.logger.LogrusLoggerWithContext(ctx).Infof("Tag %s doesnt exist", tag)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.TagDoesntExistError{Err: fmt.Errorf("tag %#v doesnt exist", su)})
	}

	articles, err := su.searchRepository.GetArticlesByTag(ctx, tag)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return articles, nil
}

func (su *searchUsecase) GetArticlesBySearchParameters(ctx context.Context, substringToSearch string, login string, categoryName string, tagName string) ([]*models.Article, error) {
	su.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetArticlesBySearchParameters function.")

	wrappingErrorMessage := "error while getting articles by search parameters"

	articles, err := su.searchRepository.GetArticlesBySearchParameters(ctx, substringToSearch, login, categoryName, tagName)
	if err != nil {
		su.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return articles, nil
}
