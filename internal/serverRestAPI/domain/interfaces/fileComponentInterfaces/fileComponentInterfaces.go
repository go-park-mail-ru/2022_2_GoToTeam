package fileComponentInterfaces

import (
	"context"
	"mime/multipart"
)

type FileUsecaseInterface interface {
	UploadProfilePhoto(ctx context.Context, fileHeader *multipart.FileHeader) error
}

type FileRepositoryInterface interface {
	UploadProfilePhoto(ctx context.Context, fileBytes []byte, fileType string, email string) error
}
