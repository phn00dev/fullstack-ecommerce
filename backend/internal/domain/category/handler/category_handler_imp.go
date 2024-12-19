package handler

import (
	"eCommerce/internal/domain/category/dto"
	"eCommerce/internal/domain/category/service"
	"eCommerce/internal/utils/response"
	"eCommerce/internal/utils/validate"
	"eCommerce/pkg/config"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type categoryHandlerImp struct {
	categoryService service.CategoryService
	config          config.Config
}

func NewCategoryHandler(categoryService service.CategoryService, config config.Config) CategoryHandler {
	return categoryHandlerImp{categoryService: categoryService, config: config}
}

func (c categoryHandlerImp) GetAll(ctx *fiber.Ctx) error {
	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "Categories not found", err)
	}
	return response.Success(ctx, fiber.StatusOK, "categories", categories)
}

func (c categoryHandlerImp) GetOne(ctx *fiber.Ctx) error {
	categoryIdStr := ctx.Params("categoryId")
	categoryId, _ := strconv.Atoi(categoryIdStr)
	category, err := c.categoryService.GetCategoryById(categoryId)
	if err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "Category not found", err)
	}
	return response.Success(ctx, fiber.StatusOK, "category", category)
}

func (c categoryHandlerImp) Create(ctx *fiber.Ctx) error {
	var createRequest dto.CreateCategoryRequest
	if err := ctx.BodyParser(&createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}
	// validate data
	if err := validate.ValidateStruct(createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}
	// create category
	if err := c.categoryService.CreateCategory(ctx, c.config, createRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Error creating category", err)
	}
	return response.Success(ctx, fiber.StatusOK, "category created successfully", nil)
}

func (c categoryHandlerImp) Update(ctx *fiber.Ctx) error {
	var updateRequest dto.UpdateCategoryRequest
	categoryIdStr := ctx.Params("categoryId")
	categoryId, _ := strconv.Atoi(categoryIdStr)
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}
	// validate data
	if err := validate.ValidateStruct(updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}
	// update category
	if err := c.categoryService.UpdateCategory(ctx, c.config, categoryId, updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Error updating category", err)
	}
	return response.Success(ctx, fiber.StatusOK, "category updated successfully", nil)
}

func (c categoryHandlerImp) Delete(ctx *fiber.Ctx) error {
	categoryIdStr := ctx.Params("categoryId")
	categoryId, _ := strconv.Atoi(categoryIdStr)

	// delete category
	if err := c.categoryService.DeleteCategory(categoryId); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Error deleting category", err)
	}
	return response.Success(ctx, fiber.StatusOK, "category deleted successfully", nil)
}
