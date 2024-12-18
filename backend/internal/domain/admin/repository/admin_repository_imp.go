package repository

import (
	"eCommerce/internal/domain/admin/dto"
	"eCommerce/internal/models"
	"errors"
	"gorm.io/gorm"
)

type adminRepositoryImp struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return adminRepositoryImp{db: db}
}

func (a adminRepositoryImp) GetAll() ([]dto.AdminResponse, error) {
	var admins []dto.AdminResponse
	if err := a.db.Select("id, username, phone_number, admin_role, last_login, created_at, updated_at").Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (a adminRepositoryImp) GetById(adminID int) (*dto.AdminResponse, error) {
	var admin dto.AdminResponse
	if err := a.db.Select("id, username, phone_number, admin_role, last_login, created_at, updated_at").
		First(&admin, "id = ?", adminID).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (a adminRepositoryImp) Create(createAdmin models.Admin) error {
	if err := a.db.Create(&createAdmin).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("This phone number is already exist")
		}
		return err
	}
	return nil
}

func (a adminRepositoryImp) Update(adminID int, updateAdmin models.Admin) error {
	if err := a.db.Model(&models.Admin{}).Where("id=?", adminID).Updates(&updateAdmin).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("This phone number is already exist")
		}
		return err
	}
	return nil
}

func (a adminRepositoryImp) Delete(adminID int) error {
	if err := a.db.Delete(models.Admin{ID: uint64(adminID)}).Error; err != nil {
		return err
	}
	return nil
}

// helper functions

func (a adminRepositoryImp) VerifyPhoneNumber(phoneNumber string) (bool, error) {
	var admin models.Admin
	if err := a.db.Where("phone_number = ?", phoneNumber).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Phone number does not exist
			return false, nil
		}
		// Other errors
		return false, err
	}
	// Phone number exists
	return true, nil
}

func (a adminRepositoryImp) GetByPhoneNumberWithID(adminID int, phoneNumber string) (*dto.AdminResponse, error) {
	var admin dto.AdminResponse
	if err := a.db.Where("id != ? AND phone_number = ?", adminID, phoneNumber).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
