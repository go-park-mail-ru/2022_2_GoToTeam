package delivery

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/fileComponentErrors/usecaseToDeliveryErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/fileComponentInterfaces"
	"2022_2_GoTo_team/pkg/utils/logger"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type FileController struct {
	fileUsecase fileComponentInterfaces.FileUsecaseInterface
	logger      *logger.Logger
}

func NewFileController(fileUsecase fileComponentInterfaces.FileUsecaseInterface, logger *logger.Logger) *FileController {
	logger.LogrusLogger.Debug("Enter to the NewFileController function.")

	fileController := &FileController{
		fileUsecase: fileUsecase,
		logger:      logger,
	}

	logger.LogrusLogger.Info("FileController has created.")

	return fileController
}

func (fc *FileController) UploadProfilePhotoHandler(c echo.Context) error {
	ctx := c.Request().Context()
	fc.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UploadProfilePhotoHandler function.")
	defer c.Request().Body.Close()

	fileHeader, err := c.FormFile("file")
	if err != nil {
		fc.logger.LogrusLoggerWithContext(ctx).Error(err)
		return c.JSON(http.StatusBadRequest, "incorrect request")
	}
	fc.logger.LogrusLoggerWithContext(ctx).Debug("Input filename: ", fileHeader.Filename, " fileSize: ", fileHeader.Size, " header: ", fileHeader.Header)

	if err := fc.fileUsecase.UploadProfilePhoto(ctx, fileHeader); err != nil {
		switch errors.Unwrap(err).(type) {
		case *usecaseToDeliveryErrors.EmailForSessionDoesntExistError:
			return c.NoContent(http.StatusUnauthorized)
		case *usecaseToDeliveryErrors.FileSizeBigError:
			return c.JSON(http.StatusBadRequest, "file size big")
		case *usecaseToDeliveryErrors.OpenFileError:
			return c.JSON(http.StatusBadRequest, "broken file")
		case *usecaseToDeliveryErrors.NotImageError:
			return c.JSON(http.StatusBadRequest, "incorrect file type")
		default:
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	fc.logger.LogrusLoggerWithContext(ctx).Info("File uploaded successfully!")

	return c.NoContent(http.StatusOK)
}
