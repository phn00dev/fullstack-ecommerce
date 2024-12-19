package constructor

import (
	"eCommerce/internal/app"
	adminConstructor "eCommerce/internal/domain/admin/constructor"
	bannerConstructor "eCommerce/internal/domain/banner/constructor"
	brandConstructor "eCommerce/internal/domain/brand/constructor"
	sectionConstructor "eCommerce/internal/domain/section/constructor"
)

func Build(dependencies app.Dependencies) {
	adminConstructor.InitAdminRequirementCreator(dependencies.DB)
	bannerConstructor.InitBannerRequirementCreator(dependencies.DB, *dependencies.Config)
	brandConstructor.InitBrandRequirementCreator(dependencies.DB, *dependencies.Config)
	sectionConstructor.InitSectionRequirementCreator(dependencies.DB, *dependencies.Config)
}
