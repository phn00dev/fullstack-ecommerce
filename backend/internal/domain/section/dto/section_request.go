package dto

type CreateSectionRequest struct {
	SectionName   string `json:"section_name"`
	SectionStatus string `json:"section_status"`
	SectionIcon   string `json:"section_icon"`
}

type UpdateSectionRequest struct {
	SectionName   string `json:"section_name"`
	SectionStatus string `json:"section_status"`
	SectionIcon   string `json:"section_icon"`
}
