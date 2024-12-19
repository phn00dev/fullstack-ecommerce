package constructor

import (
	"eCommerce/internal/domain/banner/handler"
	"eCommerce/internal/domain/banner/repository"
	"eCommerce/internal/domain/banner/service"
	"eCommerce/pkg/config"
	"gorm.io/gorm"
)

var (
	bannerRepo    repository.BannerRepository
	bannerService service.BannerService
	BannerHandler handler.BannerHandler
)

func InitBannerRequirementCreator(db *gorm.DB, config config.Config) {
	bannerRepo = repository.NewBannerRepository(db)
	bannerService = service.NewBannerService(bannerRepo)
	BannerHandler = handler.NewBannerHandler(bannerService, config)
}
