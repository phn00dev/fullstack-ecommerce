package dto

import "eCommerce/internal/constants"

type CreateCategoryRequest struct {
	CategoryNameTk string           `json:"category_name_tk"`
	CategoryNameRu string           `json:"category_name_ru"`
	CategoryNameEn string           `json:"category_name_en"`
	CategoryIcon   string           `json:"category_icon"`
	CategoryStatus constants.STATUS `json:"category_status"`
	SectionID      int              `json:"section_id"`
}

type UpdateCategoryRequest struct {
	CategoryNameTk string           `json:"category_name_tk"`
	CategoryNameRu string           `json:"category_name_ru"`
	CategoryNameEn string           `json:"category_name_en"`
	CategoryIcon   string           `json:"category_icon"`
	CategoryStatus constants.STATUS `json:"category_status"`
	SectionID      int              `json:"section_id"`
}
