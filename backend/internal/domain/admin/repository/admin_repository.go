package repository

import (
	"eCommerce/internal/domain/admin/dto"
	"eCommerce/internal/models"
)

type AdminRepository interface {
	GetAll() ([]dto.AdminResponse, error)
	GetById(adminID int) (*models.Admin, error)
	Create(createAdmin models.Admin) error
	Update(adminID int, updateAdmin models.Admin) error
	Delete(adminID int) error

	// helper functions
	VerifyPhoneNumber(phoneNumber string) (bool, error)
	GetByPhoneNumberWithID(adminID int, phoneNumber string) (*dto.AdminResponse, error)
}
