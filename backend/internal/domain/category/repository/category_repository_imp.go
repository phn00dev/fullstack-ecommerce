package repository

import (
	"eCommerce/internal/models"
	"gorm.io/gorm"
)

type categoryRepositoryImp struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return categoryRepositoryImp{db: db}
}

func (c categoryRepositoryImp) GetAll() ([]models.Category, error) {
	var categories []models.Category
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (c categoryRepositoryImp) GetOne(categoryID int) (*models.Category, error) {
	var category models.Category
	if err := c.db.First(&category, categoryID).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c categoryRepositoryImp) Create(category models.Category) error {
	return c.db.Create(&category).Error
}

func (c categoryRepositoryImp) Update(categoryID int, category models.Category) error {
	return c.db.Model(&models.Category{}).Where("id = ?", categoryID).Updates(&category).Error
}

func (c categoryRepositoryImp) Delete(categoryID int) error {
	return c.db.Delete(&models.Category{}, categoryID).Error
}

// helpers functions

func (c categoryRepositoryImp) CheckCategoryNameTk(categoryNameTk string) (bool, error) {
	err := c.db.Where("category_name_tk=?", categoryNameTk).First(&models.Section{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	return true, nil
}

func (c categoryRepositoryImp) CheckCategoryNameRu(categoryNameRu string) (bool, error) {
	err := c.db.Where("category_name_ru=?", categoryNameRu).First(&models.Category{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	return true, nil
}

func (c categoryRepositoryImp) CheckCategoryNameEn(categoryNameEn string) (bool, error) {
	err := c.db.Where("category_name_en=?", categoryNameEn).First(&models.Category{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	return true, nil
}
