package repository

import (
	"2022_2_GoTo_team/internal/serverRestAPI/domain/customErrors/categoryComponentErrors/repositoryToUsecaseErrors"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/interfaces/categoryComponentInterfaces"
	"2022_2_GoTo_team/internal/serverRestAPI/domain/models"
	"2022_2_GoTo_team/pkg/utils/logger"
	"context"
	"database/sql"
	"fmt"
)

type categoryPostgreSQLRepository struct {
	database *sql.DB
	logger   *logger.Logger
}

func NewCategoryPostgreSQLRepository(database *sql.DB, logger *logger.Logger) categoryComponentInterfaces.CategoryRepositoryInterface {
	logger.LogrusLogger.Debug("Enter to the NewCategoryPostgreSQLRepository function.")

	categoryRepository := &categoryPostgreSQLRepository{
		database: database,
		logger:   logger,
	}

	logger.LogrusLogger.Debug("All categories in storage:  \n" + func() string {
		allCategories, err := categoryRepository.GetAllCategories(context.Background())
		if err != nil {
			return repositoryToUsecaseErrors.CategoryRepositoryError.Error()
		}
		return categoryRepository.getCategoriesString(allCategories)
	}())

	logger.LogrusLogger.Info("categoryPostgreSQLRepository has created.")

	return categoryRepository
}

func (cpsr *categoryPostgreSQLRepository) getCategoriesString(categories []*models.Category) string {
	categoriesString := ""
	for _, v := range categories {
		categoriesString += fmt.Sprintf("%#v\n", v)
	}

	return categoriesString
}

func (cpsr *categoryPostgreSQLRepository) GetCategoryInfo(ctx context.Context, categoryName string) (*models.Category, error) {
	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetCategoryInfo function.")

	row := cpsr.database.QueryRow(`
SELECT category_name, description, subscribers_count
FROM categories
WHERE category_name = $1;
`, categoryName)

	category := &models.Category{}
	if err := row.Scan(&category.CategoryName, &category.Description, &category.SubscribersCount); err != nil {
		if err == sql.ErrNoRows {
			cpsr.logger.LogrusLoggerWithContext(ctx).Debug(err)
			return nil, repositoryToUsecaseErrors.CategoryRepositoryCategoryDoesntExistError
		}
		cpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.CategoryRepositoryError
	}

	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Got category: %#v", category)

	return category, nil
}

func (cpsr *categoryPostgreSQLRepository) IsUserSubscribedOnCategory(ctx context.Context, userEmail string, categoryName string) (bool, error) {
	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the IsUserSubscribedOnCategory function.")

	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Input userEmail = ", userEmail, " categoryName = ", categoryName)

	row := cpsr.database.QueryRow(`
SELECT COUNT(*) count
FROM categories C
JOIN users_categories_subscriptions UCS ON C.category_id = UCS.category_id
JOIN users U ON U.user_id = UCS.user_id 
WHERE U.email = $1 AND C.category_name = $2;
`, userEmail, categoryName)

	entriesFound := 0
	if err := row.Scan(&entriesFound); err != nil {
		cpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return false, repositoryToUsecaseErrors.CategoryRepositoryError
	}

	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Got entriesFound: ", entriesFound)

	result := false
	if entriesFound == 1 {
		result = true
	}

	return result, nil
}

func (cpsr *categoryPostgreSQLRepository) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	cpsr.logger.LogrusLoggerWithContext(ctx).Debug("Enter to the GetAllCategories function.")

	categories := make([]*models.Category, 0, 10)

	rows, err := cpsr.database.Query(`
SELECT category_name
FROM categories;
`)
	if err != nil {
		cpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
		return nil, repositoryToUsecaseErrors.CategoryRepositoryError
	}
	defer rows.Close()

	for rows.Next() {
		category := &models.Category{}
		if err := rows.Scan(&category.CategoryName); err != nil {
			cpsr.logger.LogrusLoggerWithContext(ctx).Error(err)
			return nil, repositoryToUsecaseErrors.CategoryRepositoryError
		}

		categories = append(categories, category)
	}

	cpsr.logger.LogrusLoggerWithContext(ctx).Debugf("Got categories: %#v\n", categories)

	return categories, nil
}
