package models

import "time"

type Section struct {
	ID            uint64    `json:"id"`
	SectionName   string    `json:"section_name"`
	SectionSlug   string    `json:"section_slug"`
	SectionStatus string    `json:"section_status"`
	SectionIcon   string    `json:"section_icon"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
