package repository

import (
	"eCommerce/internal/models"
	"gorm.io/gorm"
)

type brandRepositoryImp struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return brandRepositoryImp{db: db}
}

func (b brandRepositoryImp) GetAll() ([]models.Brand, error) {
	var brands []models.Brand
	if err := b.db.Find(&brands).Error; err != nil {
		return nil, err
	}
	return brands, nil
}

func (b brandRepositoryImp) GetOneByID(brandId int) (*models.Brand, error) {
	var brand models.Brand
	if err := b.db.First(&brand, brandId).Error; err != nil {
		return nil, err
	}
	return &brand, nil
}

func (b brandRepositoryImp) Create(brand models.Brand) error {
	return b.db.Create(&brand).Error
}

func (b brandRepositoryImp) Update(brandId int, brand *models.Brand) error {
	return b.db.Model(&models.Brand{}).Where("id=?", brandId).Updates(&brand).Error
}

func (b brandRepositoryImp) Delete(brandId int) error {
	return b.db.Delete(&models.Brand{}, brandId).Error
}

func (b brandRepositoryImp) GetOneBrandByBrandName(brandName string) (models.Brand, error) {
	var brand models.Brand
	if err := b.db.Where("brand_name=?", brandName).First(&brand).Error; err != nil {
		return brand, err
	}
	return brand, nil
}
