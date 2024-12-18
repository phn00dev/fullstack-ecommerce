package service

import (
	"eCommerce/internal/domain/admin/dto"
	"eCommerce/internal/models"
)

type AdminService interface {
	GetAllAdmins() ([]dto.AdminResponse, error)
	GetAdminByID(adminID int) (*models.Admin, error)
	CreateAdmin(createRequest dto.CreateAdminRequest) error
	UpdateAdmin(adminID int, updateRequest dto.UpdateAdminRequest) error
	DeleteAdmin(adminID int) error
}
