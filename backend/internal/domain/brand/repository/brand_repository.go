package repository

import "eCommerce/internal/models"

type BrandRepository interface {
	GetAll() ([]models.Brand, error)
	GetOneByID(brandId int) (*models.Brand, error)
	Create(models.Brand) error
	Update(brandId int, brand *models.Brand) error
	Delete(brandId int) error

	// helper functions

	GetOneBrandByBrandName(brandName string) (models.Brand, error)
	CheckBrandNameTk(brandNameTk string) (bool, error)
	CheckBrandNameRu(brandNameRu string) (bool, error)
	CheckBrandNameEn(brandNameEn string) (bool, error)
}
