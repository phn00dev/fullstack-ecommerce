package repository

import "eCommerce/internal/models"

type BannerRepository interface {
	GetAll() ([]models.Banner, error)
	GetOneByID(bannerId int) (*models.Banner, error)
	Create(banner *models.Banner) error
	Update(bannerID int, banner *models.Banner) error
	Delete(bannerID int) error
}
