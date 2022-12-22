package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/fileComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/fileComponentInterfaces"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"database/sql"
	"os"
	"path"
	"strconv"
	"time"
)

type fileRepository struct {
	database              *sql.DB
	staticDirAbsolutePath string
	profilePhotosDirPath  string
	logger                *logger.Logger
}

func NewFileRepository(database *sql.DB, staticDirAbsolutePath string, profilePhotosDirPath string, logger *logger.Logger) fileComponentInterfaces.FileRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewTagPostgreSQLRepository function.")

	fileRepository := &fileRepository{
		database:              database,
		staticDirAbsolutePath: staticDirAbsolutePath,
		profilePhotosDirPath:  profilePhotosDirPath,
		logger:                logger,
	}

	logger.LogrusLogger.Info("fileRepository has created.")

	return fileRepository
}

func (fr *fileRepository) UploadProfilePhoto(ctx context.Context, fileBytes []byte, fileType string, email string) error {
	fr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the UploadProfilePhoto function.")

	//fr.logger.LogrusLoggerWithContext(ctx).Debugf("Input fileBytes: %#v", fileBytes)
	/*
		fileBytes, err := io.ReadAll(*file)
		if err != nil {
			fr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return repositoryToUsecaseErrors.FileRepositoryError
		}
		fr.logger.LogrusLoggerWithContext(ctx).Debug("fileBytes: ", fileBytes)
	*/

	newFileName := strconv.FormatInt(time.Now().Unix(), 10) + "." + fileType

	newFilePathInFS := fr.staticDirAbsolutePath + fr.profilePhotosDirPath + "/" + newFileName
	fr.logger.LogrusLoggerWithContext(ctx).Debug("newFilePathInFS: ", newFilePathInFS)

	newFilePathInPostgreSQL := "/" + fr.profilePhotosDirPath + "/" + newFileName
	fr.logger.LogrusLoggerWithContext(ctx).Debug("newFilePathInPostgreSQL: ", newFilePathInPostgreSQL)

	if err := os.MkdirAll(path.Dir(newFilePathInFS), 0750); err != nil && !os.IsExist(err) {
		fr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return repositoryToUsecaseErrors.FileRepositoryError
	}

	// Destination
	/*
		dst, err := os.Create(newFilePathInFS)
		if err != nil {
			fr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return repositoryToUsecaseErrors.FileRepositoryError
		}
		defer dst.Close()
	*/

	err := os.WriteFile(newFilePathInFS, fileBytes, 0777)
	if err != nil {
		fr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return repositoryToUsecaseErrors.FileRepositoryError
	}
	//fr.logger.LogrusLoggerWithContext(ctx).Debug("Copy file: ", file)
	//fr.logger.LogrusLoggerWithContext(ctx).Debug("Written bytes: ", written)

	_, err = fr.database.Exec(`
UPDATE users SET avatar_img_path = $1
WHERE email = $2;
`, newFilePathInPostgreSQL, email)

	if err != nil {
		fr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return repositoryToUsecaseErrors.FileRepositoryError
	}

	fr.logger.LogrusLoggerWithContext(ctx).Debug("File saved successfully\n")

	return nil
}
