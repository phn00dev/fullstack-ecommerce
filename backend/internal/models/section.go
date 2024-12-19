package models

import "time"

type Section struct {
	ID            uint64    `json:"id"`
	SectionNameTk string    `json:"section_name_tk"`
	SectionNameRu string    `json:"section_name_ru"`
	SectionNameEn string    `json:"section_name_en"`
	SectionSlug   string    `json:"section_slug"`
	SectionStatus string    `json:"section_status"`
	SectionIcon   string    `json:"section_icon"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
