package dto

import (
	"eCommerce/internal/constants"
	"eCommerce/internal/models"
)

type BrandResponse struct {
	ID          uint             `json:"id"`
	BrandName   string           `json:"brand_name"`
	BrandSlug   string           `json:"brand_slug"`
	BrandIcon   string           `json:"brand_icon"`
	BrandStatus constants.STATUS `json:"brand_status"`
	CreatedAt   string           `json:"created_at"`
	UpdatedAt   string           `json:"updated_at"`
	DeletedAt   string           `json:"deleted_at"`
}

func GetAllBrandResponse(brands []models.Brand) []BrandResponse {
	var brandResponses []BrandResponse
	for _, brand := range brands {
		brandResponses = append(brandResponses, GetBrandResponse(brand))
	}
	return brandResponses
}

func GetBrandResponse(brand models.Brand) BrandResponse {
	return BrandResponse{
		ID:          brand.ID,
		BrandName:   brand.BrandName,
		BrandSlug:   brand.BrandSlug,
		BrandIcon:   brand.BrandIcon,
		BrandStatus: brand.BrandStatus,
		CreatedAt:   brand.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:   brand.UpdatedAt.Format("02-01-2006 15:04:05"),
		DeletedAt:   brand.DeletedAt.Format("02-01-2006 15:04:05"),
	}
}
