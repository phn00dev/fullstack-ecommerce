package service

import (
	"eCommerce/internal/domain/section/dto"
	"eCommerce/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type SectionService interface {
	GetAllSections() ([]dto.SectionResponse, error)
	GetOneSectionByID(sectionId int) (*dto.SectionResponse, error)
	CreateSection(ctx *fiber.Ctx, config config.Config, createRequest dto.CreateSectionRequest) error
	UpdateSection(ctx *fiber.Ctx, config config.Config, sectionID int, updateRequest dto.UpdateSectionRequest) error
	DeleteSection(sectionID int) error
}
