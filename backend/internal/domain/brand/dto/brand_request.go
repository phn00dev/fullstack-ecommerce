package dto

import "eCommerce/internal/constants"

type CreateBrandRequest struct {
	BrandName   string           `json:"brand_name"`
	BrandStatus constants.STATUS `json:"brand_status"`
	BrandIcon   string           `json:"brand_icon"`
}

type UpdateBrandRequest struct {
	BrandName   string           `json:"brand_name"`
	BrandStatus constants.STATUS `json:"brand_status"`
	BrandIcon   string           `json:"brand_icon"`
}
