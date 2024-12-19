package service

import (
	"eCommerce/internal/domain/brand/dto"
	"eCommerce/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type BrandService interface {
	GetAllBrands() ([]dto.BrandResponse, error)
	GetOneBrandByID(brandID int) (*dto.BrandResponse, error)
	CreateBrand(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateBrandRequest) error
	UpdateBrand(ctx *fiber.Ctx, config *config.Config, brandID int, updateRequest dto.UpdateBrandRequest) error
	DeleteBrand(brandID int) error
}
