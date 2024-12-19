package handler

import (
	"eCommerce/internal/domain/brand/dto"
	"eCommerce/internal/domain/brand/service"
	"eCommerce/internal/utils/response"
	"eCommerce/internal/utils/validate"
	"eCommerce/pkg/config"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type brandHandlerImp struct {
	brandService service.BrandService
	config       config.Config
}

func NewBrandHandler(service service.BrandService, config config.Config) BrandHandler {
	return brandHandlerImp{
		brandService: service,
		config:       config,
	}
}

func (b brandHandlerImp) GetAll(ctx *fiber.Ctx) error {
	brands, err := b.brandService.GetAllBrands()
	if err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "brands not found", err)
	}
	return response.Success(ctx, fiber.StatusOK, "brands", brands)
}

func (b brandHandlerImp) GetOne(ctx *fiber.Ctx) error {
	brandIdStr := ctx.Params("brandID")
	brandId, _ := strconv.Atoi(brandIdStr)

	brand, err := b.brandService.GetOneBrandByID(brandId)
	if err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "brand not found", err)
	}
	return response.Success(ctx, fiber.StatusOK, "brand", brand)
}

func (b brandHandlerImp) Create(ctx *fiber.Ctx) error {
	var createRequest dto.CreateBrandRequest
	if err := ctx.BodyParser(&createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body", err)
	}
	// validate data
	if err := validate.ValidateStruct(createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body", err)
	}
	// create brand
	if err := b.brandService.CreateBrand(ctx, &b.config, createRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "brand creation error", err)
	}
	return response.Success(ctx, fiber.StatusOK, "brand created successfully", nil)
}

func (b brandHandlerImp) Update(ctx *fiber.Ctx) error {
	var updateRequest dto.UpdateBrandRequest
	brandIdStr := ctx.Params("brandID")
	brandId, _ := strconv.Atoi(brandIdStr)
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body", err)
	}
	// validate data
	if err := validate.ValidateStruct(updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body", err)
	}
	// create brand
	if err := b.brandService.UpdateBrand(ctx, &b.config, brandId, updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "brand creation error", err)
	}
	return response.Success(ctx, fiber.StatusOK, "brand updated successfully", nil)
}

func (b brandHandlerImp) Delete(ctx *fiber.Ctx) error {
	brandIdStr := ctx.Params("brandID")
	brandId, _ := strconv.Atoi(brandIdStr)
	// delete brand
	if err := b.brandService.DeleteBrand(brandId); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "brand creation error", err)
	}
	return response.Success(ctx, fiber.StatusOK, "brand deleted successfully", nil)
}
