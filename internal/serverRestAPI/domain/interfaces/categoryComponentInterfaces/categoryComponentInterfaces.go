package categoryComponentInterfaces

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"context"
)

type CategoryUsecaseInterface interface {
	GetCategoryInfo(ctx context.Context, category string) (*models.Category, error)
	IsUserSubscribedOnCategory(ctx context.Context, categoryName string) (bool, error)
	GetCategoryList(ctx context.Context) ([]*models.Category, error)
	SubscribeOnCategory(ctx context.Context, categoryName string) error
	UnsubscribeFromCategory(ctx context.Context, categoryName string) error
}

type CategoryRepositoryInterface interface {
	GetCategoryInfo(ctx context.Context, category string) (*models.Category, error)
	IsUserSubscribedOnCategory(ctx context.Context, userEmail string, categoryName string) (bool, error)
	GetAllCategories(ctx context.Context) ([]*models.Category, error)
	SubscribeOnCategory(ctx context.Context, email string, categoryName string) error
	UnsubscribeFromCategory(ctx context.Context, email string, categoryName string) (int64, error)
}
