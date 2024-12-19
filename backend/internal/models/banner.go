package models

import (
	"eCommerce/internal/constants"
	"time"
)

type Banner struct {
	ID           uint             `json:"id"`
	BannerImage  string           `json:"banner_image"`
	BannerStatus constants.STATUS `json:"banner_status"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
	DeletedAt    time.Time        `json:"deleted_at"`
}
