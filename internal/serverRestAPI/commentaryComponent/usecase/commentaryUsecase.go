package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/commentaryComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/commentaryComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/sessionUtils"
	"2022_2_GoTo_team/pkg/domain"
	"2022_2_GoTo_team/pkg/utils/errorsUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
)

type commentaryUsecase struct {
	commentaryRepository commentaryComponentInterfaces.CommentaryRepositoryInterface
	logger               *logger.Logger
}

func NewCommentaryUsecase(commentaryRepository commentaryComponentInterfaces.CommentaryRepositoryInterface, logger *logger.Logger) commentaryComponentInterfaces.CommentaryUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewCommentaryUsecase function.")

	commentaryUsecase := &commentaryUsecase{
		commentaryRepository: commentaryRepository,
		logger:               logger,
	}

	logger.LogrusLogger.Info("commentaryUsecase has created.")

	return commentaryUsecase
}

func (acbs *commentaryUsecase) AddCommentaryBySession(ctx context.Context, commentary *models.Commentary) error {
	acbs.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the AddCommentaryBySession function.")

	wrappingErrorMessage := "error while adding new commentary by session"

	email := ctx.Value(domain.USER_EMAIL_KEY_FOR_CONTEXT)
	acbs.logger.LogrusLoggerWithContext(ctx).Debug("Email from context = ", email)

	if email == nil || email.(string) == "" {
		acbs.logger.LogrusLoggerWithContext(ctx).Error("Email from context is empty.")
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{Err: errors.New("email from context is empty")})
	}
	commentary.Publisher.Email = email.(string)

	_, err := acbs.commentaryRepository.AddCommentaryByEmail(ctx, commentary)
	if err != nil {
		acbs.logger.LogrusLoggerWithContext(ctx).Error(err)
		return errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	return nil
}

func (acbs *commentaryUsecase) GetAllCommentariesForArticle(ctx context.Context, articleId int) ([]*models.Commentary, error) {
	acbs.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetAllCommentariesForArticle function.")

	wrappingErrorMessage := "error while getting commentaries for articleId"

	articles, err := acbs.commentaryRepository.GetAllCommentsForArticle(ctx, articleId)
	if err != nil {
		acbs.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return articles, nil
}

func (acbs *commentaryUsecase) ProcessLike(ctx context.Context, likeData *models.LikeData) (int, error) {
	acbs.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the ProcessLike function.")

	wrappingErrorMessage := "error while liking the commentary by session"

	email, err := sessionUtils.GetEmailFromContext(ctx, acbs.logger)
	if err != nil {
		acbs.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{Err: err})
	}

	sign := likeData.Sign
	if sign == 1 { // Like
		_, err = acbs.commentaryRepository.AddLike(ctx, true, likeData.Id, email)
	} else if sign == -1 { // Dislike
		_, err = acbs.commentaryRepository.AddLike(ctx, false, likeData.Id, email)
	} else if sign == 0 { // Remove like if exist
		_, err = acbs.commentaryRepository.RemoveLike(ctx, likeData.Id, email)
	}
	if err != nil {
		acbs.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	updatedRating, err := acbs.commentaryRepository.GetCommentaryRating(ctx, likeData.Id)
	if err != nil {
		acbs.logger.LogrusLoggerWithContext(ctx).Error(err)
		return 0, errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{Err: err})
	}

	return updatedRating, err
}
