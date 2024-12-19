package service

import (
	"eCommerce/internal/domain/category/dto"
	"eCommerce/internal/domain/category/repository"
	sectionRepository "eCommerce/internal/domain/section/repository"
	"eCommerce/internal/models"
	"eCommerce/internal/utils/images"
	"eCommerce/pkg/config"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

type categoryServiceImp struct {
	categoryRepo repository.CategoryRepository
	sectionRepo  sectionRepository.SectionRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository,
	sectionRepo sectionRepository.SectionRepository) CategoryService {
	return categoryServiceImp{
		categoryRepo: categoryRepo,
		sectionRepo:  sectionRepo,
	}
}

func (c categoryServiceImp) GetAllCategories() ([]dto.CategoryResponse, error) {
	categories, err := c.categoryRepo.GetAll()
	if err != nil {
		return nil, err
	}
	categoryResponses := dto.GetAllCategoryResponse(categories)
	return categoryResponses, nil
}

func (c categoryServiceImp) GetCategoryById(categoryId int) (*dto.CategoryResponse, error) {
	category, err := c.categoryRepo.GetOne(categoryId)
	if err != nil {
		return nil, err
	}
	categoryResponse := dto.GetCategoryResponse(*category)
	return &categoryResponse, nil
}

func (c categoryServiceImp) CreateCategory(ctx *fiber.Ctx, config config.Config, createRequest dto.CreateCategoryRequest) error {
	// get section
	section, err := c.sectionRepo.GetOneById(createRequest.SectionID)
	if err != nil {
		return err
	}
	if section.ID == 0 {
		return errors.New("section id is 0 or section not found")
	}

	// check category name tk
	checkCategoryNameTk, err := c.categoryRepo.CheckCategoryNameTk(createRequest.CategoryNameTk)
	if err != nil {
		return fmt.Errorf("error checking category name TK: %w", err)
	}
	if checkCategoryNameTk {
		return errors.New("category name TK already exists")
	}

	// check category name ru
	checkCategoryNameRu, err := c.categoryRepo.CheckCategoryNameRu(createRequest.CategoryNameTk)
	if err != nil {
		return fmt.Errorf("error checking category name Ru: %w", err)
	}
	if checkCategoryNameRu {
		return errors.New("category name Ru already exists")
	}
	// check category name en
	checkCategoryNameEn, err := c.categoryRepo.CheckCategoryNameEn(createRequest.CategoryNameEn)
	if err != nil {
		return fmt.Errorf("error checking category name En: %w", err)
	}
	if checkCategoryNameEn {
		return errors.New("category name En already exists")
	}

	//  upload category image
	categoryIconPath, err := images.UploadFile(ctx, "category_icon", config.FolderConfig.PublicPath, "category-icons")
	if err != nil {
		return err
	}
	newCategory := models.Category{
		CategoryNameTk: createRequest.CategoryNameTk,
		CategoryNameRu: createRequest.CategoryNameRu,
		CategoryNameEn: createRequest.CategoryNameEn,
		CategorySlug:   slug.Make(createRequest.CategoryNameEn),
		CategoryIcon:   *categoryIconPath,
		CategoryStatus: createRequest.CategoryStatus,
		SectionID:      section.ID,
	}
	if err = c.categoryRepo.Create(newCategory); err != nil {
		if errDeleteCategoryImage := images.DeleteFile(*categoryIconPath); errDeleteCategoryImage != nil {
			return errDeleteCategoryImage
		}
		return err
	}
	return nil
}

func (c categoryServiceImp) UpdateCategory(ctx *fiber.Ctx, config config.Config, categoryId int, updateRequest dto.UpdateCategoryRequest) error {
	// get category
	category, err := c.categoryRepo.GetOne(categoryId)
	if err != nil {
		return err
	}
	// get section
	section, err := c.sectionRepo.GetOneById(updateRequest.SectionID)
	if err != nil {
		return err
	}
	if section.ID == 0 {
		return errors.New("section id is 0 or section not found")
	}
	// check category icon
	file, _ := ctx.FormFile("category_icon")
	if file != nil {
		// delete old category icon
		if errDeleteOldCategoryIcon := images.DeleteFile(category.CategoryIcon); errDeleteOldCategoryIcon != nil {
			return errDeleteOldCategoryIcon
		}
		// upload new category icon
		newCategoryIconPath, errUpload := images.UploadFile(ctx, "category_icon", config.FolderConfig.PublicPath, "category-icons")
		if errUpload != nil {
			return errUpload
		}
		category.CategoryIcon = *newCategoryIconPath
	}
	// update data
	category.CategoryNameTk = updateRequest.CategoryNameTk
	category.CategoryNameRu = updateRequest.CategoryNameRu
	category.CategoryNameEn = updateRequest.CategoryNameEn
	category.CategorySlug = slug.Make(updateRequest.CategoryNameEn)
	category.CategoryStatus = updateRequest.CategoryStatus
	category.SectionID = uint(updateRequest.SectionID)
	return c.categoryRepo.Update(int(category.ID), *category)
}

func (c categoryServiceImp) DeleteCategory(categoryId int) error {
	category, err := c.categoryRepo.GetOne(categoryId)
	if err != nil {
		return err
	}
	// delete category icon
	if errDeleteCategoryIcon := images.DeleteFile(category.CategoryIcon); errDeleteCategoryIcon != nil {
		return errDeleteCategoryIcon
	}
	return c.categoryRepo.Delete(int(category.ID))
}
