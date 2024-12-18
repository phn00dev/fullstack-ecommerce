package service

import (
	"eCommerce/internal/domain/admin/dto"
	"eCommerce/internal/domain/admin/repository"
	"eCommerce/internal/models"
	"eCommerce/internal/utils/password"
	"errors"
)

type adminServiceImp struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(adminRepo repository.AdminRepository) AdminService {
	return adminServiceImp{adminRepo: adminRepo}
}

func (a adminServiceImp) GetAllAdmins() ([]dto.AdminResponse, error) {
	admins, err := a.adminRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (a adminServiceImp) GetAdminByID(adminID int) (*models.Admin, error) {
	admin, err := a.adminRepo.GetById(adminID)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (a adminServiceImp) CreateAdmin(createRequest dto.CreateAdminRequest) error {
	// verify phoneNumber
	verifyPhoneNumber, err := a.adminRepo.VerifyPhoneNumber(createRequest.PhoneNumber)
	if err != nil {
		return err
	}
	if verifyPhoneNumber {
		return errors.New("This phone number is already used")
	}

	hashPassword := password.HashPassword(createRequest.Password)
	createAdmin := models.Admin{
		Username:    createRequest.Username,
		PhoneNumber: createRequest.PhoneNumber,
		AdminRole:   createRequest.AdminRole,
		Password:    hashPassword, // hashed password
	}
	// create admin
	return a.adminRepo.Create(createAdmin)
}

func (a adminServiceImp) UpdateAdmin(adminID int, updateRequest dto.UpdateAdminRequest) error {
	admin, err := a.adminRepo.GetById(adminID)
	if err != nil {
		return err
	}
	if admin.ID == 0 {
		return errors.New("admin not found")
	}
	// verify phoneNumber
	verifyPhoneNumberByID, err := a.adminRepo.GetByPhoneNumberWithID(int(admin.ID), admin.PhoneNumber)
	if err != nil {
		return err
	}
	if verifyPhoneNumberByID.ID != 0 {
		return errors.New("phone number is already in use by another admin")
	}
	// update admin
	admin.Username = updateRequest.Username
	admin.PhoneNumber = updateRequest.PhoneNumber
	admin.AdminRole = updateRequest.AdminRole

	return a.adminRepo.Update(int(admin.ID), *admin)

}

func (a adminServiceImp) DeleteAdmin(adminID int) error {
	// get admin before delete
	admin, err := a.adminRepo.GetById(adminID)
	if err != nil {
		return err
	}
	if admin.ID == 0 {
		return errors.New("admin not found")
	}
	// delete admin
	if err := a.adminRepo.Delete(adminID); err != nil {
		return err
	}
	return nil
}
