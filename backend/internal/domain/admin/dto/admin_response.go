package dto

import (
	"eCommerce/internal/constants"
	"time"
)

type AdminResponse struct {
	ID          uint64               `json:"id"`
	Username    string               `json:"username"`
	PhoneNumber string               `json:"phone_number"`
	AdminRole   constants.ADMIN_ROLE `json:"admin_role"`
	LastLogin   time.Time            `json:"last_login"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}
