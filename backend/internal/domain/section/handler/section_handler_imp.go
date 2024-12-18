package handler

import (
	"eCommerce/internal/domain/section/dto"
	"eCommerce/internal/domain/section/service"
	"eCommerce/internal/utils/response"
	"eCommerce/internal/utils/validate"
	"eCommerce/pkg/config"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type sectionHandlerImp struct {
	sectionService service.SectionService
	config         config.Config
}

func NewSectionHandler(sectionService service.SectionService, config config.Config) SectionHandler {
	return sectionHandlerImp{
		sectionService: sectionService,
		config:         config,
	}
}

func (s sectionHandlerImp) GetAll(ctx *fiber.Ctx) error {
	sections, err := s.sectionService.GetAllSections()
	if err != nil {
		return response.Error(ctx, fiber.StatusNotFound, "sections not found", err)
	}
	return response.Success(ctx, fiber.StatusOK, "sections", sections)
}

func (s sectionHandlerImp) GetOne(ctx *fiber.Ctx) error {
	sectionIdStr := ctx.Params("sectionID")
	sectionID, _ := strconv.Atoi(sectionIdStr)
	// get section
	section, err := s.sectionService.GetOneSectionById(sectionID)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "section not found", err)
	}
	return response.Success(ctx, fiber.StatusOK, "section", section)
}

func (s sectionHandlerImp) Create(ctx *fiber.Ctx) error {
	var createRequest dto.CreateSectionRequest
	if err := ctx.BodyParser(&createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "body parse failed", err)
	}
	// validate data
	if err := validate.ValidateStruct(createRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "validation failed", err)
	}
	// create request
	if err := s.sectionService.CreateSection(ctx, s.config, createRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "section creation failed", err)
	}
	return response.Success(ctx, fiber.StatusOK, "section created successfully", nil)
}

func (s sectionHandlerImp) Update(ctx *fiber.Ctx) error {
	var updateRequest dto.UpdateSectionRequest
	sectionIdStr := ctx.Params("sectionID")
	sectionID, _ := strconv.Atoi(sectionIdStr)
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "body parse failed", err)
	}
	// validate data
	if err := validate.ValidateStruct(updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	// update section
	if err := s.sectionService.UpdateSection(ctx, s.config, sectionID, updateRequest); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "section update failed", err)
	}
	return response.Success(ctx, fiber.StatusOK, "section updated successfully", nil)
}

func (s sectionHandlerImp) Delete(ctx *fiber.Ctx) error {
	sectionIdStr := ctx.Params("sectionID")
	sectionID, _ := strconv.Atoi(sectionIdStr)
	if err := s.sectionService.DeleteSection(sectionID); err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "section delete failed", err)
	}
	return response.Success(ctx, fiber.StatusOK, "section deleted successfully", nil)
}
