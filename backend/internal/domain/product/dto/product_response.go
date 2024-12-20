package dto

import (
	"eCommerce/internal/constants"
	"eCommerce/internal/models"
)

type ProductResponse struct {
	ID                        uint             `json:"id"`
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
	CreatedAt                 string           `json:"created_at"`
	UpdatedAt                 string           `json:"updated_at"`
	DeletedAt                 string           `json:"deleted_at"`
}

func GetAllProductResponses(products []models.Product) []ProductResponse {
	var productResponses []ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, GetOneProductResponse(product))
	}
	return productResponses
}

func GetOneProductResponse(product models.Product) ProductResponse {
	return ProductResponse{
		ID:                        product.ID,
		ProductNameTk:             product.ProductNameTk,
		ProductNameRu:             product.ProductNameRu,
		ProductNameEn:             product.ProductNameEn,
		ProductShortDescriptionTk: product.ProductShortDescriptionTk,
		ProductShortDescriptionRu: product.ProductShortDescriptionRu,
		ProductShortDescriptionEn: product.ProductShortDescriptionEn,
		ProductLongDescriptionTk:  product.ProductLongDescriptionTk,
		ProductLongDescriptionRu:  product.ProductLongDescriptionRu,
		ProductLongDescriptionEn:  product.ProductLongDescriptionEn,
		ProductAllSpecificationTk: product.ProductAllSpecificationTk,
		ProductAllSpecificationRu: product.ProductAllSpecificationRu,
		ProductAllSpecificationEn: product.ProductAllSpecificationEn,
		ProductPrice:              product.ProductPrice,
		ProductTotalCount:         product.ProductTotalCount,
		ProductRemainingNumber:    product.ProductRemainingNumber,
		ProductMainImageOne:       product.ProductMainImageOne,
		ProductMainImageTwo:       product.ProductMainImageTwo,
		ProductStatus:             product.ProductStatus,
		SectionID:                 int(product.SectionID),
		CategoryID:                int(product.CategoryID),
		BrandID:                   int(product.BrandID),
		CreatedAt:                 product.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:                 product.UpdatedAt.Format("02-01-2006 15:04:05"),
		DeletedAt:                 product.DeletedAt.Format("02-01-2006 15:04:05"),
	}
}
