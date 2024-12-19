package repository

import "eCommerce/internal/models"

type CategoryRepository interface {
	GetAll() ([]models.Category, error)
	GetOne(categoryID int) (*models.Category, error)
	Create(category models.Category) error
	Update(categoryID int, category models.Category) error
	Delete(categoryID int) error

	// helpers functions
	CheckCategoryNameTk(categoryNameTk string) (bool, error)
	CheckCategoryNameRu(categoryNameRu string) (bool, error)
	CheckCategoryNameEn(categoryNameEn string) (bool, error)
}
