package constructor

import (
	"eCommerce/internal/domain/brand/handler"
	"eCommerce/internal/domain/brand/repository"
	"eCommerce/internal/domain/brand/service"
	"eCommerce/pkg/config"
	"gorm.io/gorm"
)

var (
	brandRepo    repository.BrandRepository
	brandService service.BrandService
	BrandHandler handler.BrandHandler
)

func InitBrandRequirementCreator(db *gorm.DB, config config.Config) {
	brandRepo = repository.NewBrandRepository(db)
	brandService = service.NewBrandService(brandRepo)
	BrandHandler = handler.NewBrandHandler(brandService, config)
}
