package models

import (
	"eCommerce/internal/constants"
	"time"
)

type Brand struct {
	ID          uint             `json:"id"`
	BrandName   string           `json:"brand_name"`
	BrandSlug   string           `json:"brand_slug"`
	BrandIcon   string           `json:"brand_icon"`
	BrandStatus constants.STATUS `json:"brand_status"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	DeletedAt   time.Time        `json:"deleted_at"`
}
