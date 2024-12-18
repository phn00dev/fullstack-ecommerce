package constructor

import (
	"eCommerce/internal/domain/section/handler"
	"eCommerce/internal/domain/section/repository"
	"eCommerce/internal/domain/section/service"
	"eCommerce/pkg/config"
	"gorm.io/gorm"
)

var (
	sectionRepo    repository.SectionRepository
	sectionService service.SectionService
	SectionHandler handler.SectionHandler
)

func InitSectionRequirementCreator(db *gorm.DB, config config.Config) {
	sectionRepo = repository.NewSectionRepository(db)
	sectionService = service.NewSectionService(sectionRepo)
	SectionHandler = handler.NewSectionHandler(sectionService, config)
}
