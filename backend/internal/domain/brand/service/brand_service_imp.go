package service

import (
	"eCommerce/internal/domain/brand/dto"
	"eCommerce/internal/domain/brand/repository"
	"eCommerce/internal/models"
	"eCommerce/internal/utils/images"
	"eCommerce/pkg/config"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

type brandServiceImp struct {
	brandRepo repository.BrandRepository
}

func NewBrandService(brandRepo repository.BrandRepository) BrandService {
	return brandServiceImp{
		brandRepo: brandRepo,
	}
}

func (b brandServiceImp) GetAllBrands() ([]dto.BrandResponse, error) {
	brands, err := b.brandRepo.GetAll()
	if err != nil {
		return nil, err
	}
	brandResponses := dto.GetAllBrandResponse(brands)
	return brandResponses, nil
}

func (b brandServiceImp) GetOneBrandByID(brandID int) (*dto.BrandResponse, error) {
	brand, err := b.brandRepo.GetOneByID(brandID)
	if err != nil {
		return nil, err
	}
	brandResponse := dto.GetBrandResponse(*brand)
	return &brandResponse, nil
}

func (b brandServiceImp) CreateBrand(ctx *fiber.Ctx, config *config.Config, createRequest dto.CreateBrandRequest) error {
	// check brand name tk
	checkBrandNameTK, err := b.brandRepo.CheckBrandNameTk(createRequest.BrandNameTk)
	if err != nil {
		return fmt.Errorf("error checking brand name TK: %w", err)
	}
	if checkBrandNameTK {
		return errors.New("brand name TK already exists")
	}
	// check brand name ru
	checkBrandNameRu, err := b.brandRepo.CheckBrandNameRu(createRequest.BrandNameRu)
	if err != nil {
		return fmt.Errorf("error checking brand name Ru: %w", err)
	}
	if checkBrandNameRu {
		return errors.New("brand name TK already exists")
	}
	// check brand name en
	checkBrandNameEn, err := b.brandRepo.CheckBrandNameEn(createRequest.BrandNameEn)
	if err != nil {
		return fmt.Errorf("error checking brand name En: %w", err)
	}
	if checkBrandNameEn {
		return errors.New("brand name TK already exists")
	}

	// upload image
	brandIconPath, err := images.UploadFile(ctx, "brand_icon", config.FolderConfig.PublicPath, "brand-icons")
	if err != nil {
		return err
	}
	newBrand := models.Brand{
		BrandNameTk: createRequest.BrandNameTk,
		BrandNameRu: createRequest.BrandNameRu,
		BrandNameEn: createRequest.BrandNameEn,
		BrandSlug:   slug.Make(createRequest.BrandNameEn),
		BrandIcon:   *brandIconPath,
		BrandStatus: createRequest.BrandStatus,
	}

	if err = b.brandRepo.Create(newBrand); err != nil {
		// image delete
		if errDeleteBrandIcon := images.DeleteFile(*brandIconPath); errDeleteBrandIcon != nil {
			return errDeleteBrandIcon
		}
		return err
	}
	return nil

}

func (b brandServiceImp) UpdateBrand(ctx *fiber.Ctx, config *config.Config, brandID int, updateRequest dto.UpdateBrandRequest) error {
	brand, err := b.brandRepo.GetOneByID(brandID)
	if err != nil {
		return err
	}
	// check image
	file, _ := ctx.FormFile("brand_icon")
	if file != nil {
		// delete old icon
		if errDeleteOldBrandIcon := images.DeleteFile(brand.BrandIcon); errDeleteOldBrandIcon != nil {
			return errDeleteOldBrandIcon
		}
		// upload new brand icon
		newBrandIconPath, errUpload := images.UploadFile(ctx, "brand_icon", config.FolderConfig.PublicPath, "brand-icons")
		if errUpload != nil {
			return errUpload
		}
		brand.BrandIcon = *newBrandIconPath
	}

	brand.BrandNameTk = updateRequest.BrandNameTk
	brand.BrandNameRu = updateRequest.BrandNameRu
	brand.BrandNameEn = updateRequest.BrandNameEn
	brand.BrandSlug = slug.Make(updateRequest.BrandNameEn)
	brand.BrandStatus = updateRequest.BrandStatus

	// update brand
	return b.brandRepo.Update(int(brand.ID), brand)
}

func (b brandServiceImp) DeleteBrand(brandID int) error {
	// get brand
	brand, err := b.brandRepo.GetOneByID(brandID)
	if err != nil {
		return errors.New("something went wrong. Brand not found")
	}
	// delete brand image
	if errDeleteBrandIcon := images.DeleteFile(brand.BrandIcon); errDeleteBrandIcon != nil {
		return errDeleteBrandIcon
	}
	// delete brand
	return b.brandRepo.Delete(int(brand.ID))
}
