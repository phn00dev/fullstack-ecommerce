package service

import (
	"eCommerce/internal/domain/section/dto"
	"eCommerce/internal/domain/section/repository"
	"eCommerce/internal/models"
	"eCommerce/internal/utils/images"
	"eCommerce/pkg/config"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

type sectionServiceImp struct {
	sectionRepo repository.SectionRepository
}

func NewSectionService(sectionRepo repository.SectionRepository) SectionService {
	return sectionServiceImp{sectionRepo: sectionRepo}
}

func (s sectionServiceImp) GetAllSections() ([]dto.SectionResponse, error) {
	sections, err := s.sectionRepo.GetAll()
	if err != nil {
		return nil, err
	}
	sectionResponses := dto.GetAllSectionResponse(sections)
	return sectionResponses, nil
}

func (s sectionServiceImp) GetOneSectionById(sectionId int) (*dto.SectionResponse, error) {
	section, err := s.sectionRepo.GetOneById(sectionId)
	if err != nil {
		return nil, err
	}
	sectionResponse := dto.GetOneSectionResponse(*section)
	return &sectionResponse, nil
}

func (s sectionServiceImp) CreateSection(ctx *fiber.Ctx, config config.Config, createRequest dto.CreateSectionRequest) error {
	section, err := s.sectionRepo.GetOneBySectionName(createRequest.SectionName)
	if err != nil {
		return errors.New("Section name already exists")
	}
	if section.ID == 0 {
		return errors.New("Section name already exists")
	}

	// upload section icon image
	sectionIconPath, err := images.UploadFile(ctx, "section_icon", config.FolderConfig.PublicPath, "section-icons")
	if err != nil {
		return err
	}

	newSection := models.Section{
		SectionName:   createRequest.SectionIcon,
		SectionSlug:   slug.Make(createRequest.SectionName),
		SectionStatus: createRequest.SectionStatus,
		SectionIcon:   *sectionIconPath,
	}

	if err := s.sectionRepo.Create(&newSection); err != nil {
		// eger-de section doretmekde error bar bolsa onda upload edilen section icon image delete edilmeli
		if err := images.DeleteFile(*sectionIconPath); err != nil {
			return errors.New("Something wrong! Failed to delete section icon")
		}
		return err
	}
	return nil
}

func (s sectionServiceImp) UpdateSection(ctx *fiber.Ctx, config config.Config, sectionID int, updateRequest dto.UpdateSectionRequest) error {
	// get section
	section, err := s.sectionRepo.GetOneById(sectionID)
	if err != nil {
		return errors.New("section not found")
	}
	if section.ID == 0 {
		return errors.New("section not found")
	}

	// new section image barlamaly
	file, _ := ctx.FormFile("section_icon")
	if file != nil {
		// old section image delete
		if errOldSectionIconImage := images.DeleteFile(section.SectionIcon); errOldSectionIconImage != nil {
			return errOldSectionIconImage
		}
		// new image upload
		sectionIconPath, errUploadIcon := images.UploadFile(ctx, "section_icon", config.FolderConfig.PublicPath, "section-icons")
		if errUploadIcon != nil {
			return errUploadIcon
		}
		updateRequest.SectionIcon = *sectionIconPath
	}
	section.SectionName = updateRequest.SectionName
	section.SectionStatus = updateRequest.SectionStatus
	return s.sectionRepo.Update(sectionID, section)
}

func (s sectionServiceImp) DeleteSection(sectionID int) error {
	section, err := s.sectionRepo.GetOneById(sectionID)
	if err != nil {
		return err
	}
	// delete section icon
	if errDeleteSectionIcon := images.DeleteFile(section.SectionIcon); errDeleteSectionIcon != nil {
		return errDeleteSectionIcon
	}
	// delete section
	return s.sectionRepo.Delete(sectionID)
}
