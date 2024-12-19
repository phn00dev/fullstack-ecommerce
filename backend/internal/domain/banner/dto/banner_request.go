package dto

import "eCommerce/internal/constants"

type CreateBannerRequest struct {
	BannerImage  string           `json:"banner_image"`
	BannerStatus constants.STATUS `json:"banner_status"`
}

type UpdateBannerRequest struct {
	BannerStatus constants.STATUS `json:"banner_status"`
}
