package tagComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

//go:generate mockgen -destination=./mock/tagRepositoryMock.go -package=mock 2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/tagComponentInterfaces TagRepositoryInterface

type TagUsecaseInterface interface {
	GetTagsList(ctx context.Context) ([]*models.Tag, error)
}

type TagRepositoryInterface interface {
	GetAllTags(ctx context.Context) ([]*models.Tag, error)
}
