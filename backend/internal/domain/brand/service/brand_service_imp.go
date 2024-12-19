package service

import (
	"eCommerce/internal/domain/brand/dto"
	"eCommerce/internal/domain/brand/repository"
	"eCommerce/internal/models"
	"eCommerce/internal/utils/images"
	"eCommerce/pkg/config"
	"errors"
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
	// check brand name
	checkBrandName, err := b.brandRepo.GetOneBrandByBrandName(createRequest.BrandName)
	if err != nil {
		return err
	}

	if checkBrandName.ID != 0 {
		return errors.New("Brand name already exists")
	}
	// upload image
	brandIconPath, err := images.UploadFile(ctx, "brand_icon", config.FolderConfig.PublicPath, "brand-icons")
	if err != nil {
		return err
	}
	newBrand := models.Brand{
		BrandName:   createRequest.BrandName,
		BrandSlug:   slug.Make(createRequest.BrandName),
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
	// image barlag

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
		updateRequest.BrandIcon = *newBrandIconPath
	}
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
