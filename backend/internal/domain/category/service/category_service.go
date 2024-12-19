package service

import (
	"eCommerce/internal/domain/category/dto"
	"eCommerce/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type CategoryService interface {
	GetAllCategories() ([]dto.CategoryResponse, error)
	GetCategoryById(categoryId int) (*dto.CategoryResponse, error)
	CreateCategory(ctx *fiber.Ctx, config config.Config, createRequest dto.CreateCategoryRequest) error
	UpdateCategory(ctx *fiber.Ctx, config config.Config, categoryId int, updateRequest dto.UpdateCategoryRequest) error
	DeleteCategory(categoryId int) error
}
