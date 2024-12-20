package service

import (
	"eCommerce/internal/domain/product/dto"
	"eCommerce/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type ProductService interface {
	GetAllProducts() ([]dto.ProductResponse, error)
	GetOneProductById(productId int) (*dto.ProductResponse, error)
	CreateProduct(ctx *fiber.Ctx, config config.Config, createRequest dto.CreateProductRequest) error
	UpdateProduct(ctx *fiber.Ctx, config config.Config, productID int, updateRequest dto.UpdateProductRequest) error
	DeleteProduct(productID int) error
}
