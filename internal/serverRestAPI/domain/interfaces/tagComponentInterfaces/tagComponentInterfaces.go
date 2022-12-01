package tagComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type TagUsecaseInterface interface {
	GetTagsList(ctx context.Context) ([]*models.Tag, error)
}

type TagRepositoryInterface interface {
	GetAllTags(ctx context.Context) ([]*models.Tag, error)
}
