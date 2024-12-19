package dto

import "eCommerce/internal/constants"

type CreateBrandRequest struct {
	BrandNameTk string           `json:"brand_name_tk"`
	BrandNameRu string           `json:"brand_name_ru"`
	BrandNameEn string           `json:"brand_name_en"`
	BrandStatus constants.STATUS `json:"brand_status"`
	BrandIcon   string           `json:"brand_icon"`
}

type UpdateBrandRequest struct {
	BrandNameTk string           `json:"brand_name_tk"`
	BrandNameRu string           `json:"brand_name_ru"`
	BrandNameEn string           `json:"brand_name_en"`
	BrandStatus constants.STATUS `json:"brand_status"`
	BrandIcon   string           `json:"brand_icon"`
}
