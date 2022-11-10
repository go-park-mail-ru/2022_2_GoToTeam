package categoryComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type CategoryUsecaseInterface interface {
	GetCategoryInfo(ctx context.Context, category string) (*models.Category, error)
}

type CategoryRepositoryInterface interface {
	GetCategoryInfo(ctx context.Context, category string) (*models.Category, error)
}
