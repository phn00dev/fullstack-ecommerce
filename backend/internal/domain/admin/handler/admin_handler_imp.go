package handler

import (
	"eCommerce/internal/domain/admin/dto"
	"eCommerce/internal/domain/admin/service"
	"eCommerce/internal/utils/response"
	"eCommerce/internal/utils/validate"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type adminHandlerImp struct {
	adminService service.AdminService
}

func NewAdminHandler(adminService service.AdminService) AdminHandler {
	return adminHandlerImp{adminService: adminService}
}

func (a adminHandlerImp) GetAll(ctx *fiber.Ctx) error {
	admins, err := a.adminService.GetAllAdmins()
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "admins not found", err.Error())
	}
	return response.Success(ctx, fiber.StatusOK, "admins", admins)
}

func (a adminHandlerImp) GetOne(ctx *fiber.Ctx) error {
	adminIDStr := ctx.Params("adminID")
	adminID, _ := strconv.Atoi(adminIDStr)
	// getAdmin
	admin, err := a.adminService.GetAdminByID(adminID)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "admin not found", err.Error())
	}
	return response.Success(ctx, fiber.StatusOK, "admin", admin)
}

func (a adminHandlerImp) Create(ctx *fiber.Ctx) error {
	var createRequest dto.CreateAdminRequest
	// body parser
	if err := ctx.BodyParser(&createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "body parse failed", err.Error())
	}
	// validate data
	if err := validate.ValidateStruct(createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "validation failed", err.Error())
	}
	// validate phoneNumber
	if !validate.ValidatePhoneNumber(createRequest.PhoneNumber) {
		return response.Error(ctx, http.StatusBadRequest, "Invalid phone number", "Phone number format is not valid")
	}
	// create admin
	if err := a.adminService.CreateAdmin(createRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "admin creation failed", err.Error())
	}
	// return response
	return response.Success(ctx, fiber.StatusCreated, "admin created successfully", nil)
}

func (a adminHandlerImp) Update(ctx *fiber.Ctx) error {
	var updateRequest dto.UpdateAdminRequest

	adminIDStr := ctx.Params("adminID")
	adminID, _ := strconv.Atoi(adminIDStr)
	// body parser
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "body parse failed", err.Error())
	}
	// validate data
	if err := validate.ValidateStruct(updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "validation failed", err.Error())
	}
	// validate phoneNumber
	if !validate.ValidatePhoneNumber(updateRequest.PhoneNumber) {
		return response.Error(ctx, http.StatusBadRequest, "Invalid phone number", "Phone number format is not valid")
	}
	// update admin
	if err := a.adminService.UpdateAdmin(adminID, updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "admin update failed", err.Error())
	}
	return response.Success(ctx, fiber.StatusOK, "admin updated successfully", nil)
}

func (a adminHandlerImp) Delete(ctx *fiber.Ctx) error {
	adminIDStr := ctx.Params("adminID")
	adminID, _ := strconv.Atoi(adminIDStr)
	err := a.adminService.DeleteAdmin(adminID)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "admin not found", err.Error())
	}
	return response.Success(ctx, fiber.StatusOK, "admin deleted successfully", nil)
}
