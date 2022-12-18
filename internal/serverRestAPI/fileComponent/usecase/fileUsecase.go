package usecase

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/fileComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/fileComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/utils/sessionUtils"
	"2022_2_GoTo_team/pkg/utils/errorsUtils"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

const (
	_MAX_BYTES_FILE_SIZE = 4194304
)

type fileUsecase struct {
	fileRepository fileComponentInterfaces.FileRepositoryInterface
	logger         *logger.Logger
}

func NewFileUsecase(fileRepository fileComponentInterfaces.FileRepositoryInterface, logger *logger.Logger) fileComponentInterfaces.FileUsecaseInterface {
	logger.LogrusLogger.Debug("Enter to the NewFileUsecase function.")

	fileUsecase := &fileUsecase{
		fileRepository: fileRepository,
		logger:         logger,
	}

	logger.LogrusLogger.Info("fileUsecase has created.")

	return fileUsecase
}

func (fu *fileUsecase) UploadProfilePhoto(ctx context.Context, fileHeader *multipart.FileHeader) error {
	fu.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UploadProfilePhoto function.")

	wrappingErrorMessage := "error while file uploading"

	email, err := sessionUtils.GetEmailFromContext(ctx, fu.logger)
	if err != nil {
		fu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.EmailForSessionDoesntExistError{Err: err})
	}

	fu.logger.LogrusLoggerWithContext(ctx).Debug("fileName: ", fileHeader.Filename)
	fileSize := fileHeader.Size
	fu.logger.LogrusLoggerWithContext(ctx).Debug("fileSize: ", fileSize)
	if fileSize > _MAX_BYTES_FILE_SIZE {
		fu.logger.LogrusLoggerWithContext(ctx).Error("File size is greater then ", _MAX_BYTES_FILE_SIZE, " bytes.")
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.FileSizeBigError{Err: fmt.Errorf("file size is greater then %d bytes", _MAX_BYTES_FILE_SIZE)})
	}

	file, err := fileHeader.Open()
	if err != nil {
		fu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.OpenFileError{Err: err})
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fu.logger.LogrusLoggerWithContext(ctx).Error(err)
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.OpenFileError{Err: err})
	}

	mimeFileType := http.DetectContentType(fileBytes)
	fu.logger.LogrusLoggerWithContext(ctx).Debug("mimeFileType: ", mimeFileType)

	if !mimeTypeIsSubtypeOf(mimeFileType, "image") {
		fu.logger.LogrusLoggerWithContext(ctx).Error("Mime type is not image.")
		return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.NotImageError{Err: errors.New("mime type is not image")})
	}

	fileType := extractTypeFromMimeType(mimeFileType)
	fu.logger.LogrusLoggerWithContext(ctx).Debug("fileType: ", fileType)

	err = fu.fileRepository.UploadProfilePhoto(ctx, fileBytes, fileType, email)
	if err != nil {
		switch err {
		default:
			fu.logger.LogrusLoggerWithContext(ctx).Error(err)
			return errorsUtils.WrapError(wrappingErrorMessage, &usecaseToDeliveryErrors.RepositoryError{
				Err: err,
			})
		}
	}

	return nil
}

func mimeTypeIsSubtypeOf(mimeType string, pattern string) bool {
	if strings.HasPrefix(mimeType, pattern) {
		return true
	}

	return false
}

func extractTypeFromMimeType(mimeType string) string {
	splittedStrings := strings.Split(mimeType, "/")

	return splittedStrings[len(splittedStrings)-1]
}
