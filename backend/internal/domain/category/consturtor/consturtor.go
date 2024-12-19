package consturtor

import (
	"eCommerce/internal/domain/category/handler"
	"eCommerce/internal/domain/category/repository"
	"eCommerce/internal/domain/category/service"
	sectionRepository "eCommerce/internal/domain/section/repository"
	"eCommerce/pkg/config"
	"gorm.io/gorm"
)

var (
	categoryRepo    repository.CategoryRepository
	sectionRepo     sectionRepository.SectionRepository
	categoryService service.CategoryService
	CategoryHandler handler.CategoryHandler
)

func InitCategoryRequirementCreator(db *gorm.DB, config config.Config) {
	categoryRepo = repository.NewCategoryRepository(db)
	sectionRepo = sectionRepository.NewSectionRepository(db)
	categoryService = service.NewCategoryService(categoryRepo, sectionRepo)
	CategoryHandler = handler.NewCategoryHandler(categoryService, config)
}
