package repository

import (
	"eCommerce/internal/models"
	"gorm.io/gorm"
)

type productRepositoryImp struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return productRepositoryImp{db: db}
}

func (p productRepositoryImp) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p productRepositoryImp) GetOne(productID int) (*models.Product, error) {
	var product models.Product
	err := p.db.Preload("Section").Preload("Category").Preload("Brand").First(&product, "id = ?", productID).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p productRepositoryImp) Create(product models.Product) error {
	return p.db.Create(&product).Error
}

func (p productRepositoryImp) Update(productID int, product models.Product) error {
	return p.db.Model(&models.Product{}).Where("id = ?", productID).Updates(&product).Error
}

func (p productRepositoryImp) Delete(productID int) error {
	return p.db.Delete(&models.Product{}, "id = ?", productID).Error
}

// helper functions

func (p productRepositoryImp) CheckProductNameTk(sectionNameTk string) (bool, error) {
	err := p.db.Where("product_name_tk=?", sectionNameTk).First(&models.Product{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	return true, nil
}

func (p productRepositoryImp) CheckProductNameRu(sectionNameRu string) (bool, error) {
	err := p.db.Where("product_name_ru=?", sectionNameRu).First(&models.Product{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	return true, nil
}

func (p productRepositoryImp) CheckProductNameEn(sectionNameEn string) (bool, error) {
	err := p.db.Where("product_name_en=?", sectionNameEn).First(&models.Product{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	return true, nil
}
