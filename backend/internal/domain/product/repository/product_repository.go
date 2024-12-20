package repository

import "eCommerce/internal/models"

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetOne(productID int) (*models.Product, error)
	Create(product models.Product) error
	Update(productID int, product models.Product) error
	Delete(productID int) error

	// helper functions

	CheckProductNameTk(sectionNameTk string) (bool, error)
	CheckProductNameRu(sectionNameRu string) (bool, error)
	CheckProductNameEn(sectionNameEn string) (bool, error)
}
