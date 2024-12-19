package service

import (
	"eCommerce/internal/domain/banner/dto"
	"eCommerce/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type BannerService interface {
	GetAllBanners() ([]dto.BannerResponse, error)
	GetOneBanner(bannerID int) (*dto.BannerResponse, error)
	CreateBanner(ctx *fiber.Ctx, config config.Config, createRequest dto.CreateBannerRequest) error
	UpdateBanner(bannerID int, updateRequest dto.UpdateBannerRequest) error
	DeleteBanner(bannerID int) error
}
