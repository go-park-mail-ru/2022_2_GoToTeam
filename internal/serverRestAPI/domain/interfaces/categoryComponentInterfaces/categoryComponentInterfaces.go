package categoryComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type CategoryUsecaseInterface interface {
	GetCategoryInfo(ctx context.Context, category string) (*models.Category, error)
	IsUserSubscribedOnCategory(ctx context.Context, categoryName string) (bool, error)
	GetCategoryList(ctx context.Context) ([]*models.Category, error)
}

type CategoryRepositoryInterface interface {
	GetCategoryInfo(ctx context.Context, category string) (*models.Category, error)
	IsUserSubscribedOnCategory(ctx context.Context, userEmail string, categoryName string) (bool, error)
	GetAllCategories(ctx context.Context) ([]*models.Category, error)
}
