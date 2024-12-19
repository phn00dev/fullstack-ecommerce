package repository

import (
	"eCommerce/internal/models"
	"gorm.io/gorm"
)

type sectionRepositoryImp struct {
	db *gorm.DB
}

func NewSectionRepository(db *gorm.DB) SectionRepository {
	return sectionRepositoryImp{db: db}
}

func (s sectionRepositoryImp) GetAll() ([]models.Section, error) {
	var sections []models.Section
	if err := s.db.Find(&sections).Error; err != nil {
		return nil, err
	}
	return sections, nil
}

func (s sectionRepositoryImp) GetOneById(sectionID int) (*models.Section, error) {
	var section models.Section
	if err := s.db.First(&section, sectionID).Error; err != nil {
		return nil, err
	}
	return &section, nil
}

func (s sectionRepositoryImp) Create(section *models.Section) error {
	return s.db.Create(&section).Error
}

func (s sectionRepositoryImp) Update(sectionID int, section *models.Section) error {
	return s.db.Model(&models.Section{}).Where("id=?", sectionID).Updates(&section).Error
}

func (s sectionRepositoryImp) Delete(sectionID int) error {
	return s.db.Delete(&models.Section{}, sectionID).Error
}

func (s sectionRepositoryImp) GetOneBySectionName(sectionName string) (*models.Section, error) {
	var section models.Section
	if err := s.db.Where("section_name=?", sectionName).First(&section).Error; err != nil {
		return nil, err
	}
	return &section, nil
}

func (s sectionRepositoryImp) CheckSectionNameTk(sectionNameTk string) (bool, error) {
	err := s.db.Where("section_name_tk=?", sectionNameTk).First(&models.Section{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s sectionRepositoryImp) CheckSectionNameRu(sectionNameRu string) (bool, error) {
	err := s.db.Where("section_name_ru=?", sectionNameRu).First(&models.Section{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s sectionRepositoryImp) CheckSectionNameEn(sectionNameEn string) (bool, error) {
	err := s.db.Where("section_name_en=?", sectionNameEn).First(&models.Section{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	return true, nil
}
