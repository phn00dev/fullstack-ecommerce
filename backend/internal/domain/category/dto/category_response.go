package dto

import (
	"eCommerce/internal/constants"
	"eCommerce/internal/models"
)

type CategoryResponse struct {
	ID             uint             `json:"id"`
	CategoryNameTk string           `json:"category_name_tk"`
	CategoryNameRu string           `json:"category_name_ru"`
	CategoryNameEn string           `json:"category_name_en"`
	CategorySlug   string           `json:"category_slug"`
	CategoryIcon   string           `json:"category_icon"`
	CategoryStatus constants.STATUS `json:"category_status"`
	CreatedAt      string           `json:"created_at"`
	UpdatedAt      string           `json:"updated_at"`
	DeletedAt      string           `json:"deleted_at"`
}

func GetAllCategoryResponse(categories []models.Category) []CategoryResponse {
	var categoriesResponse []CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, GetCategoryResponse(category))
	}
	return categoriesResponse
}

func GetCategoryResponse(category models.Category) CategoryResponse {
	return CategoryResponse{
		ID:             category.ID,
		CategoryNameTk: category.CategoryNameTk,
		CategoryNameRu: category.CategoryNameRu,
		CategoryNameEn: category.CategoryNameEn,
		CategorySlug:   category.CategorySlug,
		CategoryIcon:   category.CategoryIcon,
		CategoryStatus: category.CategoryStatus,
		CreatedAt:      category.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:      category.UpdatedAt.Format("02-01-2006 15:04:05"),
		DeletedAt:      category.DeletedAt.Format("02-01-2006 15:04:05"),
	}
}
