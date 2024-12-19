package repository

import (
	"eCommerce/internal/models"
	"gorm.io/gorm"
)

type bannerRepositoryImp struct {
	db *gorm.DB
}

func NewBannerRepository(db *gorm.DB) BannerRepository {
	return bannerRepositoryImp{db: db}
}

func (b bannerRepositoryImp) GetAll() ([]models.Banner, error) {
	var banners []models.Banner
	if err := b.db.Find(&banners).Error; err != nil {
		return nil, err
	}
	return banners, nil
}

func (b bannerRepositoryImp) GetOneByID(bannerId int) (*models.Banner, error) {
	var banner models.Banner
	if err := b.db.First(&banner, bannerId).Error; err != nil {
		return nil, err
	}
	return &banner, nil
}

func (b bannerRepositoryImp) Create(banner *models.Banner) error {
	return b.db.Create(banner).Error
}

func (b bannerRepositoryImp) Update(bannerID int, banner *models.Banner) error {
	return b.db.Model(&models.Banner{}).Where("id = ?", bannerID).Updates(banner).Error
}

func (b bannerRepositoryImp) Delete(bannerID int) error {
	return b.db.Delete(&models.Banner{}, bannerID).Error
}
