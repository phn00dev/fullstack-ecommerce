package dto

import (
	"eCommerce/internal/constants"
)

type CreateProductRequest struct {
	ProductNameTk             string           `json:"product_name_tk"`
	ProductNameRu             string           `json:"product_name_ru"`
	ProductNameEn             string           `json:"product_name_en"`
	ProductShortDescriptionTk string           `json:"product_short_description_tk"`
	ProductShortDescriptionRu string           `json:"product_short_description_ru"`
	ProductShortDescriptionEn string           `json:"product_short_description_en"`
	ProductLongDescriptionTk  string           `json:"product_long_description_tk"`
	ProductLongDescriptionRu  string           `json:"product_long_description_ru"`
	ProductLongDescriptionEn  string           `json:"product_long_description_en"`
	ProductAllSpecificationTk string           `json:"product_all_specifications_tk"`
	ProductAllSpecificationRu string           `json:"product_all_specifications_ru"`
	ProductAllSpecificationEn string           `json:"product_all_specifications_en"`
	ProductPrice              float64          `json:"product_price"`
	ProductTotalCount         int              `json:"product_count"`
	ProductRemainingNumber    int              `json:"product_remaining_number"`
	ProductMainImageOne       string           `json:"product_main_image_one"`
	ProductMainImageTwo       string           `json:"product_main_image_two"`
	ProductStatus             constants.STATUS `json:"product_status"`
	SectionID                 int              `json:"section_id"`
	CategoryID                int              `json:"category_id"`
	BrandID                   int              `json:"brand_id"`
}

type UpdateProductRequest struct {
	ProductNameTk             string           `json:"product_name_tk"`
	ProductNameRu             string           `json:"product_name_ru"`
	ProductNameEn             string           `json:"product_name_en"`
	ProductShortDescriptionTk string           `json:"product_short_description_tk"`
	ProductShortDescriptionRu string           `json:"product_short_description_ru"`
	ProductShortDescriptionEn string           `json:"product_short_description_en"`
	ProductLongDescriptionTk  string           `json:"product_long_description_tk"`
	ProductLongDescriptionRu  string           `json:"product_long_description_ru"`
	ProductLongDescriptionEn  string           `json:"product_long_description_en"`
	ProductAllSpecificationTk string           `json:"product_all_specifications_tk"`
	ProductAllSpecificationRu string           `json:"product_all_specifications_ru"`
	ProductAllSpecificationEn string           `json:"product_all_specifications_en"`
	ProductPrice              float64          `json:"product_price"`
	ProductTotalCount         int              `json:"product_count"`
	ProductRemainingNumber    int              `json:"product_remaining_number"`
	ProductMainImageOne       string           `json:"product_main_image_one"`
	ProductMainImageTwo       string           `json:"product_main_image_two"`
	ProductStatus             constants.STATUS `json:"product_status"`
	SectionID                 int              `json:"section_id"`
	CategoryID                int              `json:"category_id"`
	BrandID                   int              `json:"brand_id"`
}
