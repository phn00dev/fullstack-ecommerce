package dto

import (
	"eCommerce/internal/constants"
	"eCommerce/internal/models"
)

type BannerResponse struct {
	ID           uint             `json:"id"`
	BannerImage  string           `json:"banner_image"`
	BannerStatus constants.STATUS `json:"banner_status"`
	CreatedAt    string           `json:"created_at"`
	UpdatedAt    string           `json:"updated_at"`
	DeletedAt    string           `json:"deleted_at"`
}

func GetAllBannerResponse(banners []models.Banner) []BannerResponse {
	var bannerResponse []BannerResponse
	for _, banner := range banners {
		bannerResponse = append(bannerResponse, GetBannerResponse(banner))
	}
	return bannerResponse
}

func GetBannerResponse(banner models.Banner) BannerResponse {
	return BannerResponse{
		ID:           banner.ID,
		BannerImage:  banner.BannerImage,
		BannerStatus: banner.BannerStatus,
		CreatedAt:    banner.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:    banner.UpdatedAt.Format("02-01-2006 15:04:05"),
		DeletedAt:    banner.DeletedAt.Format("02-01-2006 15:04:05"),
	}
}
