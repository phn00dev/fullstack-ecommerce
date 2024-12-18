package repository

import "eCommerce/internal/models"

type SectionRepository interface {
	GetAll() ([]models.Section, error)
	GetOneById(sectionID int) (*models.Section, error)
	Create(section *models.Section) error
	Update(sectionID int, section *models.Section) error
	Delete(sectionID int) error

	// helper functions

	GetOneBySectionName(sectionName string) (*models.Section, error)
}
