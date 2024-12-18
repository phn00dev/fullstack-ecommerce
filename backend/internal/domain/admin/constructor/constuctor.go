package constructor

import (
	"eCommerce/internal/domain/admin/handler"
	"eCommerce/internal/domain/admin/repository"
	"eCommerce/internal/domain/admin/service"
	"gorm.io/gorm"
)

var (
	adminRepo    repository.AdminRepository
	adminService service.AdminService
	AdminHandler handler.AdminHandler
)

func InitAdminRequirementCreator(db *gorm.DB) {
	adminRepo = repository.NewAdminRepository(db)
	adminService = service.NewAdminService(adminRepo)
	AdminHandler = handler.NewAdminHandler(adminService)
}
