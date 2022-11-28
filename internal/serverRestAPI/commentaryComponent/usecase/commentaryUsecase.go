package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/commentaryComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/commentaryComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/errorsUtils"
	"2022_2_GoTo_team/pkg/logger"
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

	//authorEmail, err := au.sessionRepository.GetEmailBySession(ctx, session)
	authorEmail := ctx.Value(domain.USER_EMAIL_KEY_FOR_CONTEXT).(string)
	acbs.logger.LogrusLoggerWithContext(ctx).Debug("Email from context = ", authorEmail)

	if authorEmail == "" {
		acbs.logger.LogrusLoggerWithContext(ctx).Error("Email from context is empty.")
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{Err: errors.New("email from context is empty")})
	}
	commentary.Publisher.Email = authorEmail

	_, err := acbs.commentaryRepository.AddCommentaryByEmail(ctx, commentary)
	if err != nil {
		acbs.logger.LogrusLoggerWithContext(ctx).Error(err)
		return errorsUtils.WrapError(wrappingErrorMessage, err)
	}

	return nil
}
