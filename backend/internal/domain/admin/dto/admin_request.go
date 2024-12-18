package dto

import "eCommerce/internal/constants"

type CreateAdminRequest struct {
	Username        string               `json:"username"`
	PhoneNumber     string               `json:"phone_number"`
	AdminRole       constants.ADMIN_ROLE `json:"admin_role"`
	Password        string               `json:"password"`
	ConfirmPassword string               `json:"confirm_password"`
}

type UpdateAdminRequest struct {
	Username    string               `json:"username"`
	PhoneNumber string               `json:"phone_number"`
	AdminRole   constants.ADMIN_ROLE `json:"admin_role"`
}

type UpdateAdminPassword struct {
}
