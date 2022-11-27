package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/tagComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/tagComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/errorsUtils"
	"2022_2_GoTo_team/pkg/logger"
	"context"
)

type tagUsecase struct {
	tagRepository tagComponentInterfaces.TagRepositoryInterface
	logger        *logger.Logger
}

func NewTagUsecase(tagRepository tagComponentInterfaces.TagRepositoryInterface, logger *logger.Logger) tagComponentInterfaces.TagUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewTagUsecase function.")

	tagUsecase := &tagUsecase{
		tagRepository: tagRepository,
		logger:        logger,
	}

	logger.LogrusLogger.Info("tagUsecase has created.")

	return tagUsecase
}

func (tu *tagUsecase) GetTagsList(ctx context.Context) ([]*models.Tag, error) {
	tu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetTagsList function.")

	wrappingErrorMessage := "error while getting tags list"

	tags, err := tu.tagRepository.GetAllTags(ctx)
	if err != nil {
		switch err {
		default:
			tu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
		}
	}

	return tags, nil
}
