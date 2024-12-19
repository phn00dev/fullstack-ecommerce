package dto

type CreateSectionRequest struct {
	SectionNameTk string `json:"section_name_tk"`
	SectionNameRu string `json:"section_name_ru"`
	SectionNameEn string `json:"section_name_en"`
	SectionStatus string `json:"section_status"`
	SectionIcon   string `json:"section_icon"`
}

type UpdateSectionRequest struct {
	SectionNameTk string `json:"section_name_tk"`
	SectionNameRu string `json:"section_name_ru"`
	SectionNameEn string `json:"section_name_en"`
	SectionStatus string `json:"section_status"`
	SectionIcon   string `json:"section_icon"`
}
