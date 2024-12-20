package constructor

import (
	brandRepository "eCommerce/internal/domain/brand/repository"
	categoryRepository "eCommerce/internal/domain/category/repository"
	"eCommerce/internal/domain/product/handler"
	"eCommerce/internal/domain/product/repository"
	"eCommerce/internal/domain/product/service"
	sectionRepository "eCommerce/internal/domain/section/repository"
	"eCommerce/pkg/config"
	"gorm.io/gorm"
)

var (
	productRepo    repository.ProductRepository
	sectionRepo    sectionRepository.SectionRepository
	categoryRepo   categoryRepository.CategoryRepository
	brandRepo      brandRepository.BrandRepository
	productService service.ProductService
	ProductHandler handler.ProductHandler
)

func InitProductRequirementCreator(db *gorm.DB, config config.Config) {
	productRepo = repository.NewProductRepository(db)
	sectionRepo = sectionRepository.NewSectionRepository(db)
	categoryRepo = categoryRepository.NewCategoryRepository(db)
	brandRepo = brandRepository.NewBrandRepository(db)
	productService = service.NewProductService(productRepo, sectionRepo, categoryRepo, brandRepo)
	ProductHandler = handler.NewProductHandler(productService, config)
}
