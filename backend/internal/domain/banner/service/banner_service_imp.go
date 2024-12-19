package service

import (
	"eCommerce/internal/domain/banner/dto"
	"eCommerce/internal/domain/banner/repository"
	"eCommerce/internal/models"
	"eCommerce/internal/utils/images"
	"eCommerce/pkg/config"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type bannerServiceImp struct {
	bannerRepo repository.BannerRepository
}

func NewBannerService(bannerRepo repository.BannerRepository) BannerService {
	return bannerServiceImp{bannerRepo: bannerRepo}
}

func (b bannerServiceImp) GetAllBanners() ([]dto.BannerResponse, error) {
	banners, err := b.bannerRepo.GetAll()
	if err != nil {
		return nil, err
	}
	bannerResponses := dto.GetAllBannerResponse(banners)
	return bannerResponses, nil
}

func (b bannerServiceImp) GetOneBanner(bannerID int) (*dto.BannerResponse, error) {
	banner, err := b.bannerRepo.GetOneByID(int(bannerID))
	if err != nil {
		return nil, err
	}
	bannerResponse := dto.GetBannerResponse(*banner)
	return &bannerResponse, nil
}

func (b bannerServiceImp) CreateBanner(ctx *fiber.Ctx, config config.Config, createRequest dto.CreateBannerRequest) error {
	// upload image
	bannerImagePath, err := images.UploadFile(ctx, "banner_image", config.FolderConfig.PublicPath, "banner-images")
	if err != nil {
		return err
	}
	newBanner := models.Banner{
		BannerImage:  *bannerImagePath,
		BannerStatus: createRequest.BannerStatus,
	}
	if err = b.bannerRepo.Create(&newBanner); err != nil {
		if errDeleteBannerImage := images.DeleteFile(*bannerImagePath); errDeleteBannerImage != nil {
			return errDeleteBannerImage
		}
		return err
	}
	return nil
}

func (b bannerServiceImp) UpdateBanner(bannerID int, updateRequest dto.UpdateBannerRequest) error {
	// get banner
	banner, err := b.bannerRepo.GetOneByID(int(bannerID))
	if err != nil {
		return errors.New("banner not found")
	}

	banner.BannerStatus = updateRequest.BannerStatus

	return b.bannerRepo.Update(int(banner.ID), banner)
}

func (b bannerServiceImp) DeleteBanner(bannerID int) error {
	banner, err := b.bannerRepo.GetOneByID(int(bannerID))
	if err != nil {
		return errors.New("something went wrong. Banner not found")
	}
	// delete banner image
	if errDeleteBannerImage := images.DeleteFile(banner.BannerImage); errDeleteBannerImage != nil {
		return errDeleteBannerImage
	}
	// delete banner
	return b.bannerRepo.Delete(int(banner.ID))
}
