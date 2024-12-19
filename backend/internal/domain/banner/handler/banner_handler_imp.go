package handler

import (
	"eCommerce/internal/domain/banner/dto"
	"eCommerce/internal/domain/banner/service"
	"eCommerce/internal/utils/response"
	"eCommerce/internal/utils/validate"
	"eCommerce/pkg/config"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type bannerHandlerImp struct {
	bannerService service.BannerService
	config        config.Config
}

func NewBannerHandler(bannerService service.BannerService, config config.Config) BannerHandler {
	return bannerHandlerImp{
		bannerService: bannerService,
		config:        config,
	}
}

func (b bannerHandlerImp) GetAll(ctx *fiber.Ctx) error {
	banners, err := b.bannerService.GetAllBanners()
	if err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "banners not found", err)
	}
	return response.Success(ctx, fiber.StatusOK, "banners", banners)
}

func (b bannerHandlerImp) GetOne(ctx *fiber.Ctx) error {
	bannerIDStr := ctx.Params("bannerID")
	bannerID, _ := strconv.Atoi(bannerIDStr)
	banner, err := b.bannerService.GetOneBanner(bannerID)
	if err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "banner not found", err)
	}
	return response.Success(ctx, fiber.StatusOK, "banner", banner)
}

func (b bannerHandlerImp) Create(ctx *fiber.Ctx) error {
	var createRequest dto.CreateBannerRequest
	if err := ctx.BodyParser(&createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body", err)
	}
	if err := validate.ValidateStruct(&createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body", err)
	}
	// create banner
	if err := b.bannerService.CreateBanner(ctx, b.config, createRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "banner creation failed", err)
	}
	return response.Success(ctx, fiber.StatusOK, "banner created successfully", nil)
}

func (b bannerHandlerImp) Update(ctx *fiber.Ctx) error {
	bannerIDStr := ctx.Params("bannerID")
	bannerID, _ := strconv.Atoi(bannerIDStr)
	var updateRequest dto.UpdateBannerRequest
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body", err)
	}
	if err := validate.ValidateStruct(&updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "invalid request body", err)
	}

	if err := b.bannerService.UpdateBanner(bannerID, updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "banner update failed", err)
	}
	return response.Success(ctx, fiber.StatusOK, "banner updated successfully", nil)
}

func (b bannerHandlerImp) Delete(ctx *fiber.Ctx) error {
	bannerIDStr := ctx.Params("bannerID")
	bannerID, _ := strconv.Atoi(bannerIDStr)
	if err := b.bannerService.DeleteBanner(bannerID); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "banner delete failed", err)
	}
	return response.Success(ctx, fiber.StatusOK, "banner deleted successfully", nil)
}
