package dto

import (
	"eCommerce/internal/models"
)

type SectionResponse struct {
	ID            uint64 `json:"id"`
	SectionName   string `json:"section_name"`
	SectionSlug   string `json:"section_slug"`
	SectionStatus string `json:"section_status"`
	SectionIcon   string `json:"section_icon"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	DeletedAt     string `json:"deleted_at"`
}

func GetAllSectionResponse(sections []models.Section) []SectionResponse {
	var sectionResponse []SectionResponse
	for _, section := range sections {
		sectionResponse = append(sectionResponse, GetOneSectionResponse(section))
	}
	return sectionResponse
}

func GetOneSectionResponse(section models.Section) SectionResponse {
	return SectionResponse{
		ID:            section.ID,
		SectionName:   section.SectionName,
		SectionSlug:   section.SectionSlug,
		SectionStatus: section.SectionStatus,
		SectionIcon:   section.SectionIcon,
		CreatedAt:     section.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt:     section.UpdatedAt.Format("02-01-2006 15:04:05"),
		DeletedAt:     section.DeletedAt.Format("02-01-2006 15:04:05"),
	}
}
