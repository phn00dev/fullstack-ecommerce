package models

import (
	"eCommerce/internal/constants"
	"time"
)

type Admin struct {
	ID          uint64               `json:"id"`
	Username    string               `json:"username"`
	PhoneNumber string               `json:"phone_number"`
	AdminRole   constants.ADMIN_ROLE `json:"admin_role"`
	Password    string               `json:"password"`
	LastLogin   time.Time            `json:"last_login"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}
