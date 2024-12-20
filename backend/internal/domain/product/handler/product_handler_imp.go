package handler

import (
	"eCommerce/internal/domain/product/dto"
	"eCommerce/internal/domain/product/service"
	"eCommerce/internal/utils/response"
	"eCommerce/internal/utils/validate"
	"eCommerce/pkg/config"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type productHandlerImp struct {
	productService service.ProductService
	config         config.Config
}

func NewProductHandler(productService service.ProductService, config config.Config) ProductHandler {
	return productHandlerImp{
		productService: productService,
		config:         config,
	}
}

func (p productHandlerImp) GetAll(ctx *fiber.Ctx) error {
	products, err := p.productService.GetAllProducts()
	if err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "Products not found", err)
	}
	return response.Success(ctx, fiber.StatusOK, "Products", products)
}

func (p productHandlerImp) GetOne(ctx *fiber.Ctx) error {
	productIdStr := ctx.Params("productId")
	productId, _ := strconv.Atoi(productIdStr)
	product, err := p.productService.GetOneProductById(productId)
	if err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "Product not found", err)
	}
	return response.Success(ctx, fiber.StatusOK, "Product", product)
}

func (p productHandlerImp) Create(ctx *fiber.Ctx) error {
	var createRequest dto.CreateProductRequest
	if err := ctx.BodyParser(&createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}
	// validate data
	if err := validate.ValidateStruct(createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}
	// create product
	if err := p.productService.CreateProduct(ctx, p.config, createRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to create product", err)
	}
	return response.Success(ctx, fiber.StatusOK, "Product created successfully", nil)
}

func (p productHandlerImp) Update(ctx *fiber.Ctx) error {
	productIdStr := ctx.Params("productId")
	productId, _ := strconv.Atoi(productIdStr)
	var updateRequest dto.UpdateProductRequest
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}
	// validate data
	if err := validate.ValidateStruct(updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body", err)
	}
	if err := p.productService.UpdateProduct(ctx, p.config, productId, updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to update product", err)
	}
	return response.Success(ctx, fiber.StatusOK, "Product updated successfully", nil)
}

func (p productHandlerImp) Delete(ctx *fiber.Ctx) error {
	productIdStr := ctx.Params("productId")
	productId, _ := strconv.Atoi(productIdStr)
	if err := p.productService.DeleteProduct(productId); err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "Product not found", err)
	}
	return response.Success(ctx, fiber.StatusOK, "Product deleted successfully", nil)
}
