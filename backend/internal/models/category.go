package models

import (
	"eCommerce/internal/constants"
	"time"
)

type Category struct {
	ID             uint             `json:"id"`
	CategoryNameTk string           `json:"category_name_tk"`
	CategoryNameRu string           `json:"category_name_ru"`
	CategoryNameEn string           `json:"category_name_en"`
	CategorySlug   string           `json:"category_slug"`
	CategoryIcon   string           `json:"category_icon"`
	CategoryStatus constants.STATUS `json:"category_status"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
	DeletedAt      time.Time        `json:"deleted_at"`
}
